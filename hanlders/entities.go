package main

import "time"

type TutorHistory struct {
	Id      int64 `json:"tutorId"`
	OrderId int64 `jason:"VideoOrder"`
	Qid     int64 `json:"QuestionOrder"`
}

type StudentHistory struct {
	Id      int64 `json:"studentId"`
	OrderId int64 `jason:"VideoOrder"`
	Qid     int64 `json:"QuestionOrder"`
}

type Users struct {
	Email string `json:"email"`
	Password string `json:"password"`
}
type AnswerQuestionRequestBody struct {
	questionId int `json:"questionId"`
	answer string `json:"answer"`
}
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
	Keywords	string `jason:"keywords"`
}
type Questionn struct {
	StudentId 	int64 `jason:"studentId"`
	Level 		int64	`json:"level"`
	Subject 	string	`jason:"subject"`
	Keywords 	string	`json:"keywords"`
}

type Video struct {
	Id          string `jason:"Id"`
	User        string `jason:"User"`
	Date        string `jason:"Date"`
	Subject     string `jason:"Subject"`
	Description string `jason:"Description"`
	Fulefilled  string `jason:"Fulefilled"`
}

type Student struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Level     string `json:"level"`
	// Rating           string `json:"rating"`
	// AccountBalance   string `json:"account_balance"`
	// UserId           string `json:"user_id"`
	// StudentHistoryId string `json:"student_history_id"`
	// TutorHistoryId   string `json:"tutor_history_id"`
	// isTutor bool `json:"isTutor"` //does Go have enums?
}

type Tutor struct {
	Id        string `json:"id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Level     string `json:"level"`
	Subject   string `json:"subject"`
	// Account_Balance int,
	// Rating int,
}
type Tutorr struct { // 内部的东西需要开头大写吗。其他api也用了 Tutor struct，但不一样。FY
	TutorId     int64 `json:"tutorId"`
	FirstName   string `json:"firstName"`
	LastName   	string `json:"lastName"`
	Rating		int64 `json:"rating"`
}

type PickQuestion struct {
	QId int64
	FulfilledBy int64
	Answer string
}

type Tcriteria struct {
	TutorId int64
	Level int64
	Subject string
}
