package main

type QuestionOrder struct {
	QId			int64 `json:"description"`
	StudentID	int64 `json:"studentID"`  // requested by
	TutorID		int64 `json:"tutorID"`  // fullfilled by
	Level		int64 `json:"level"`
	Subject		string `json:"subject"`
	Description string `json:"description"`
	Answer      string `json:"answer"`
	Finished	bool `json:"finished"`
}