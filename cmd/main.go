package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/handlers"
)


func main() {
	router := mux.NewRouter()

	router.HandleFunc("/books", handlers.GetAllBooks).Methods(http.MethodGet) // read
	router.HandleFunc("/books/{id}", handlers.GetBook).Methods(http.MethodGet) // read
	router.HandleFunc("/books", handlers.AddBook).Methods(http.MethodPost) // create
	router.HandleFunc("/books/{id}", handlers.UpdateBook).Methods(http.MethodPut) // update
	router.HandleFunc("/books/{id}", handlers.DeleteBook).Methods(http.MethodDelete)// deletay
	log.Println("The GO server is running you better GO catch it!")
	http.ListenAndServe(":4000", router)
}
