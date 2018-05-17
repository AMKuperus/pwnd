package pwnd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

// TODO func fetchapi haveIbeenpwnd
// TODO struct for returned passworddata
// TODO methods for passworddata
// TODO struct for returned emaildata
// TODO methods for emaildata

func Checkpassword(pass string) int {
	// TODO apirequest -> fetch data -> make object -> process data -> return data.
	ret := requestpassword(pass)
	return ret
}

func Checkemail(email string) string {
	// TODO apirequest -> fetch data -> make object -> process data -> return data.
	ret := fmt.Sprintf("Email to look for: %s", email)
	return ret
}

func requestpassword(input string) int {
	request := fmt.Sprintf("https://api.pwnedpasswords.com/pwnedpassword/%s", input)
	fmt.Printf("Request: %s\n", color.GreenString(request))
	resp, err := http.Get(request)
	if err != nil {
		fmt.Printf("%s %s\n", color.RedString("Error making HTTP-requets:"), err.Error())
	}
	defer resp.Body.Close()

	var ret int
	switch resp.StatusCode {
	case 404:
		//Not found — the password was not found in the Pwned Passwords repository
		ret = 0
	case 200:
		//Ok — the password was found in the Pwned Passwords repository and is
		//returned with a count of its prevalence
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("%s %s\n", color.RedString("Error reading data from reponse:"), err.Error())
		}
		ret, err = strconv.Atoi(string(data))
		if err != nil {
			fmt.Printf("%s %s\n", color.RedString("Error converting data string to int"), err.Error())
			return -1
		}
	case 301:
		//All API endpoints must be invoked over HTTPS. Any requests over HTTP will
		//result in a 301 response with a redirect to the same path on the secure
		//scheme. Only TLS versions 1.2 and 1.3 are supported; older versions of the
		//protocol will not allow a connection to be made.
		fmt.Printf("%s %s\n", color.RedString("Error cannot make request, https missing"), err.Error())
		ret = -1
	default:
		ret = -1
	}
	return ret
}
