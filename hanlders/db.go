package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

//数据库配置
const (
	userName = "root"
	// password = "自己密码"
	password = "rootroot"
	ip       = "127.0.0.1"
	port     = "3306"
	// dbName   = "qa1"
	dbName = "flagcamp"
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

//初始化表单
func initialTable() {
	{ // Create a new table
		query := `
		CREATE TABLE IF NOT EXISTS VideoOrder(
			OrderId int AUTO_INCREMENT, 
			StartTime DateTime, EndTime DateTime, 
			RequestBy int, FulfilledBy int, Duration Integer,
			Subject varchar(50), Level int, Keywords varchar(256), 
			Agreed int, T_s_rating int, S_t_rating int, 
			Primary Key(OrderId));
			`
		if _, err := DB.Exec(query); err != nil {
			log.Fatal(err)
		}
	}

	{
		query := `
		CREATE TABLE IF NOT EXISTS QuestionOrder(
			QId int AUTO_INCREMENT,
			StartTime DateTime, EndTime DateTime, 
			RequestBy int, FulfilledBy int,
			Subject varchar(50), Level int, Keywords varchar(256), 
			Answer varchar(256), S_t_rating int, 
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
			TutorId int AUTO_INCREMENT,
			Level int,
			Subject varchar(50),
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
			StudentId int AUTO_INCREMENT, 
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
			id int AUTO_INCREMENT, 
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
			id int AUTO_INCREMENT, 
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
			id int AUTO_INCREMENT, 
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
