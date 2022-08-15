package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/db"
	"github.com/tutorials/go/crud/pkg/handlers"
)


func main() {
	// initialize the database
	DB := db.Init()
	// pass the database to the handlers
	h := handlers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/books", h.GetAllBooks).Methods(http.MethodGet) // read
	router.HandleFunc("/books/{id}", h.GetBook).Methods(http.MethodGet) // read
	router.HandleFunc("/books", h.AddBook).Methods(http.MethodPost) // create
	router.HandleFunc("/books/{id}", h.UpdateBook).Methods(http.MethodPut) // update
	router.HandleFunc("/books/{id}", h.DeleteBook).Methods(http.MethodDelete)// deletay
	log.Println("The GO server is running you better GO catch it!")
	http.ListenAndServe(":4000", router)
}
