package main

import "database/sql"

const (
	DATABASE      = "root:rootroot@tcp(localhost:3306)/flagcamp?charset=utf8&parseTime=True&loc=Local"
	USER_TABLE    = "users"
	STUDENT_TABLE = "student"
	TUTOR_TABLE   = "Tutor"
)

type User struct {
	// Id       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
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

func checkUser(email string, password string) (bool, error) {
	db, err := sql.Open("mysql", DATABASE)
	if err != nil {
		return false, err
	}
	defer db.Close()
	var fetched_password string
	err = db.QueryRow("select password from "+USER_TABLE+" where email = ?", email).Scan(&fetched_password)
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

func addTutor(tutor *Tutor) (bool, error) {
	db, err := sql.Open("mysql", DATABASE)
	if err != nil {
		return false, err
	}
	defer db.Close()
	email := tutor.Email
	rows, err := db.Query("select email from "+USER_TABLE+" where email = ?", email)
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
	sql := "INSERT INTO " + TUTOR_TABLE + " (TutorId, FirstName, LastName, Level, Subject) VALUES (?,?,?,?,?)"
	_, err = db.Exec(sql, id, fn, ln, lev, subject)
	sql = "INSERT INTO " + USER_TABLE + " (id, password, email) VALUES (?,?,?)"
	_, err = db.Exec(sql, id, pass, email)
	if err != nil {
		return false, err
	}
	return true, nil
}

//should id be auto-incremented?
func addStudent(student *Student) (bool, error) {
	db, err := sql.Open("mysql", DATABASE)
	if err != nil {
		return false, err
	}
	defer db.Close()
	email := student.Email
	rows, err := db.Query("select email from "+USER_TABLE+" where email = ?", email)
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
	_, err = db.Exec(sql, id, fn, ln, lev)
	sql = "INSERT INTO " + USER_TABLE + "(id, password, email) VALUES (?,?,?)"
	_, err = db.Exec(sql, id, pass, email)
	if err != nil {
		return false, err
	}
	return true, nil
}
