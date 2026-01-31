package http

import (
	"net/http"
	"restapi/todo"
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

}

func (h *HTTPHanlers) HandleGetTask(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHanlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {

}

func (h *HTTPHanlers) HandleGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	
}
