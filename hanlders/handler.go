package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"

	"time"

	jwt "github.com/dgrijalva/jwt-go"
)


// receive JSON, return ?
func questionPostHandler(w http.ResponseWriter, r *http.Request) {
	//setting:
	fmt.Println("Received one request for upload")
	w.Header().Set("Content-Type", "application/json") // return type is json
	//cross domain access:option
	w.Header().Set("Access-Control-Allow-Origin", "*")                           // can be accessed by any domain
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization") //back end allow http header which is content -type and Authorization
	// if request is options do nothing just return ok
	if r.Method == "OPTIONS" {
		return
	}

	//get user's information from body
	user := r.Context().Value("user")
	claims := user.(*jwt.Token).Claims
	username := claims.(jwt.MapClaims)["username"]

	q := Post{
		User:    username.(string),
		Message: r.FormValue("message"),
	}

	// post the question
	err := saveQuestion(&p)
	if err != nil {
		http.Error(w, "Failed to post question", http.StatusInternalServerError)
		return
	}
	fmt.Println("Post is saved successfully")
}


func QuestionSearchHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received one request for search")
	w.Header().Set("Content-Type", "application/json")
	//cross domain access:option
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	if r.Method == "OPTIONS" {
		return
	}
	keywords := r.URL.Query().Get("keywords")

	var questions []Question
	var err error
	quesitons, err = searchQuestionsByKeywords(keywords)

	if err != nil {
		http.Error(w, "Failed to read question from Elasticsearch", http.StatusInternalServerError)
		fmt.Printf("Failed to read question from Elasticsearch %v.\n", err)
		return
	}

	js, err := json.Marshal(questions)
	if err != nil {
		http.Error(w, "Failed to parse questions into JSON format", http.StatusInternalServerError)
		fmt.Printf("Failed to parse questions into JSON format %v.\n", err)
		return
	}
	w.Write(js)
}

// signin sign up handler:
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