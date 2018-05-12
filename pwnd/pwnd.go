package pwnd

import "fmt"

// TODO func fetchapi haveIbeenpwnd
// TODO struct for returned passworddata
// TODO methods for passworddata
// TODO struct for returned emaildata
// TODO methods for emaildata

func Checkpassword(pass string) string {
	// TODO apirequest -> fetch data -> make object -> process data -> return data.
	ret := fmt.Sprintf("Password to look for: %s", pass)
	return ret
}

func Checkemail(email string) string {
	// TODO apirequest -> fetch data -> make object -> process data -> return data.
	ret := fmt.Sprintf("Email to look for: %s", email)
	return ret
}
