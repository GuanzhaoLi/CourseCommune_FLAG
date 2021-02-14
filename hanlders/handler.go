package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

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
func studentSignupHandler(w http.ResponseWriter, r *http.Request) {
	//request body should contain a user

	//check if the username already exists

	// will return http.StatusBadRequest if username already exists

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var student Student
	if err := decoder.Decode(&student); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode user data from client: %v\n", err)
		return
	}

	if student.Firstname == "" || student.Lastname == "" {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		fmt.Printf("Invalid username or password\n")
		return
	}

	success, err := addStudent(&student)

	if err != nil {
		http.Error(w, "Failed to save user to database", http.StatusInternalServerError)
		fmt.Printf("Failed to save user to database: %v\n", err)
		return
	}

	if !success {
		http.Error(w, "User already exists", http.StatusBadRequest)
		fmt.Printf("User already exists\n")
		return
	}

	fmt.Printf("Student added successfully: %s %s.\n", student.Firstname, student.Lastname)
}

//url: /signin
func SigninHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	//It returns an encrypted token if sign in succeeds.
	//It returns status unauthorized if username or password is incorrect.
	//It could also return an status internal server error if an error occurs that is caused by the database operations or the token encryption.

	decoder := json.NewDecoder(r.Body)
	var user User
	if err := decoder.Decode(&user); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode user data from client: %v\n", err)
		return
	}
	exists, err := checkUser(user.Email, user.Password)

	if err != nil {
		http.Error(w, "Failed to read user from db", http.StatusInternalServerError)
		fmt.Printf("Failed to read user from db: %v\n", err)
		return
	}
	if !exists {
		http.Error(w, "User does not exist or password is incorrect", http.StatusUnauthorized)
		fmt.Printf("User does not exist or password is incorrect\n")
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		fmt.Printf("Failed to generate token: %v\n", err)
		return
	}

	w.Write([]byte(tokenString))
}
