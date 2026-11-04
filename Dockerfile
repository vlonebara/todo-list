FROM golang:1.26 AS builder

WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY cmd/ ./cmd/
COPY pkg/ ./pkg/

RUN CGO_ENABLED=0 GOOS=linux go build -o /todo-list ./cmd

FROM ubuntu:latest

WORKDIR /app

COPY --from=builder /todo-list ./todo-list
COPY web/ ./web/

ENV TODO_PORT=7540 TODO_DBFILE=/app/scheduler.db

EXPOSE 7540

CMD ["./todo-list"]
