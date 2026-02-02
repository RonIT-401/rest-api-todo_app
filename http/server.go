package http

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPServer struct {
	httpHanlers *HTTPHanlers
}

func NewHTTPHServer(httpHandler *HTTPHanlers) *HTTPServer {
	return &HTTPServer{
		httpHanlers: httpHandler,
	}
}

func (s *HTTPServer) StartServer() error {
	router := mux.NewRouter()

	router.Path("/tasks").Methods("POST").HandlerFunc(s.httpHanlers.HandleCreateTask)
	router.Path("/tasks/{title}").Methods("GET").HandlerFunc(s.httpHanlers.HandleGetTask)
	router.Path("/tasks").Methods("GET").Queries("comleted", "true").HandlerFunc(s.httpHanlers.HandleGetAllUncompletedTasks)
	router.Path("/tasks").Methods("GET").HandlerFunc(s.httpHanlers.HandleGetAllTasks)
	router.Path("/tasks/{title}").Methods("PATCH").HandlerFunc(s.httpHanlers.HandleCompleteTasks)
	router.Path("/tasks/{title}").Methods("DELETE").HandlerFunc(s.httpHanlers.HandleDeleteTask)

	if err := http.ListenAndServe(":9091", router); err != nil {
		if errors.Is(err, http.ErrServerClosed) {
			return nil
		}

		return err
	}

	return nil
}
