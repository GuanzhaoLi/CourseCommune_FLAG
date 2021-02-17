package main

import (
	"fmt"
	"log"
	"net/http"

<<<<<<< Updated upstream
	"github.com/gorilla/mux"
)

// let's declare a global Quetsions and videos array
// that we can then populate in our main function
// to simulate a database
var Questions []Question
var Videos []Video

func QuestionPage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: QuestionPage")
}
func handleRequestsOfQuestion() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", http.HandlerFunc(QuestionPage)).Methods("GET")
	myRouter.HandleFunc("/{user}/questions", returnAllQuestions).Methods("GET", "POST")
	myRouter.HandleFunc("/{user}/questions/{id}", returnSingleQuestion).Methods("GET")
	myRouter.HandleFunc("/{user}/video", returnAllVideo).Methods("GET", "POST")
	myRouter.HandleFunc("/{user}/video/{id}", returnVideo).Methods("GET")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))

	// jianheng added. 照葫芦画瓢
	myRouter.handleFunc("/{user}/tutor_search", searchTutor).Methods("GET")
	http.handleFunc("/{user}/schedule_meeing", requestVideo).Methods("POST")
=======
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	// jwt "github.com/dgrijalva/jwt-go"
	jwt "github.com/form3tech-oss/jwt-go"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func handler() {
	// creates a new instance of a mux router
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(mySigningKey), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})
	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/user/tuotorsignup", tutorSignupHandler).Methods("POST", "OPTIONS")
	// myRouter.HandleFunc("/user/studentsignup", studentSignupHandler).Methods("POST", "OPTIONS")
	// myRouter.HandleFunc("/user/signin", SigninHandler).Methods("POST", "OPTIONS")
	// myRouter.HandleFunc("/student={user}/history", studentHistory).Methods("GET", "OPTIONS")
	// myRouter.HandleFunc("/tutor={user}/history", tutorHistory).Methods("GET", "OPTIONS")
	// myRouter.HandleFunc("/student={user}/postquestion", studentPostQuestion).Methods("POST", "OPTIONS")
	// myRouter.HandleFunc("/student={user}/searchqustion", studentSearchQuestions).Methods("GET", "OPTIONS")
	// myRouter.HandleFunc("/student={user}/requestvideo", studentRequestVideo).Methods("POST", "OPTIONS")
	// myRouter.HandleFunc("/student={user}/searchtutor", studentSearchTutors).Methods("POST", "OPTIONS")
	// myRouter.HandleFunc("/tutor={user}/pickquestion", tutorPickQuestion).Methods("POST", "OPTIONS")
	// myRouter.HandleFunc("/tutor={user}/pickvideo", tutorPickVideo).Methods("POST", "OPTIONS")
	// myRouter.HandleFunc("/tutor={user}/searchquestions", tutorSearchQuestions).Methods("GET", "OPTIONS")
	// myRouter.HandleFunc("/tutor={user}/searchvideos", tutorSearchVideos).Methods("GET", "OPTIONS")

	r := mux.NewRouter()
	r.Handle("/student={user}/history", jwtMiddleware.Handler(http.HandlerFunc(studentHistory))).Methods("GET", "OPTIONS")
	r.Handle("/tutor={user}/history", jwtMiddleware.Handler(http.HandlerFunc(tutorHistory))).Methods("GET", "OPTIONS")
	r.Handle("/student={user}/postquestion", jwtMiddleware.Handler(http.HandlerFunc(studentPostQuestion))).Methods("POST", "OPTIONS")
	r.Handle("/student={user}/searchqustion", jwtMiddleware.Handler(http.HandlerFunc(studentSearchQuestions))).Methods("GET", "OPTIONS")
	r.Handle("/student={user}/requestvideo", jwtMiddleware.Handler(http.HandlerFunc(studentRequestVideo))).Methods("POST", "OPTIONS")
	r.Handle("/student={user}/searchtutor", jwtMiddleware.Handler(http.HandlerFunc(studentSearchTutors))).Methods("POST", "OPTIONS")
	r.Handle("/tutor={user}/pickquestion", jwtMiddleware.Handler(http.HandlerFunc(tutorPickQuestion))).Methods("POST", "OPTIONS")
	r.Handle("/tutor={user}/pickvideo", jwtMiddleware.Handler(http.HandlerFunc(tutorPickVideo))).Methods("POST", "OPTIONS")
	r.Handle("/tutor={user}/searchquestions", jwtMiddleware.Handler(http.HandlerFunc(tutorSearchQuestions))).Methods("GET", "OPTIONS")
	r.Handle("/tutor={user}/searchvideos", jwtMiddleware.Handler(http.HandlerFunc(tutorSearchVideos))).Methods("GET", "OPTIONS")
	r.Handle("/tutorsignup", http.HandlerFunc(tutorSignupHandler)).Methods("POST", "OPTIONS")
	r.Handle("/studentsignup", http.HandlerFunc(studentSignupHandler)).Methods("POST", "OPTIONS")
	r.Handle("/signin", http.HandlerFunc(SigninHandler)).Methods("POST", "OPTIONS")

	// log.Fatal(http.ListenAndServe(":10000", myRouter))
	log.Fatal(http.ListenAndServe(":10000", r))
>>>>>>> Stashed changes
}

func main() {
	fmt.Println("Started")
	//Get Questions from database
	Questions = []Question{
		Question{Id: "1", Name: "Sin", Desc: "Article Description", Content: "Article Content"},
		Question{Id: "2", Name: "Lee", Desc: "Article Description", Content: "Article Content"},
	}
	Videos = []Video{
		Video{Id: "1", User: "a", Date: "02/28/1993", Subject: "Math", Description: "Math question", Fulefilled: "ture"},
		Video{Id: "2", User: "b", Date: "03/12/1993", Subject: "English", Description: "English q", Fulefilled: "False"},
	}
	handleRequestsOfQuestion()

}
