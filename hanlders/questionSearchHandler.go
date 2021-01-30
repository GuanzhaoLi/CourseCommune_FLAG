package main

import {

}

type Question struct {
	questionText, subject, level, // and more...
}

// receive JSON, return JSON
func questionSearchHandler(w http.ResponseWriter, r *http.Request) {
	// ...
	q := Question {
		questionText: r.formValue("questionText"),
		subject: r.formValue("subject"),
		level: r.formValue("level"), // high school? college junoir?
		// ...
	}
	// search for similar questions
	similarQuestions, err := searchSimilarQuestions(q.questionText)

	// JSON encode
	js, err := JSON.Marshal(similarQuestions)

	w.write(js)
}

func searchSimilarQuestions(String, question) []Question {
	// return type []Question or []string, which is better?

	// create a database client
	
	// search for similar questions. need some research.
}