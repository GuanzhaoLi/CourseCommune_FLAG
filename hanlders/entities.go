package main

import (
	"time"
)

type VideoOrder struct {
	OrderId		int64 `json:"orderId"`
	StartTime	time.Time `json:"startTime"`
	EndTime		time.Time `json:"endTime"`
	RequestBy	int64 `json:"requestedBy"`
	FulfilledBy	int64 `json:"fulfilledBy"`
	Duration	int64 `json:"duration"`	
	Subject		string `json:"subject"`
	Level		int64 `json:"level"`
	Keywords	string `json:"keywords"`
	Agreed		int64 `json:"agreed"`
	T_s_rating	int64 `json:"t_s_rating"`  // tutor gives rate on student
	S_t_rating	int64 `json:"s_t_rating"`	
}

type QuestionOrder struct {
	QId			int64 `json:"QId"`
	StartTime	time.Time `json:"startTime"`
	EndTime		time.Time `json:"endTime"`
	RequestBy	int64 `json:"requestBy"`
	FulfilledBy	int64 `json:"fulfilledBy"`	
	Subject		string `json:"subject"`
	Level		int64 `json:"level"`
	Keywords	string `json:"keywords"`
	Answer      string `json:"answer"`
	S_t_rating	int64 `json:"s_t_rating"`
}

type Criteria struct {
	Level     int64 
	Subject   string 
}
type Keywordst struct {
	Keywords	string
}
type Questionn struct {
	StudentId int64
	Level int64
	Subject string
	Keywords string
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsTutor bool `json:"isTutor"` //does Go have enums?
}

type Video struct {
	Id          string `jason:"Id"`
	User        string `jason:"User"`
	Date        string `jason:"Date"`
	Subject     string `jason:"Subject"`
	Description string `jason:"Description"`
	Fulefilled  string `jason:"Fulefilled"`
}
