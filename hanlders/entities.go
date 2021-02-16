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
type keywordst struct {
	Keywords	string
}