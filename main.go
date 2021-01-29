package main

import (
       
)

func main() {
    
    http.handleFunc(“/question_search”, questionSearchHandler)
    http.handleFunc(“/question_post”, questionPostHandler)
	http.handleFunc(“/tutor_search”, tutorSearchHandler)
	http.handleFunc(“/schedule_meeing”, scheduleMeetingHandler)

}
