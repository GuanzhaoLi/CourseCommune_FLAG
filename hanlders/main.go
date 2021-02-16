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
	myRouter.HandleFunc("/user/tuotorsignup", tutorSignupHandler).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/user/studentsignup", studentSignupHandler).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/user/signin", SigninHandler).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/student={user}/history", studentHistory).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/tutor={user}/history", tutorHistory).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/student={user}/postquestion", studentPostQuestion).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/student={user}/searchqustion", studentSearchQuestions).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/student={user}/requestvideo", studentRequestVideo).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/student={user}/searchtutor", studentSearchTutors).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/tutor={user}/pickquestion", tutorPickQuestion).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/tutor={user}/pickvideo", tutorPickVideo).Methods("POST", "OPTIONS")
	myRouter.HandleFunc("/tutor={user}/searchquestions", tutorSearchQuestions).Methods("GET", "OPTIONS")
	myRouter.HandleFunc("/tutor={user}/searchvideos", tutorSearchVideos).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Started")
	InitDB()
	handleRequestsOfQuestion()
}
