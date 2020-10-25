package main

import (
	"log"
	"net/http"

	"github.com/EduAR/quizapi/handlers"
	"github.com/EduAR/quizapi/repository"
	"github.com/gorilla/mux"
)

func main() {
	repo := repository.NewRepository(
		"mongodb://localhost:27017",
		"eduAR",
		"questions",
	)
	defer repo.Close()

	h := handlers.Handlers{
		Repo: repo,
	}

	r := mux.NewRouter()
	r.HandleFunc("/quizapi", h.All).Methods("GET")
	r.HandleFunc("/quizapi/{category}", h.GetByCat).Methods("GET")

	r.HandleFunc("/quizapi", h.Insert).Methods("POST")
	r.HandleFunc("/quizapi/{category}", h.Delete).Methods("DELETE")
	r.HandleFunc("/quizapi/{category}", h.Update).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8080", r))

}
