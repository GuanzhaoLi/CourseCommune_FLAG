package main

import (
	"database/sql"
	"fmt"
)

var db sql.DB

func checkUser(username, password string) (bool, error) {

	db, err1 := sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/QuestionOrder")
	if err1 !=nil {
		return false, err1
	}

	rows,err := db.Query("SELECT * FROM User")
	if err != nil {
		return false, err
	}
	var s User
	for rows.Next() {
		err := rows.Scan(&s.Username,&s.Password,&s.Uid,&s.IsTutor)
		if err != nil {
			fmt.Println(err)
		}
	}

	if username == s.Username && password == s.Password {
		return true, nil
	}
	rows.Close()
	return false, nil
}

func addUser(user *User) (bool, error) {

	// initialize database


	db, err1 := sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/User")
	if err1 !=nil {
		return false, err1
	}

	rows, err2 := db.Query("SELECT * FROM User")
	if err2 != nil {
		return false, err2
	}

	for rows.Next() {
		var s User
		err := rows.Scan(&s.Username,&s.Password,&s.Uid,&s.IsTutor)
		if err != nil{
			fmt.Println(err)
		}
		fmt.Println(s.Username)
		if user.Username != s.Username{
			result,err := db.Exec("INSERT INTO User (username, password, Id, isTutor) VALUES (?,?,?,?);",user.Username,user.Password,user.Uid, user.IsTutor)
			if err != nil {
				return false,err
			} else {
				rows,_ := result.RowsAffected()
				if rows != 1 {
					return false, nil
				}
			}
		} else {
			return false,nil
		}
	}
	rows.Close()

	fmt.Printf("User is added: %s\n", user.Username)
	return true, nil
}