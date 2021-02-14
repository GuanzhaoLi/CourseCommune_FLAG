package main

type Student struct {
	Id               string `json:"id"`
	Firstname        string `json:"firstname"`
	Lastname         string `json:"lastname"`
	Level            string `json:"level"`
	Rating           string `json:"rating"`
	AccountBalance   string `json:"account_balance"`
	UserId           string `json:"user_id"`
	StudentHistoryId string `json:"student_history_id"`
	TutorHistoryId   string `json:"tutor_history_id"`
}
