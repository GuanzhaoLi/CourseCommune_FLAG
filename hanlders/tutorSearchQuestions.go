package main

import (
	"encoding/json"
	"net/http"
    _ "github.com/go-sql-driver/mysql"
)


func tutorSearchQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 

	var c Criteria
	err := json.NewDecoder(r.Body).Decode(&c)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return
    }
	questions, err1 := findQuestionsToAnswer(&c)
	if (err1 != nil) {
		panic(err1)
	}
	
	er := json.NewEncoder(w).Encode(&questions)
	if (er != nil) {
		http.Error(w, "Failed to read questions from mySQL", http.StatusInternalServerError)
	}
}

func findQuestionsToAnswer(c *Criteria) ([]QuestionOrder, error) {
	// search for questions
	results, err2 := DB.Query("select QId, StartTime, RequestBy, Keywords from QuestionOrder where Level = ? AND Subject = ? AND FulfilledBy is null", c.Level, c.Subject) 

	if err2 != nil {
		panic(err2.Error())
	}
	// get all questions searched
	var questions []QuestionOrder
	for results.Next() {
		var q QuestionOrder
		err3 := results.Scan(&q.QId, &q.StartTime, &q.RequestBy, &q.Keywords)
		if (err3 != nil) {
			panic(err3)
		}
		questions = append(questions, q)
	}
	return questions, nil
}