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
	myRouter.HandleFunc("/user/tuotorsignup", tutorSignupHandler).Methods("POST")
	myRouter.HandleFunc("/user/studentsignup", studentSignupHandler).Methods("POST")
	myRouter.HandleFunc("/user/signin", SigninHandler).Methods("POST")
	myRouter.HandleFunc("/student={user}/history", studentHistory).Methods("GET")
	myRouter.HandleFunc("/tutor={user}/history", tutorHistory).Methods("GET")
	myRouter.HandleFunc("/student={user}/postquestion", studentPostQuestion).Methods("POST")
	myRouter.HandleFunc("/student={user}/searchqustion", studentSearchQuestions).Methods("GET")
	myRouter.HandleFunc("/student={user}/requestvideo", studentRequestVideo).Methods("POST")
	myRouter.HandleFunc("/student={user}/searchtutor", studentSearchTutors).Methods("POST")
	myRouter.HandleFunc("/tutor={user}/pickquestion", tutorPickQuestion).Methods("POST")
	myRouter.HandleFunc("/tutor={user}/pickvideo", tutorPickVideo).Methods("POST")
	myRouter.HandleFunc("/tutor={user}/searchquestions", tutorSearchQuestions).Methods("GET")
	myRouter.HandleFunc("/tutor={user}/searchvideos", tutorSearchVideos).Methods("GET")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Started")
	InitDB()
	handleRequestsOfQuestion()
}
