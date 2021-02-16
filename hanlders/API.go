package main

import (
	"database/sql"
	"fmt"
	"time"
)

func checkUser(email string, password string) (bool, error) {
	InitDB()
	var fetched_password string
	err := DB.QueryRow("select password from Users where email = ?", email).Scan(&fetched_password)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
	}
	if password == fetched_password {
		return true, nil
	} else {
		return false, nil
	}
}
//
func addTutor(tutor *Tutor) (bool, error) {
	InitDB()
	email := tutor.Email
	rows, err := DB.Query("select email from Users where email = ?", email)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if hasnext := rows.Next(); hasnext {
		return false, nil
	}
	id := tutor.Id
	fn := tutor.Firstname
	ln := tutor.Lastname
	lev := tutor.Level
	pass := tutor.Password
	subject := tutor.Subject
	sql := "INSERT INTO Tutor (TutorId, FirstName, LastName, Level, Subject) VALUES (?,?,?,?,?)"
	_, err = DB.Exec(sql, id, fn, ln, lev, subject)
	sql = "INSERT INTO Users (id, password, email) VALUES (?,?,?)"
	_, err = DB.Exec(sql, id, pass, email)
	if err != nil {
		return false, err
	}
	return true, nil
}

//should id be auto-incremented?
func addStudent(student *Student) (bool, error) {
	InitDB()
	email := student.Email
	rows, err := DB.Query("select email from Users where email = ?", email)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	if hasnext := rows.Next(); hasnext {
		return false, nil
	}
	id := student.Id
	fn := student.Firstname
	ln := student.Lastname
	lev := student.Level
	pass := student.Password
	sql := "INSERT INTO student(StudentId, FirstName, LastName, Level) VALUES (?,?,?,?)"
	_, err = DB.Exec(sql, id, fn, ln, lev)
	sql = "INSERT INTO Users (id, password, email) VALUES (?,?,?)"
	_, err = DB.Exec(sql, id, pass, email)
	if err != nil {
		return false, err
	}
	return true, nil
}
//
func postQToDB(qr *QuestionOrder) error {
	_, err := DB.Exec("insert into QuestionOrder values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", nil, qr.StartTime, qr.EndTime, qr.RequestBy, nil, qr.Subject, qr.Level, qr.Keywords, qr.Answer, qr.S_t_rating)
	return err
}
//
func findQuestions(k *Keywordst) ([]QuestionOrder, error) {
	results, err2 := DB.Query("select QId, StartTime, RequestBy, Subject, Level, Keywords, Answer, S_t_rating from QuestionOrder where keywords = ?", k.Keywords)  // SQL 语句 FY
	if err2 != nil {
		panic(err2.Error())
	}
	// get all questions searched from database
	var questions []QuestionOrder
	for results.Next() {
		var t QuestionOrder
		err3 := results.Scan(&t.QId, &t.StartTime, &t.RequestBy, &t.Subject,&t.Level, &t.Keywords, &t.Answer, &t.S_t_rating)

		if (err3 != nil) {
			panic(err3) // error 处理。 branch中 print 获取失败
		}

		questions = append(questions, t)
	}
	return questions, nil
}
//
func postVrToDB(vr *VideoOrder) error {
	_, err2 := DB.Exec("insert into VideoOrder values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", nil, vr.StartTime, vr.EndTime, vr.RequestBy, vr.FulfilledBy, vr.Duration, vr.Subject, vr.Level, vr.Keywords, nil, vr.T_s_rating, vr.S_t_rating)

	return err2
}
//
func findTutors(c *Criteria) ([]Tutorr, error) {
	results, err2 := DB.Query("select TutorId, FirstName, LastName, Rating from Tutor where Level = ? and Subject = ?", c.Level, c.Subject)  // SQL 语句 FY
	if err2 != nil {
		panic(err2.Error())
	}
	// get all tutors searched from database
	var tutors []Tutorr
	for results.Next() {
		var t Tutorr
		err3 := results.Scan(&t.TutorId, &t.FirstName, &t.LastName, &t.Rating)

		if (err3 != nil) {
			panic(err3) // error 处理。 branch中 print 获取失败
		}

		tutors = append(tutors, t)
	}
	return tutors, nil
}
//
func updateQuestionToDB(q *QuestionOrder) error {
	q.EndTime = time.Now()
	_, err2 := DB.Exec("update QuestionOrder set FulfilledBy = ?, EndTime = ?, Answer = ? where QId = ?", q.FulfilledBy, q.EndTime, q.Answer, q.QId) // 语句 FY
	// unsolved, table need to add Answer column
	return err2
}
//
func updateVideoToDB(vt *VideoOrder) error {
	_, err2 := DB.Exec("update VideoOrder set EndTime = ?, FulfilledBy = ?, Agreed = ? where OrderId = ?", time.Now(), vt.FulfilledBy, vt.Agreed, vt.OrderId)
	return err2
}
//
func findQuestionsToAnswer(c *Tcriteria) ([]QuestionOrder, error) {
	// search for questions
	// _, err2 := DB.Exec("update QuestionOrder set Answer = '' where Answer is null")
	results, err2 := DB.Query("select QId, StartTime, RequestBy, Keywords, Answer from QuestionOrder where Level = ? AND Subject = ? AND (FulfilledBy is null OR FulfilledBy = ?)", c.Level, c.Subject, c.TutorId)
	if err2 != nil {
		panic(err2.Error())
	}
	// get all questions searched
	var questions []QuestionOrder
	for results.Next() {
		var q QuestionOrder
		err3 := results.Scan(&q.QId, &q.StartTime, &q.RequestBy, &q.Keywords, &q.Answer)
		if (err3 != nil) {
			panic(err3)
		}
		fmt.Println(q.QId)
		questions = append(questions, q)
	}
	return questions, nil
}
//
func findVideoTasks(c *Tcriteria) ([]VideoOrder, error) {
	// search for questions
	results, err2 := DB.Query("select OrderId, StartTime, RequestBy, Keywords from VideoOrder where Level = ? AND Subject = ? AND FulfilledBy = ?", c.Level, c.Subject, c.TutorId)

	if err2 != nil {
		panic(err2.Error())
	}
	// get all VideoOrder searched
	var vts []VideoOrder
	for results.Next() {
		var vt VideoOrder
		err3 := results.Scan(&vt.OrderId, &vt.StartTime, &vt.RequestBy, &vt.Keywords)
		if (err3 != nil) {
			panic(err3)
		}
		vts = append(vts, vt)
	}
	return vts, nil
}
//Student history
func studentHis(studenthis StudentHistory) interface{}{
	InitDB()
	var target = studenthis.Id
	var studentres = make([]StudentHistory, 0)

	rows, err := DB.Query("select SH.id, SH.VideoOrder, SH.QuestionOrder from Student_History SH join Student ST on SH.id = ST.StudentId where SH.id = ?", target)
	if err != nil {
		fmt.Println(err.Error())
	}

	for rows.Next() {
		var id,videoid,questionid int64
		if e := rows.Scan(&id, &videoid, &questionid);e != nil {
			fmt.Println("获取失败。。")
		}
		comment := StudentHistory{Id: id, OrderId: videoid, Qid: questionid}
		studentres = append(studentres,comment)
	}
	defer rows.Close()
	defer DB.Close()
	return studentres
}

func toturHis(tutorhis TutorHistory) interface{}{
	InitDB()
	var find = tutorhis.Id
	var tutors = make([]TutorHistory, 0)
	rows, err := DB.Query("select TH.id, TH.VideoOrder, TH.QuestionOrder from Tutor_History TH join Tutor TU on TH.id = TU.TutorId where TH.id = ?" ,find)
	if err != nil {
		fmt.Println(1)
	}

	for rows.Next() {
		var id,videoid,questionid int64
		if e := rows.Scan(&id, &videoid, &questionid);e != nil {
			fmt.Println("获取失败。。")
		}
		comment := TutorHistory{Id: id, OrderId: videoid, Qid: questionid}
		tutors = append(tutors,comment)
	}
	defer rows.Close()
	defer DB.Close()
	return tutors
}