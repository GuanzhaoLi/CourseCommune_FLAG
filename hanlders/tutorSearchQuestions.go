package main

import (
	"encoding/json"
	"net/http"
    _ "github.com/go-sql-driver/mysql"
	"fmt"
)

type Tcriteria struct {
	TutorId int64
	Level int64
	Subject string
}

func tutorSearchQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 

	
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

func findQuestionsToAnswer(c *Tcriteria) ([]QuestionOrder, error) {
	// search for questions
	_, err2 := DB.Exec("update QuestionOrder set Answer = '' where Answer is null")
	results, err2 := DB.Query("select QId, StartTime, RequestBy, Keywords, Answer from QuestionOrder where Level = ? AND Subject = ? AND (FulfilledBy is null OR FulfilledBy = ?)", c.Level, c.Subject, c.TutorId)
	if err2 != nil {
		panic(err2.Error())
	}
	// get all questions searched
	var questions []QuestionOrder
	for results.Next() {
		var q QuestionOrder
		err3 := results.Scan(&q.QId, &q.StartTime, &q.RequestBy, &q.Keywords, &q.Answer)
		if (err3 != nil) {
			panic(err3)
		}
		fmt.Println(q.QId)
		questions = append(questions, q)
	}
	return questions, nil
}