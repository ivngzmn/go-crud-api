package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tutorials/go/crud/pkg/models"
)

func (h handler) AddBook(w http.ResponseWriter, r *http.Request) {
	//read to request body
	defer r.Body.Close()
	body, err :=  ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var book models.Book
	json.Unmarshal(body, &book)
	// append to the books table
	if result := h.DB.Create(&book); result.Error != nil {
		fmt.Println(result.Error)
	}


	// send a 201 created response
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Created")
}