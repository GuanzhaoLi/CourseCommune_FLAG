package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Schedule struct {
	subject     string `json:"subject"`
	startTime 	string `json:"startTime"`
	endTime		string `json:"endTime"`
	level     	string `json:"level"`
}

type Schedule struct {
	// tutorID, subject, level(high school, college?), date, time, questionText ...
}

type Tutor struct {
	// tutorID, tutorName, imageURL, subject, level, available table...
}

// receive JSON, return JSON
func tutorSearchHandler(w http.ResponseWriter, r *http.Request) {
	// get the Schedule
	schedule := Schedule {
		subject: r.formValue("subject"),
		startTime: r.formValue("startTime") 
		endTime: r.formValue("endTime") 
		level: r.formValue("level")
	}

	// search for available tutors
	tutors, err := findTutors(&schedule)

	// encode to JSON
	js, err := JSON.Marshal(tutors)

	w.write(js)
}

func findTutors(schedule *Schedule) []Tutor {
	// create database client
	db, err = sql.open("mysql", "root:pass1@tcp(127.0.0.1:3306)/Tutor") // need modify dataSourceName
	                                                                     // name, password, and IP address
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// search for available tutors, based on subject, level, and available table
	results, err := db.Query("SELECT * FROM Tutor WHERE Level = schedule.level 
		AND schedule.startTime >= StartTime 
		AND schedule.endTime <= EndTime") // 命令可能不对

	if err != nil {
		panic(err.Error())
	}

	// return a list of Tutors
	return results // ??
}