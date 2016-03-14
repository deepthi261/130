package main

import (
	"net/http"
	"html/template"
	"log"
	"encoding/json"
	"encoding/base64"
	"strings"
)


type User struct {
	Name string
	Age string
}


func set_user(req *http.Request, user *User) string {
	user.Name = req.FormValue("user_name")
	user.Age = req.FormValue("user_age")

	bs, _ := json.Marshal(user) //marshal to type user to json

	b64 := base64.URLEncoding.EncodeToString(bs) //encode type user to base64


	return b64
}


//Serves index.html
func serve_the_webpage(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}


	cookie, err := req.Cookie("session-fino")
	//create cookie if not exists.
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",

			HttpOnly: true,
		}
	}


	user := new(User) //create new User instance.
	b64 := set_user(req, user)

	if req.FormValue("user_name") != "" {

		cookie_id := strings.Split(cookie.Value, "|")
		cookie.Value = cookie_id[0] + "|" + b64
	}

	http.SetCookie(res, cookie)

	tpl.Execute(res, nil)
}


func main() {
	http.HandleFunc("/", serve_the_webpage)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)
}
