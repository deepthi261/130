package storage

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"html/template"
	"log"
	"net/http"
)

func init() {
	http.HandleFunc("/", indexHandler)
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/isUser", isUserHandler)
}

// Index handler simply shows the index.html file
func indexHandler(res http.ResponseWriter, req *http.Request) {
	//Parsing the template
	tpl := template.Must(template.ParseFiles("index.html"))
	err := tpl.Execute(res, nil)
	logError(err)
}

// Checks to see if the new-word set on the request exists in memcache or not.
func isUserHandler(res http.ResponseWriter, req *http.Request) {
	userName := req.FormValue("new-word")
	log.Println("new-word: " + userName)
	userExists := isExistingUser(userName, req)
	if !userExists {
		saveUser(userName, req)
	}
	json.NewEncoder(res).Encode(userExists)
}

// Saves the userName given on memcache
func saveUser(userName string, req *http.Request) {
	ctx := appengine.NewContext(req)
	user := memcache.Item{
		Key:   userName,
		Value: []byte(""), // The user name is important for us only.
	}
	err := memcache.Set(ctx, &user)
	logError(err)
}

// Returns true if the userName given exists in memcache
func isExistingUser(userName string, req *http.Request) bool {
	ctx := appengine.NewContext(req)
	item, err := memcache.Get(ctx, userName)
	if err != nil {
		logError(err)
		return false
	}
	log.Println("item: " + item.Key)
	if item.Key == "" {
		return false
	}
	return true
}

// Logs the error given into log
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}
Status API Training Shop Blog About
