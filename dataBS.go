package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strings"
	"time"
)
//数据库配置
const (
	userName = "root"
	password = "自己密码"
	ip       = "127.0.0.1"
	port     = "3306"
	dbName   = "QA"
)

//Db数据库连接池
var DB *sql.DB

//初始化数据库
func InitDB() {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?parseTime=true&multiStatements=true"}, "")
	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	DB, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	DB.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	DB.SetMaxIdleConns(10)
	//
	DB.SetConnMaxLifetime(time.Minute * 3)
	// 尝试与数据库建立连接（校验dsn是否正确）
	if err := DB.Ping(); err != nil {
		fmt.Println("open database fail")
		return
	}
	fmt.Println("connnect success")
	initialTable()
}
var studentres = make([]StudentHistory, 0)

//Student history
func studentHis(studenthis StudentHistory) {
	var target = studenthis.Id
	//studentres := make([] StudentHistory, 0)

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
	//rows.Close()
	//DB.Close()

}
var tutors = make([]TutorHistory, 0)

func toturHis(tutorhis TutorHistory){
	var find = tutorhis.Id
	rows, err := DB.Query("select TH.id, TH.VideoOrder, TH.QuestionOrder from Tutor_History TH join Tutor TU on TH.id = TU.TutorId where TH.id = ?" ,find)
	if err != nil {
		fmt.Println(err.Error())
	}
	for rows.Next() {
		var id,videoid,questionid int64
		if e := rows.Scan(&id, &videoid, &questionid);e != nil {
			fmt.Println("获取失败。。")
		}
		comment := TutorHistory{Id: id, OrderId: videoid, Qid: questionid}
		tutors = append(tutors,comment)
	}
// 	rows.Close()
// 	DB.Close()
}
//初始化表单
func initialTable(){
	{ // Create a new table
		query := `
		CREATE TABLE IF NOT EXISTS VideoOrder(
			OrderId int, 
			StartTime DateTime, EndTime DateTime, 
			RequestBy int, FulfilledBy int, Duration Integer,
			Subject varchar(50), Level int, Keywords varchar(256), 
			UsersId int,
			Primary Key(OrderId));
			`
		if _, err := DB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		query := `
		CREATE TABLE IF NOT EXISTS QuestionOrder(
			QId int,
			StartTime DateTime, EndTime DateTime, 
			RequestBy int, FulfilledBy int,
			Subject varchar(50), Level int, Keywords varchar(256), 
			UserId int,
			Primary Key(QId)
			);
		`
		if _, err := DB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		query := `
		CREATE TABLE IF NOT EXISTS Tutor(
			TutorId int,
			Level int,
			Subject int,
			FirstName varchar(50),
			LastName varchar(50),
			Account_Balance int,
			Rating int,
			Primary Key(TutorId)
		);
		`
		if _, err := DB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		query := `
		CREATE TABLE IF NOT EXISTS Student(
			StudentId int, 
			FirstName varchar(50),
			LastName varchar(50),
			Level int,
			Account_Balance int,
			Rating int,
			Primary Key(StudentId)
		);
		`
		if _, err := DB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		query := `
		CREATE TABLE IF NOT EXISTS Users(
			Id int,
			email varchar(256),
			password varchar(32),
			birthday date,
			StudentId int,
			TutorId int,
			Foreign Key(StudentId) References Student(StudentId) ON DELETE CASCADE ON Update CASCADE,
			Foreign Key(TutorId) References Tutor(TutorId) ON DELETE CASCADE ON Update CASCADE
		);
		`
		if _, err := DB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		query := `
		CREATE TABLE IF NOT EXISTS Student_Preference(
			id int, 
			Prefered_tutor int,
			Comments varchar(256),
			Foreign Key(id) References Student(StudentId) ON DELETE CASCADE ON Update CASCADE
		);
		`
		if _, err := DB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		query := `
		CREATE TABLE IF NOT EXISTS Student_History(
			id int, 
			VideoOrder int,
			QuestionOrder int,
			Foreign Key(id) References Student(StudentId) ON DELETE CASCADE ON Update CASCADE
		);
		`
		if _, err := DB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		query := `
		CREATE TABLE IF NOT EXISTS Tutor_History(
			id int, 
			VideoOrder int,
			QuestionOrder int,
			Foreign Key(id) References Tutor(TutorId) ON DELETE CASCADE ON Update CASCADE
		);
		`
		if _, err := DB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}
}


