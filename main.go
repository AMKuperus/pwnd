package main

import (
	"flag"
	"fmt"

	"github.com/AMKuperus/pwnd/pwnd"
)

func main() {
	// TODO make better flags
	cmdpass := flag.String("pass", "", "Search Have I been pwnd for a password.")
	cmdemail := flag.String("email", "", "Search Have I been pwnd for a email")
	flag.Parse()

	// TODO improve flag processing -> call -> return
	if *cmdpass != "" {
		fmt.Printf("%s\n", pwnd.Checkpassword(*cmdpass))
	}
	if *cmdemail != "" {
		fmt.Printf("%s\n", pwnd.Checkemail(*cmdemail))
	}
}

// TODO API-Request
// TODO NICETOHAVE GUI / Webinterface
