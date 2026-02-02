package main

import (
	"fmt"
	"restapi/http"
	"restapi/todo"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := http.NewHTTPHandlers(todoList)
	httpServer := http.NewHTTPHServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server:", err)
	}
}