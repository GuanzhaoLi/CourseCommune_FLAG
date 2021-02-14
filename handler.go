package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/gorilla/mux"
)


func studentHistory(w http.ResponseWriter, r *http.Request)  {
	var s StudentHistory
	//解析json
	err := json.NewDecoder(r.Body).Decode(&s)
	if err != nil {
		fmt.Println(2)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//res 接收return
	res :=studentHis(s)

	w.Header().Set("Content-Type", "application/json")
	//转回json
	er := json.NewEncoder(w).Encode(&res)
	if er != nil {
		fmt.Println(3)
		http.Error(w, er.Error(), http.StatusAccepted)
		return
	}

}

func tutorHistory(w http.ResponseWriter, r *http.Request){
	var t TutorHistory
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res := toturHis(t)

	w.Header().Set("Content-Type", "application/json")

	er := json.NewEncoder(w).Encode(&res)
	if er != nil {
		http.Error(w, er.Error(), http.StatusAccepted)
		return
	}
}
