package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from golang!")
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TEST")
}
