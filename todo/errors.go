package todo

import "errors"

var ErrTaskNotFound = errors.New("Task not found")
var ErrTaskAlreadyExist = errors.New("TaskAlreadyExist")