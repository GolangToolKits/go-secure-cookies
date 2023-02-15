package gosecurecookies

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
)

const (
	envalidValue = "there is an invalid value in the cookie"
	enkeyerror   = "secretKey must be at least 16 char long"
)

// GCMEncrypt GCMEncryption
type GCMEncrypt struct {
	secretKey []byte
}

// NewEncrypt NewEncrypt
func NewEncrypt(secretKey string) (Encrypt, error) {
	//securekey must be at least 16 char long
	// The key argument should be the AES key,
	// either 16, 24, or 32 bytes to select
	// AES-128, AES-192, or AES-256.
	var rtn Encrypt
	var rtnObj GCMEncrypt
	var rtnErr error
	if len(secretKey) > 32 {
		// trim to 32
		secretKey = secretKey[0:32]
	} else if len(secretKey) > 24 {
		// trim to 24
		secretKey = secretKey[0:24]
	} else if len(secretKey) < 16 {
		rtnErr = errors.New(enkeyerror)
	}
	if rtnErr == nil {
		skstr := hex.EncodeToString([]byte(secretKey))
		rtnObj.secretKey, _ = hex.DecodeString(skstr)
		rtn = &rtnObj
	}
	return rtn, rtnErr
}

// Encrypt Encrypt
func (e *GCMEncrypt) Encrypt(s string) (string, error) {
	var rtn string
	var rtnErr error
	block, err := aes.NewCipher(e.secretKey)
	if err == nil {
		aesGCM, err := cipher.NewGCM(block)
		if err == nil {
			nsiz := aesGCM.NonceSize()
			nonce := make([]byte, nsiz)
			_, err = io.ReadFull(rand.Reader, nonce)
			if err == nil {
				encryptedValue := aesGCM.Seal(nonce, nonce, []byte(s), nil)
				rtn = string(encryptedValue)
			}
		}
	} else {
		rtnErr = err
	}
	return rtn, rtnErr
}

// Decrypt Decrypt
func (e *GCMEncrypt) Decrypt(es string) (string, error) {
	var rtn string
	var rtnErr error
	block, err := aes.NewCipher(e.secretKey)
	if err == nil {
		aesGCM, err := cipher.NewGCM(block)
		if err == nil {
			nonceSize := aesGCM.NonceSize()
			esLen := len(es)
			if esLen >= nonceSize {
				nonce := es[:nonceSize]
				ciphertext := es[nonceSize:]
				plaintext, err := aesGCM.Open(nil, []byte(nonce), []byte(ciphertext), nil)
				if err == nil {
					rtn = string(plaintext)
				} else {
					rtnErr = err
				}
			}
		}
	} else {
		rtnErr = err
	}
	return rtn, rtnErr
}
