package main

import (
	"awesomeProject/internal/core/application"
	"awesomeProject/internal/framework/primary/rest"
	"awesomeProject/internal/framework/secondary/memory"
	"fmt"
	"net/http"
)

func main() {
	repo := memory.NewMemoryRepository()
	service := application.NewTodoService(repo)
	handler := rest.NewAppHandler(service)

	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handler.GetAllTodos(w)
		} else if r.Method == http.MethodPost {
			handler.CreateTodo(w, r)
		} else {
			http.NotFound(w, r)
		}
	})
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetTodoByID(w, r)
		case http.MethodPut:
			handler.UpdateTodoStatus(w, r)
		case http.MethodDelete:
			handler.DeleteTodoByID(w, r)
		default:
			http.NotFound(w, r)
		}
	})

	fmt.Println("Server started on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("something went wrong %s \n", err)
	}

}
