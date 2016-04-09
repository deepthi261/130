package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"io"
	"log"
	"net/http"
	"strings"
)

type User struct {
	Name string
	Age  string
}

func main() {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		data := foo(req.FormValue("name"), req.FormValue("age"))
		code := getCode(data)
		cookie, err := req.Cookie("session-fino")
		if err != nil {
			id, _ := uuid.NewV4()
			cookie = &http.Cookie{
				Name:  "session-fino",
				Value: id.String() + "|" + data + "|" + code,
				// Secure: true,
				HttpOnly: true,
			}
			http.SetCookie(res, cookie)
		}
		res.Header().Set("Content-Type", "text/html; charset=utf-8")
		xs := strings.Split(cookie.Value, "|")
		usrPics := xs[1] + "Different"
		usrCode := xs[2]
		if usrCode == getCode(usrPics) {
			fmt.Fprintf(res, "Code Valid\n")
		} else {
			fmt.Fprintf(res, "Code Invalid\n")
		}
		err = tpl.Execute(res, nil)
		if err != nil {
			log.Fatalln(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}

func foo(name string, age string) string {
	user := User{
		Name: name,
		Age:  age,
	}

	bs, err := json.Marshal(user)
	if err != nil {
		fmt.Printf("error: ", err)
	}
	str := base64.URLEncoding.EncodeToString(bs)
	return str
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}