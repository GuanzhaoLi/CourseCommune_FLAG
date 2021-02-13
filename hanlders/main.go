package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// let's declare a global Quetsions and videos array
// that we can then populate in our main function
// to simulate a database
var Questions []Question
var Videos []Video

func QuestionPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: QuestionPage")
}
func handleRequestsOfQuestion() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", http.HandlerFunc(QuestionPage)).Methods("GET")
	myRouter.HandleFunc("/{user}/questions", returnAllQuestions).Methods("GET", "POST")
	myRouter.HandleFunc("/{user}/questions/{id}", returnSingleQuestion).Methods("GET")
	myRouter.HandleFunc("/{user}/video", returnAllVideo).Methods("GET", "POST")
	myRouter.HandleFunc("/{user}/video/{id}", returnVideo).Methods("GET")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))

	// jianheng added. 照葫芦画瓢
	myRouter.handleFunc("/{user}/tutor_search", searchTutor).Methods("GET")
	http.handleFunc("/{user}/schedule_meeing", requestVideo).Methods("POST")
}

func main() {
	fmt.Println("Started")
	//Get Questions from database
	Questions = []Question{
		Question{Id: "1", Name: "Sin", Desc: "Article Description", Content: "Article Content"},
		Question{Id: "2", Name: "Lee", Desc: "Article Description", Content: "Article Content"},
	}
	Videos = []Video{
		Video{Id: "1", User: "a", Date: "02/28/1993", Subject: "Math", Description: "Math question", Fulefilled: "ture"},
		Video{Id: "2", User: "b", Date: "03/12/1993", Subject: "English", Description: "English q", Fulefilled: "False"},
	}
	handleRequestsOfQuestion()

}
