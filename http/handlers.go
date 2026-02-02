package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"restapi/todo"
	"time"

	"github.com/gorilla/mux"
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

	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
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
	title:= mux.Vars(r)["title"]
	
	task, err := h.todoList.GetTask(title)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time: time.Now(),
		}

		if errors.Is(err, todo.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}

		return
	}

	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response")
		return
	}
}

func (h *HTTPHanlers) HandleGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

func (h *HTTPHanlers) HandleGetAllUncompletedTasks(w http.ResponseWriter, r *http.Request) {
	uncomletedTasks := h.todoList.ListUncomletedTasks()
	b, err := json.MarshalIndent(uncomletedTasks, "", "    ")
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

}

func (h *HTTPHanlers) HandleCompleteTasks(w http.ResponseWriter, r *http.Request) {
	var CompleteDTO CompleteTaskDTO
	if err := json.NewDecoder(r.Body).Decode(&CompleteDTO); err != nil {
		errDTO := ErrorDTO {
			Message: err.Error(),
			Time: time.Now(),
		}

		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return
	}

	title := mux.Vars(r)["title"]

	if CompleteDTO.Complete {
		if err := h.todoList.CompleteTask(title); err != nil {
			errDTO := ErrorDTO {
				Message: err.Error(),
				Time: time.Now(),
			}

			if errors.Is(err, todo.ErrTaskNotFound) {
				http.Error(w, errDTO.ToString(), http.StatusNotFound)
			} else {
				http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			}

			return
		}
	} else {
		if err := h.todoList.UncomleteTask(title); err != nil {
			errDTO := ErrorDTO {
				Message: err.Error(),
				Time: time.Now(),
			}

			if errors.Is(err, todo.ErrTaskNotFound) {
				http.Error(w, errDTO.ToString(), http.StatusNotFound)
			} else {
				http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			}

			return
	}
}
}

func (h *HTTPHanlers) HandleDeleteTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]

	if err := h.todoList.DeleteTask(title); err != nil {
		errDTO := ErrorDTO {
				Message: err.Error(),
				Time: time.Now(),
			}

			if errors.Is(err, todo.ErrTaskNotFound) {
				http.Error(w, errDTO.ToString(), http.StatusNotFound)
			} else {
				http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
			}

			return
	}
}
 
