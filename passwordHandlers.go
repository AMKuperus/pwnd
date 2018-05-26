package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AMKuperus/pwnd/pwnd"
	"github.com/fatih/color"
)

// Answer holds answer
type Answer struct {
	Found bool `json:"found"`
	Value int  `json:"value"`
}

var answers []Answer

func getPasswordAnswerHandler(w http.ResponseWriter, r *http.Request) {
	answersBytes, err := json.Marshal(answers)
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(answersBytes)
}

func checkPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// Reset answers to keep 1 answer and delete the history.
	if len(answers) > 0 {
		answers = nil
	}
	pass := pwnd.Password{}
	err := r.ParseForm()
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pass.Word = r.Form.Get("password")
	pass.Check()
	if pass.Error != nil {
		log.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !pass.Found() {
		// Found nothing
		answer := Answer{}
		answer.Found = pass.Found()
		answers = append(answers, answer)
		log.Printf("%s\n", color.GreenString("Found nothing | End of this request."))
	} else {
		answer := Answer{}
		answer.Found = pass.Found()
		answer.Value = pass.Value()
		answers = append(answers, answer)
		log.Printf("%s\n", color.GreenString("Found the answer | End of this request."))
	}
	//log.Printf("%#v\n", answers)
	http.Redirect(w, r, "/assets/password.html", http.StatusFound)
}
