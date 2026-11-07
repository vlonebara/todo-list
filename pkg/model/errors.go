package model

import "errors"

// ErrInvalidTask сообщает об ошибке в данных задачи.
var ErrInvalidTask = errors.New("invalid task")

// ErrTaskNotFound сообщает, что задачи с указанным ID нет.
var ErrTaskNotFound = errors.New("task not found")
