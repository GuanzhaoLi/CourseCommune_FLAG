package main

import (
	"net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

type Criteria struct {
	level     int64 `json:"level"`
	subject   string `json:"subject"`
}

func tutorSearchQuestions(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.body)
	w.Header().Set("Content-Type", "application/json") 
	w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type,Authorization")

	c := Criteria {
		level: 		r.formValue("level"),
		subject: 	r.formValue("subject"),
	}
	
	questions, err := findQuetions(&c)
	
	js, err := JSON.Marshal(questions)
	if (err != nil) {
		http.Error(w, "Failed to read questions from mySQL", http.StatusInternalServerError)
	}
	w.write(js)
}

func findQuestions(c *Criteria) []Question {
	db, err1 = sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/QuestionOrder") // need modify dataSourceName
	                                                                    // name, password, and IP address
	if err1 != nil {
		panic(err.Error())
	}
	defer db.Close()

	// search for questions
	results, err2 := db.Query("SELECT QId, RequestBy, FulfilledBy, Description FROM Tutor WHERE 
	Level = c.level AND Subject = c.subject AND Finished = false") 

	if err2 != nil {
		panic(err.Error())
	}
	// get all questions searched
	var questions []Question
	for results.Next() {
		var q Question
		err3 := results.Scan(&q.QId, &q.studentID, &q.TutorID, &q.Description)
		if (err3 != nil) {
			panic(err)
		}
		questions.append(questions, q)
	}
	return questions
}