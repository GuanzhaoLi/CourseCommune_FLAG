package main

type VideoOrder struct {
	tutorID   	int64 `json:"tutorID"`
	studentID   int64 `json:"studentID"`
	startTime 	string `json:"startTime"`
	endTime	  	string `json:"endTime"`
	level		int64 `json:"level"`
	subject		string `json:"subject"`
	keywords	string `json:"keywords"`
	t_s_rating	int64 `json:"t_s_rating"`  // tutor gives rate on student
	s_t_rating	int64 `json:"s_t_rating"`  // student gives rate on tutor
	finished	bool `json:"finished"`
}