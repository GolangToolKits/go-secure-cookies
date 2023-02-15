package gosecurecookies

import (
	b64 "encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// SecureCookies SecureCookies
type SecureCookies struct {
	encrypt Encrypt
}

func (s *SecureCookies) Write(w http.ResponseWriter, cookie http.Cookie) error {
	var rtnErr error
	plaintext := fmt.Sprintf("%s:%s", cookie.Name, cookie.Value)
	entxt, err := s.encrypt.Encrypt(plaintext)
	if err == nil {
		cookie.Value = b64.StdEncoding.EncodeToString([]byte(entxt))
		http.SetCookie(w, &cookie)
	} else {
		rtnErr = err
	}
	return rtnErr
}

func (s *SecureCookies) Read(r *http.Request, name string) (string, error) {
	var rtn string
	var rtnErr error
	cookie, err := r.Cookie(name)
	if err == nil {
		b64txt := cookie.Value
		etx, _ := b64.StdEncoding.DecodeString(b64txt)
		plaintext, err := s.encrypt.Decrypt(string(etx))
		if err == nil {
			expectedName, value, ok := strings.Cut(string(plaintext), ":")
			if ok && expectedName == name {
				rtn = value
			}
		}
	}
	return rtn, rtnErr
}
