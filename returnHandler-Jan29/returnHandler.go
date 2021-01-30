package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//Return single item
func returnSingleQuestion(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	// fmt.Fprintf(w, "Key: "+key)

	// Loop over all of our Questions
	// if the question.Id equals the key we pass in
	// return the question encoded as JSON
	for _, question := range Questions {
		if question.Id == key {
			json.NewEncoder(w).Encode(question)
		}

	}
}

func returnVideo(w http.ResponseWriter, r *http.Request) {
	vediovars := mux.Vars(r)
	vediokey := vediovars["id"]

	for _, vedio := range Videos {
		if vedio.Id == vediokey {
			json.NewEncoder(w).Encode(vedio)

		}
	}
}

//Return All
func returnAllQuestions(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: returnAllQuestions")
	//parsing
	json.NewEncoder(w).Encode(Questions)
}
func returnAllVideo(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Videos)
}
