package main

import (
	"log"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
	r := mux.NewRouter()
	log.Printf("[%s]\n", color.YellowString("PWND starting server"))

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))

	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")

	//r.HandleFunc("/assets/", mainHandler)

	r.HandleFunc("/answer", checkPasswordHandler).Methods("POST")
	r.HandleFunc("/answer", getPasswordAnswerHandler).Methods("GET")

	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}

/*func main() {
	// TODO make better flags
	cmdpass := flag.String("pass", "", "Search Have I been pwnd for a password.")
	cmdemail := flag.String("email", "", "Search Have I been pwnd for a email")
	flag.Parse()

	// TODO improve flag processing -> call -> return
	if *cmdpass != "" {
		pass := pwnd.Password{Word: *cmdpass}
		pass.Check()
		if !pass.Found() && pass.Error == nil {
			//Found nothing
			fmt.Printf("Found nothing\n")
		}
		// check Error
		fmt.Printf("%s - %d\n", color.CyanString(pass.Word), pass.Value())
	}

	if *cmdemail != "" {
		fmt.Printf("%s\n", pwnd.Checkemail(*cmdemail))
	}
}*/

// TODO API-Request
// TODO NICETOHAVE GUI / Webinterface
