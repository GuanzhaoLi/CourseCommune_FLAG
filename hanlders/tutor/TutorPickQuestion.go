package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
	"github.com/pborman/uuid"
)

func tutorPickQuestion(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json") 
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	var question QuestionOrder
	err1 := json.NewDecoder(r.Body).Decode(&question)
    if err1 != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	
	err2 := postToDB(&question)
	if (err2 != nil) {
		http.Error(w, "failed to send video request", http.StatusInternalServerError)
	}
	w.wirteHeader(200)
}

func postToDB(q *QuestionOrder) error {
	db, err1 = sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/QuestionOrder") // need modify dataSourceName
	                                                                    // name, password, and IP address
	if err1 != nil {
		return err1
	}
	defer db.Close()

	_, err2 := db.Exec("update QuestionOrder set FulfilledBy = q.TutorID, Answer = q.Answer, Finished = true 
	where QId = q.QId")

	return err2
}