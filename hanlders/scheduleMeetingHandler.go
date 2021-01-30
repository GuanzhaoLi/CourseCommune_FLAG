package main

import {

}

// receive JSON, return ?
func scheduleMeetingHandler(w http.ResponseWriter, r *http.Request) {
	// get the question and schedule
	question := Question {

	}
	schedule := Schedule {
		tutorID: r.formValue("tutorID"),
		// ...
	}

	// send schedule to tutor selected
	err := sendScheduleToTutor(&schedule)
	if (err != nil) {

	}
}

func sendScheduleToTutor(schedule *Schedule) error {
	// how to inform tutor? email or app UI?
	// what to do with tutor's approval or denial?
	//     approval: update database that tutor is not available at some time period
	//     denial: inform student to re-schedule
}