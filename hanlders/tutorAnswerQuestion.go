package main

import (
	"time"
	"encoding/json"
	"net/http"
    _ "github.com/go-sql-driver/mysql"
	"fmt"
)

type PickQuestion struct {
	QId int64
	FulfilledBy int64
	Answer string
}

func tutorPickQuestion(w http.ResponseWriter, r *http.Request) {
	
	w.Header().Set("Content-Type", "application/json") 

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
		http.Error(w, "failed to send video request", http.StatusInternalServerError)
	} else {
		fmt.Println("successfully accessed database")
	}
}

func updateQuestionToDB(q *QuestionOrder) error {
	q.EndTime = time.Now()
	_, err2 := DB.Exec("update QuestionOrder set FulfilledBy = ?, EndTime = ?, Answer = ? where QId = ?", q.FulfilledBy, q.EndTime, q.Answer, q.QId) // 语句 FY
	// unsolved, table need to add Answer column
	return err2
}