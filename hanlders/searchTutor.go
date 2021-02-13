package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Tutor struct {
	tutorId     int64 `json:"tutorId"`
	firstName   string `json:"firstName"`
	lastName   	string `json:"firstName"`
	rating		int64 `json:"rating"`
}

type Criteria struct {
	level     int64 `json:"level"`
	subject   string `json:"subject"`
}

// receive JSON, return JSON
func searchTutor(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.body)
	w.Header().Set("Content-Type", "application/json") 
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	c := Criteria {
		level: 		r.formValue("level"),
		subject: 	r.formValue("subject"),
	}
	// search for available tutors
	tutors, err := findTutors(&c)
	// encode to JSON
	js, err := JSON.Marshal(tutors)
	if (err != nil) {
		http.Error(w, "Failed to read tutors from mySQL", http.StatusInternalServerError)
	}
	w.write(js)
}

func findTutors(c *Criteria) []Tutor {
	db, err1 = sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/Tutor") // need modify dataSourceName
	                                                                    // name, password, and IP address
	if err1 != nil {
		panic(err.Error())
	}
	defer db.Close()

	// search for available tutors based on ?
	results, err2 := db.Query("SELECT TutorId, FirstName, LastName, Rating FROM Tutor WHERE 
	Level = c.level AND Subject = c.subject") 

	if err2 != nil {
		panic(err.Error())
	}
	// get all tutors searched from database
	var tutors []Tutor
	for results.Next() {
		var tTutor
		err3 := results.Scan(&t.id, &t.firstName, &t.lastName, &t.rating)
		if (err3 != nil) {
			panic(err)
		}
		tutors.append(tutors, t)
	}
	return tutors
}