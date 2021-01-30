package main

import {

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
		questionText: r.formValue("questionText"),
		tutorID: r.formValue("tutorID") // empty here
		// ... 
	}

	// search for available tutors
	tutors, err := findTutors(&schedule)

	// encode to JSON
	js, err := JSON.Marshal(tutors)

	w.write(js)
}

func findTutors(schedule *Schedule) []Tutor {
	// create database client

	// search for available tutors, based on subject, level, and available table

	// return a list of Tutors
}