package main

type AnswerQuestionRequestBody struct {
	questionId int `json:"questionId"`
	answer string `json:"answer"`
}