package main

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
