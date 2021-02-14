package main

import (
	"fmt"
	_ "fmt"
	"github.com/gorilla/mux"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
	_ "github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
)


func handleRequestsOfQuestion() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/login", Login)
	myRouter.HandleFunc("/student={user}/history", studentHistory).Methods("GET")
	myRouter.HandleFunc("/tutor={user}/history", tutorHistory).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	fmt.Println("Started")
	InitDB()
	handleRequestsOfQuestion()
}