package main

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	isTutor bool `json:"isTutor"` //does Go have enums?
}