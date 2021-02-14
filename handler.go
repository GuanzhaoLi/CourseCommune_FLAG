package main

import (
	"encoding/json"
	_ "fmt"
	"net/http"
	_ "github.com/gorilla/mux"
)


func studentHistory(w http.ResponseWriter, r *http.Request)  {
	var s StudentHistory
	//解析json
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//function
	studentHis(s)

	w.Header().Set("Content-Type", "application/json")

	//转回json
	er := json.NewEncoder(w).Encode(&studentres)
	if er != nil {
		http.Error(w, er.Error(), http.StatusAccepted)
		return
	}
	//清空列表
	studentres = nil

}

func tutorHistory(w http.ResponseWriter, r *http.Request){
	var t TutorHistory
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	toturHis(t)

	w.Header().Set("Content-Type", "application/json")

	er := json.NewEncoder(w).Encode(&tutors)
	if er != nil {
		http.Error(w, er.Error(), http.StatusAccepted)
		return
	}
	tutors = nil
}
