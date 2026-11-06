package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/vlonebara/todo-list/pkg/api"
	"github.com/vlonebara/todo-list/pkg/db"
	"github.com/vlonebara/todo-list/pkg/service"
)

// Run запускает сервер и корректно останавливает его по сигналу ОС.
func Run() error {
	database, err := db.OpenDatabase()
	if err != nil {
		return err
	}
	defer func() {
		if err := database.Close(); err != nil {
			log.Printf("close database: %v", err)
		}
	}()

	store := db.NewTaskStore(database)
	tasks := service.NewTaskService(store)
	mux := http.NewServeMux()
	api.Init(mux, tasks)

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}
	addr := ":" + port

	fileServer := http.FileServer(http.Dir("./web"))
	mux.Handle("/", fileServer)

	server := &http.Server{Addr: addr, Handler: mux}
	shutdown, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	log.Printf("Server is listening on %s", addr)
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- server.ListenAndServe()
	}()

	select {
	case err := <-serverErrors:
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}
		return err
	case <-shutdown.Done():
		log.Println("Shutting down server...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			_ = server.Close()
			return fmt.Errorf("shutdown server: %w", err)
		}
		if err := <-serverErrors; !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	}
}
