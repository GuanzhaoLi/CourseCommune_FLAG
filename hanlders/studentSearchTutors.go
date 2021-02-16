package main

import (
	"encoding/json"
	// "strconv"
	"net/http"
    _ "github.com/go-sql-driver/mysql"
	
)

type Tutorr struct { // 内部的东西需要开头大写吗。其他api也用了 Tutor struct，但不一样。FY
	TutorId     int64 `json:"tutorId"`
	FirstName   string `json:"firstName"`
	LastName   	string `json:"lastName"`
	Rating		int64 `json:"rating"`
}


// receive JSON, return JSON
func studentSearchTutors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 
	var c Criteria
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

func findTutors(c *Criteria) ([]Tutorr, error) {
	results, err2 := DB.Query("select TutorId, FirstName, LastName, Rating from Tutor where Level = ? and Subject = ?", c.Level, c.Subject)  // SQL 语句 FY
	if err2 != nil {
		panic(err2.Error())
	}
	// get all tutors searched from database
	var tutors []Tutorr
	for results.Next() {
		var t Tutorr
		err3 := results.Scan(&t.TutorId, &t.FirstName, &t.LastName, &t.Rating)

		if (err3 != nil) {
			panic(err3) // error 处理。 branch中 print 获取失败
		}

		tutors = append(tutors, t)
	}
	return tutors, nil
}