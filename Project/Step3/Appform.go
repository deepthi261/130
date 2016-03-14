package Appform

import (
	"net/http"
	"github.com/nu7hatch/gouuid"
	"html/template"
	"log"
	"strings"
)


//Serves index.html
func serve_the_webpage(res http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("index.html")
	if err != nil {
		log.Fatalln(err)
	}

	//create a cookie "session-fino"
	cookie, err := req.Cookie("session-fino")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name: "session-fino",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}
	}

	if req.FormValue("user_name") != "" && !strings.Contains(cookie.Value, "user_name") {
		cookie.Value = cookie.Value +
		`name=` + req.FormValue("user_name") +
		`age=` + req.FormValue("user_age")
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
