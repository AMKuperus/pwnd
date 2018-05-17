package main

import (
	"flag"
	"fmt"

	"github.com/AMKuperus/pwnd/pwnd"
	"github.com/fatih/color"
)

func main() {
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
}

// TODO API-Request
// TODO NICETOHAVE GUI / Webinterface
