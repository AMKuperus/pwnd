package pwnd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/fatih/color"
)

// Password holds information gathered for password
type Password struct {
	Word  string `json:"password"`
	Found bool   `json:"found"`
	Value int    `json:"value,omitempty"`
	Error error  `json:"error,omitempty"`
}

// Check requests password
func (p *Password) Check() {
	p.request()
}

func (p *Password) request() {
	request := fmt.Sprintf("https://api.pwnedpasswords.com/pwnedpassword/%s", p.Word)
	log.Printf("Request: %s\n", color.GreenString(request))

	resp, err := http.Get(request)
	if err != nil {
		p.Error = err
		p.Found = false
		log.Printf("Error making request: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case 200:
		//Ok — the password was found in the Pwned Passwords repository and is
		//returned with a count of its prevalence
		data, ioerr := ioutil.ReadAll(resp.Body)
		if ioerr != nil {
			p.Error = ioerr
			p.Found = false
			log.Printf("Error reading body: %s", ioerr.Error())
			return
		}
		val, strerr := strconv.Atoi(string(data))
		if strerr != nil {
			p.Error = strerr
			p.Found = false
			log.Printf("Error converting data: %s", strerr.Error())
		}
		p.Found = true
		p.Value = val
	case 301:
		//All API endpoints must be invoked over HTTPS. Any requests over HTTP will
		//result in a 301 response with a redirect to the same path on the secure
		//scheme. Only TLS versions 1.2 and 1.3 are supported; older versions of the
		//protocol will not allow a connection to be made.
		p.Error = fmt.Errorf("Error: Request needs a secure connection %s", err.Error())
		p.Found = false
	case 404:
		//Not found — the password was not found in the Pwned Passwords repository
		p.Found = false
		p.Value = 0
	default:
		p.Error = fmt.Errorf("Oops it appears we missed something")
		p.Found = false
	}
}
