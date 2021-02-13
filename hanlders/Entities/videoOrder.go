package main

type VideoOrder struct {
	TutorID		int64 `json:"tutorID"`
	StudentID	int64 `json:"studentID"`
	StartTime	string `json:"startTime"`
	EndTime		string `json:"endTime"`
	Level		int64 `json:"level"`
	Subject		string `json:"subject"`
	Keywords	string `json:"keywords"`
	T_s_rating	int64 `json:"t_s_rating"`  // tutor gives rate on student
	T_t_rating	int64 `json:"s_t_rating"`  // student gives rate on tutor
	Finished	bool `json:"finished"`
}
