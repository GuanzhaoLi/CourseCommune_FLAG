package main

import (
	"encoding/json"
	"net/http"
    _ "github.com/go-sql-driver/mysql"
	"fmt"
	// "time"
)

func studentRequestVideo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") 
	fmt.Println("0")
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

func postVrToDB(vr *VideoOrder) error {
	_, err2 := DB.Exec("insert into VideoOrder values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", nil, vr.StartTime, vr.EndTime, vr.RequestBy, vr.FulfilledBy, vr.Duration, vr.Subject, vr.Level, vr.Keywords, nil, vr.T_s_rating, vr.S_t_rating)

	return err2
}