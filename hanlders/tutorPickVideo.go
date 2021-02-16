package main

import (
	"encoding/json"
	"net/http"
    _ "github.com/go-sql-driver/mysql"
	"time"
)

func tutorPickVideo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 

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
		http.Error(w, "failed to send video request", http.StatusInternalServerError)
	}
}

func updateVideoToDB(vt *VideoOrder) error {
	_, err2 := DB.Exec("update VideoOrder set EndTime = ?, FulfilledBy = ?, Agreed = ? where OrderId = ?", time.Now(), vt.FulfilledBy, vt.Agreed, vt.OrderId)
	return err2
}