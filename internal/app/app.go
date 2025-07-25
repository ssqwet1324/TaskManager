package app

import (
	"log"
	"net/http"
	"taskmanager/internal/handlers"
	"taskmanager/internal/repository"
)

func Run() {
	repo := repository.New()

	globalHandler := handlers.New(repo)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	})

	mux.HandleFunc("/api/v1/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			globalHandler.GetTasks(w, r)
		case http.MethodPost:
			globalHandler.AddTask(w, r)
		case http.MethodPut:
			globalHandler.EditTask(w, r)
		case http.MethodDelete:
			globalHandler.DeleteTask(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	// Создаём сервер
	server := http.Server{
		Addr:    ":8081",
		Handler: mux,
	}

	// Запускаем сервер
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
