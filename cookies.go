package gosecurecookies

import (
	"net/http"
)

// Cookies Cookies
type Cookies interface {
	Write(w http.ResponseWriter, cookie http.Cookie) error
	Read(r *http.Request, name string) (string, error)
}

// NewCookies NewCookies
func NewCookies(secureKey string) (Cookies, error) {
	var rtn Cookies
	var rtnErr error
	//securekey must be at least 16 char long
	enc, err := NewEncrypt(secureKey)
	if err == nil {
		var rtnObj SecureCookies
		rtnObj.encrypt = enc
		rtn = &rtnObj
	} else {
		rtnErr = err
	}
	return rtn, rtnErr
}

// go mod init github.com/GolangToolKits/go-secure-cookies
