package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	_ "github.com/go-sql-driver/mysql"
)


func handleRequestsOfQuestion() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/login", Login)
	// student post a video tutoring request
	myRouter.HandleFunc("/student={user}/requestvideo", studentRequestVideo).Methods("POST")	
	// student search for tutors 
	myRouter.HandleFunc("/student={user}/searchtutors", studentSearchTutors).Methods("GET")
	// tutor post answer to picked question
	myRouter.HandleFunc("/tutor={user}/pickquestion", tutorPickQuestion).Methods("POST")
	// tutor post picked video reqest
	myRouter.HandleFunc("/tutor={user}/pickvideo", tutorPickVideo).Methods("POST")
	// tutor search for un-answered questions
	myRouter.HandleFunc("/tutor={user}/searchquestions", tutorSearchQuestions).Methods("GET")
	// tutor search for video requests
	myRouter.HandleFunc("/tutor={user}/searchvideos", tutorSearchVideos).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Started")
	InitDB()
	handleRequestsOfQuestion()
}
