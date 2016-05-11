package handler

import (
	"crypto/sha256"
	"fmt"
	"github.com/GoProject/GoProject/session"
	"io"
	"net/http"
)

type APlusTemplate struct {
	Header Header
	Data   interface{}
}

type Header struct {
	IsLoggedIn bool
}

// Creates a template for the given data by including the standard header to it.
func GetAPlusTemplateHeader(req *http.Request, data interface{}) APlusTemplate {
	return APlusTemplate{
		Header: Header{
			IsLoggedIn: session.GetUser(req).Email != "",
		},
		Data: data,
	}
}

// Enctypts the password using MD5
func Encrypt(pass string) string {
	h := sha256.New()
	io.WriteString(h, pass)
	return fmt.Sprintf("%x", h.Sum(nil))
}
