package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_ "github.com/gorilla/mux"
	"net/http"
	"time"
)

var mySigningKey = []byte("secret")

func studentHistory(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var s StudentHistory
	//解析json
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		fmt.Println(2)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//function
	res :=studentHis(s)

	w.Header().Set("Content-Type", "application/json")
	//转回json
	er := json.NewEncoder(w).Encode(&res)
	if er != nil {
		fmt.Println(3)
		http.Error(w, er.Error(), http.StatusAccepted)
		return
	}

}

func tutorHistory(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var t TutorHistory
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := toturHis(t)

	w.Header().Set("Content-Type", "application/json")

	er := json.NewEncoder(w).Encode(&res)
	if er != nil {
		http.Error(w, er.Error(), http.StatusAccepted)
		return
	}
}
func tutorSignupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == "OPTIONS" {
		return
	}

	decoder := json.NewDecoder(r.Body)
	var tutor Tutor
	if err := decoder.Decode(&tutor); err != nil {
		http.Error(w, "Cannot decode user data from client", http.StatusBadRequest)
		fmt.Printf("Cannot decode user data from client: %v\n", err)
	}

	if tutor.Firstname == "" || tutor.Lastname == "" || tutor.Email == "" {
		http.Error(w, "Invalid username or password", http.StatusBadRequest)
		fmt.Printf("Invalid username or password\n")
		return
	}

	success, err := addTutor(&tutor)

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

	fmt.Printf("Tutor added successfully: %s %s.\n", tutor.Firstname, tutor.Lastname)
}

// signin sign up handler:
//url: /signup
func studentSignupHandler(w http.ResponseWriter, r *http.Request) {

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
	var user Users
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
		"username": user.Email,
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

func studentPostQuestion(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	fmt.Println("0")
	var q Questionn
	err1 := json.NewDecoder(r.Body).Decode(&q)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	var qr QuestionOrder
	qr.RequestBy = q.StudentId
	qr.Level = q.Level
	qr.Subject = q.Subject
	qr.Keywords = q.Keywords
	qr.StartTime = time.Now()
	qr.Answer = ""

	// post question to database
	err2 := postQToDB(&qr)
	if (err2 != nil) {
		http.Error(w, "failed to post question", http.StatusInternalServerError) // FY, error type
		return
	}
}

func studentSearchQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")	var k Keywordst
	err := json.NewDecoder(r.Body).Decode(&k)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	questions, err0 := findQuestions(&k)
	if (err0 != nil) {
		http.Error(w, "Failed to read questions from mySQL", http.StatusInternalServerError)
	}

	// encode to JSON
	err1 := json.NewEncoder(w).Encode(&questions) // FY, Encoder vs Marshal
	if (err1 != nil) {
		http.Error(w, "Failed to read questions from mySQL", http.StatusInternalServerError)
	}
}

func studentRequestVideo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")	fmt.Println("0")
	var vr VideoOrder
	err1 := json.NewDecoder(r.Body).Decode(&vr)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("1")
	fmt.Println(vr.StartTime)
	// post video request to database
	err2 := postVrToDB(&vr)
	if (err2 != nil) {
		fmt.Println("4")
		http.Error(w, "failed to send video request", http.StatusInternalServerError) // FY, error type
		return
	}
}

func studentSearchTutors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")	var c Criteria
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	tutors, err0 := findTutors(&c)
	if (err0 != nil) {
		http.Error(w, "Failed to read questions from mySQL", http.StatusInternalServerError)
	}
	// encode to JSON
	err1 := json.NewEncoder(w).Encode(&tutors) // FY, Encoder vs Marshal
	if (err1 != nil) {
		http.Error(w, "Failed to read tutors from mySQL", http.StatusInternalServerError)
	}
}
func tutorPickQuestion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var question QuestionOrder
	var pickQ PickQuestion
	err1 := json.NewDecoder(r.Body).Decode(&pickQ)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("decode to struct")
	question.QId = pickQ.QId
	question.FulfilledBy = pickQ.FulfilledBy
	question.Answer = pickQ.Answer
	err2 := updateQuestionToDB(&question)
	if (err2 != nil) {
		http.Error(w, "failed to update answer to database", http.StatusInternalServerError)
	} else {
		fmt.Println("successfully accessed database")
	}
}

func tutorPickVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	type PickedVideo struct {
		VideoId int64
		TutorId int64
		Agreed int64
	}

	var pv PickedVideo
	err1 := json.NewDecoder(r.Body).Decode(&pv)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusBadRequest)
		return
	}
	var vt VideoOrder
	vt.OrderId = pv.VideoId
	vt.FulfilledBy = pv.TutorId
	vt.Agreed = pv.Agreed

	err2 := updateVideoToDB(&vt)
	if (err2 != nil) {
		http.Error(w, "failed to update video request", http.StatusInternalServerError)
	}
}

func tutorSearchQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var c Tcriteria
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	questions, err1 := findQuestionsToAnswer(&c, )
	if (err1 != nil) {
		panic(err1)
	}

	er := json.NewEncoder(w).Encode(&questions)
	if (er != nil) {
		http.Error(w, "Failed to read questions from mySQL", http.StatusInternalServerError)
	}
}


func tutorSearchVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var c Tcriteria
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vts, err1 := findVideoTasks(&c)
	if (err1 != nil) {
		panic(err1)
	}
	er := json.NewEncoder(w).Encode(&vts)
	if (er != nil) {
		http.Error(w, "Failed to find video from mySQL", http.StatusInternalServerError)
	}
}
