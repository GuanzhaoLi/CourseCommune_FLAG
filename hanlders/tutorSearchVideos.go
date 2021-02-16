package main

import (
	"encoding/json"
	"net/http"
    _ "github.com/go-sql-driver/mysql"
)


func tutorSearchVideos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 

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

func findVideoTasks(c *Tcriteria) ([]VideoOrder, error) {
	// search for questions
	results, err2 := DB.Query("select OrderId, StartTime, RequestBy, Keywords from VideoOrder where Level = ? AND Subject = ? AND FulfilledBy = ?", c.Level, c.Subject, c.TutorId) 

	if err2 != nil {
		panic(err2.Error())
	}
	// get all VideoOrder searched
	var vts []VideoOrder
	for results.Next() {
		var vt VideoOrder
		err3 := results.Scan(&vt.OrderId, &vt.StartTime, &vt.RequestBy, &vt.Keywords)
		if (err3 != nil) {
			panic(err3)
		}
		vts = append(vts, vt)
	}
	return vts, nil
}