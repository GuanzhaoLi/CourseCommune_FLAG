package main

import (

)

//url: /signup
func signupHandler(w http.ResponseWriter, r *http.Request) {
	//request body should contain a user

	//check if the username already exists

	// will return http.StatusBadRequest if username already exists



}

//url: /signin
func signinHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	//It returns an encrypted token if sign in succeeds.
	//It returns status unauthorized if username or password is incorrect.
 	//It could also return an status internal server error if an error occurs that is caused by the database operations or the token encryption.

}

//url: /answer
func answerQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	//are we gonna have several different answers for one question?
	//request should contain question id and answer?
}

