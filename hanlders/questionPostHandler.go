package main

import {

}

// receive JSON, return ?
func questionPostHandler(w http.ResponseWriter, r *http.Request) {
	//get the question
	q := Question{

	}

	// post the question
	err := saveQuestionPost(&q)
	if err != nil {

	}
}

func saveQuestionPost(question *Question) error {
	// create database client

	// save to data base

	// return error if any
}