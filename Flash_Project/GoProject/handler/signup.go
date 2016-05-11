package handler

import (
	"encoding/json"
	"github.com/GoProject/GoProject/datastore"
	"github.com/GoProject/GoProject/log"
	"github.com/GoProject/GoProject/session"
	"github.com/GoProject/GoProject/util"
	"html/template"
	"net/http"
	"regexp"
)

// Signup handler
func SignupHandler(res http.ResponseWriter, req *http.Request) {

	if req.Method != "POST" && session.GetUser(req).Email != "" {
		// If user is already in session, we want to redirect him to front page.
		http.Redirect(res, req, URL_ROOT, http.StatusFound)
		return
	}

	var errs []string

	if req.Method == "POST" {
		email := req.FormValue("email")
		pass1 := req.FormValue("password1")
		pass2 := req.FormValue("password2")
		// Validatation comes here
		errs = validate(email, pass1, pass2, req)

		if errs == nil {
			// Create the user
			du := datastore.User{
				Email:     email,
				Password:  Encrypt(pass1),
				FirstName: req.FormValue("fname"),
			}

			// Store the user
			util.SaveUser(req, du)

			// Create session
			session.CreateSession(&res, req, session.User{Email: du.Email})

			// Redirect the user to front page
			http.Redirect(res, req, URL_ROOT, http.StatusFound)
		}
	}

	//Parsing the template
	tpl := template.Must(template.ParseFiles("template/signup.html"))
	err := tpl.Execute(res, errs)
	log.LogError(err)
}

// A handler for Ajax calls which writes a JSON on the response back, saying if the user is already taken or not.
func IsUserTaken(res http.ResponseWriter, req *http.Request) {
	userName := req.FormValue("username")
	log.Println("Checking to see if the user [" + userName + "] is already taken...")
	u := util.GetUserWithEmail(userName, req)
	json.NewEncoder(res).Encode(u.Email != "")
}

// validates to make sure the email format is correct. Also it checks to see if
// the passwords are matching, if not it will return an error message in HTML format
func validate(email, pass1, pass2 string, req *http.Request) []string {
	var err []string
	if valid, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email); !valid {
		err = append(err, "Wrong email format")
	}
	if pass1 == "" || pass2 == "" {
		err = append(err, "Password is mandatory")
	}
	if pass1 != pass2 {
		err = append(err, "Passwords do not match")
	}
	if u := util.GetUserWithEmail(email, req); u.Email != "" {
		err = append(err, "User already exists")
	}
	return err
}
