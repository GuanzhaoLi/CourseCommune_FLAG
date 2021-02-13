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

func requestVideo(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.body)
	w.Header().Set("Content-Type", "application/json") 
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	videoRequest := VideoOrder {
		tutorID: 	r.formValue("tutorID"),
		studentID: 	r.formValue("studentID"),
		startTime : r.formValue("startTime"),
		endTime: 	r.formValue("endTime"),
		level: 		r.formValue("level"),
		subject: 	r.formValue("subject"),
		keywords: 	r.formValue("keywords"),
		finished:   false,
	}

	// send schedule to tutor selected
	err := sendScheduleToTutor(&videoRequest)
	if (err != nil) {
		http.Error(w, "failed to send video request", http.StatusInternalServerError)
	}
	w.wirteHeader(200)
}

func sendScheduleToTutor(vr *VideoRequest) error {
	db, err1 = sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/VideoOrder") // need modify dataSourceName
	                                                                    // name, password, and IP address
	if err1 != nil {
		return err1
	}
	defer db.Close()

	id := uuid.New()
	insert, err2 := db.Query("INSERT INTO VideoOrder VALUES (id, vr.startTime, vr.endTime, vr.studentID, 
		vr.tutorID, "30 min", vr.subject, vr.level, vr.keywords)")

	return err2
}