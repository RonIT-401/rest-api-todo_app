package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restapi/todo"
	"time"
)

type HTTPHanlers struct {
	todoList *todo.List
}

func NewHTTPHandlers(todoList *todo.List) *HTTPHanlers {
	return &HTTPHanlers{
		todoList: todoList,
	}
}

func (h *HTTPHanlers) HandleCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO

	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	if err := TaskDTO.ValidateForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}

	todoTask := todo.NewTask(taskDTO.Title, taskDTO.Description)
	if err := h.todoList.AddTask(todoTask); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, todo.ErrTaskAlreadyExist) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err :=json.MarshalIndent(todoTask, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *HTTPHanlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHanlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHanlers) HandleGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHanlers) HandleCompleteTasks(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHanlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {

}
