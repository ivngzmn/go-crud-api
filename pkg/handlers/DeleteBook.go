package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/mocks"
)

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	// read the dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// iterate over all the mock Books
	for index, book := range mocks.Books {
		if book.Id == id {
			// delete book and send response if the book Id matches dynamic Id
			mocks.Books = append(mocks.Books[:index], mocks.Books[index+1:]...)

			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("Deleted")
			break
		}
	}

}