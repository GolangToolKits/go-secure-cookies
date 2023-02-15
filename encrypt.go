package gosecurecookies

// Encrypt Encrypt
type Encrypt interface {
	Encrypt(s string) (string,error)
	Decrypt(es string) (string, error)
}
