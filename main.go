package main

import (
	"TodoList/myhttp"
	"TodoList/todo"
	"fmt"
)

func main() {
	todoList := todo.NewList()
	httpHandlers := myhttp.NewHTTPHandlers(todoList)
	httpServer := myhttp.NewHTTPServer(httpHandlers)

	if err := httpServer.StartServer(); err != nil {
		fmt.Println("failed to start http server", err)
	}
}
