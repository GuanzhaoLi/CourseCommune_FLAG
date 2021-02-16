package main

import (
	"encoding/json"
	"net/http"
	"time"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)


func studentPostQuestion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") 
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

func postQToDB(qr *QuestionOrder) error {
	_, err := DB.Exec("insert into QuestionOrder values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", nil, qr.StartTime, nil, qr.RequestBy, nil, qr.Subject, qr.Level, qr.Keywords, qr.Answer, qr.S_t_rating)
	return err
}

//receive keyword return []JSON
func studentSearchQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 
	var k Keywordst
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

func findQuestions(k *Keywordst) ([]QuestionOrder, error) {
	results, err2 := DB.Query("select QId, StartTime, EndTime, RequestBy, FulfilledBy, Subject, Level, Keywords, Answer, S_t_rating from QuestionOrder where keywords = ?", k.Keywords)  // SQL 语句 FY
	if err2 != nil {
		panic(err2.Error())
	}
	// get all questions searched from database
	var questions []QuestionOrder
	for results.Next() {
		var t QuestionOrder
		err3 := results.Scan(&t.QId, &t.StartTime, &t.EndTime, &t.RequestBy, &t.FulfilledBy, &t.Subject,&t.Level, &t.Keywords, &t.Answer, &t.S_t_rating)

		if (err3 != nil) {
			panic(err3) // error 处理。 branch中 print 获取失败
		}

		questions = append(questions, t)
	}
	return questions, nil
}