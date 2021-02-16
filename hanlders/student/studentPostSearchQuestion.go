import (
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"regexp"

	"time"
	_ "github.com/go-sql-driver/mysql"
	jwt "github.com/dgrijalva/jwt-go"
)


func studentPostQuestion(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json") 
	fmt.Println("0")
	var q question
	err1 := json.NewDecoder(r.Body).Decode(&q)
    if err1 != nil {
        http.Error(w, err1.Error(), http.StatusBadRequest) 
		return
    }
	fmt.Println("1")
	fmt.Println(q.StartTime)
	// post question to database
	err2 := postQToDB(&q)
	if (err2 != nil) {
		fmt.Println("4")
		http.Error(w, "failed to post question", http.StatusInternalServerError) // FY, error type
		return
	}
}

func postQToDB(vr *VideoOrder) error {
	_, err2 := DB.Exec("insert into QuestionOrder values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", nil, q.StartTime, q.EndTime, q.RequestBy, q.FulfilledBy, q.Subject, q.Level, q.Keywords, q.Answer, q.s_t_rating)

	return err2
}






//receive keyword return []JSON
func studentSearchQuestions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") 
	var k keywordst
	err := json.NewDecoder(r.Body).Decode(&k)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		return
    }
	questions, err0 := findQuestions(&k)
	if (err0 != nil) {
		http.Error(w, "Failed to read questions from mySQL", http.StatusInternalServerError)
	}

	// encode to JSON
	err1 := json.NewEncoder(w).Encode(&questions) // FY, Encoder vs Marshal
	if (err1 != nil) {
		http.Error(w, "Failed to read questions from mySQL", http.StatusInternalServerError)
	}
}

func findQuestions(k *keywords) ([]questionn, error) {
	results, err2 := DB.Query("select QId, StartTime, EndTime, RequestBy, FulfilledBy, Subject, Level, Keywords, Answer, S_t_rating from QuestionOrder where keywords = ?", k)  // SQL 语句 FY
	if err2 != nil {
		panic(err2.Error())
	}
	// get all questions searched from database
	var questions []questionn
	for results.Next() {
		var t quesitonn
		err3 := results.Scan(&t.QId, &t.StartTime, &t.EndTime, &t.RequestBy, &t.FulfilledBy, &t.Subject,&t.Level, &t.Keywords, &t.Answer, &t.t_s_rating)

		if (err3 != nil) {
			panic(err3) // error 处理。 branch中 print 获取失败
		}

		questions = append(questions, t)
	}
	return questions, nil
}