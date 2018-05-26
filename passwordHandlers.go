package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/AMKuperus/pwnd/pwnd"
	"github.com/fatih/color"
)

type Answer struct {
	Password string `json:"password"`
	Found    bool   `json:"found"`
	Value    int    `json:"value"`
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
	pass := pwnd.Password{}

	err := r.ParseForm()
	if err != nil {
		log.Printf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	pass.Word = r.Form.Get("password")
	log.Printf("Checking: %s\n", pass.Word)
	pass.Check()
	if pass.Error != nil {
		log.Printf("Error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !pass.Found() {
		// Found nothing
		log.Printf("Found nothing for %s\n", pass.Word)
	}

	answer := Answer{}
	answer.Password = pass.Word
	answer.Found = pass.Found()
	answer.Value = pass.Value()
	answers = append(answers, answer)
	log.Printf("Found %s - %d\n", color.CyanString(pass.Word), pass.Value())
	log.Printf("%#v", answer)
	http.Redirect(w, r, "/assets/password.html", http.StatusFound)
}
