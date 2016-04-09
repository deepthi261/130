package main
import (
	"net/http"
	"html/template"
	"log"
	"github.com/nu7hatch/gouuid"
	"encoding/json"
	"encoding/base64"
	"crypto/hmac"
	"crypto/sha256"
	"io"
	"fmt"
	"strconv"
)

var template1 *template.Template
type UserData struct {
	Name string
	Age int
}
type User struct {
	Uuid, Name, Hmac string
	Password string
	Valid bool
	Data UserData
}
func init() {
	var err error
	template1, err = template.ParseFiles("index.html")
	if(err != nil){
		log.Println("Error: ", err)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/user-form", upload)
	http.ListenAndServe(":8080", nil)
}

func upload(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		cookie, err := req.Cookie("user-info")
		if err == http.ErrNoCookie {
			log.Println("Error no cookie after login: ", err)
			http.RedirectHandler("/", 303)
		}
		var user User
		user = decodeJsonData(cookie)
		uname := req.FormValue("uname")
		uage, err := strconv.Atoi(req.FormValue("uage"))
		if err != nil {
			log.Println("Error while converting uage: ", err)
		}
		userData := UserData{
			Name: uname,
			Age: uage,
		}
		user.Data = userData
		log.Println("Before template data looks like: ", user)
		template1.Execute(res, user)
	}
}

func home(res http.ResponseWriter, req *http.Request) {
	cookie, err1 := req.Cookie("user-info")
	user := User{}
	if(err1 == http.ErrNoCookie){
		uuid, _ := uuid.NewV4()
		user = User {
			Uuid: uuid.String(),
			Hmac: getCode(uuid.String()),
		}
		log.Println("UUID: ", user.Uuid)
		log.Println("HMAC: ", user.Hmac)
		encodedUser := encodeJsonData(user)
		log.Println("ENCODED: ", encodedUser)
		cookie = setNewCookie(encodedUser, cookie)
		http.SetCookie(res, cookie)
		user.Valid = true
	}
	if req.Method == "POST" {
		log.Println("POST REQUEST")
		name := req.FormValue("name")
		password := req.FormValue("password")
		log.Println("POST NAME: ", name)
		log.Println("POST PASSWORD: ", password)
		user = User{
			Name: name,
			Password: password,
		}
		var err error
		cookie, err = req.Cookie("user-info")
		if err != nil {
			log.Println("ERROR: ", err)
		}
		cookie.Value = updateCookie(user, req, cookie)
		user = decodeJsonData(cookie)
	}
	if req.URL.Path == "/logout" {
		log.Println("Logout Request!!")
		cookie = setNewCookie(user.Uuid, cookie)
		http.SetCookie(res, cookie)
		user.Valid = true
	}
	log.Println("Before template data looks like: ", user)
	template1.Execute(res, user)
}

func setNewCookie(userInfo string, cookie *http.Cookie) *http.Cookie {
	cookie = &http.Cookie{
		Name: "user-info",
		Value: userInfo,
		HttpOnly: true,
//		Secure: true,
	}
	return cookie
}

func updateCookie(user User, req *http.Request, cookie *http.Cookie) string {
	decodedUser := decodeJsonData(cookie)
	if decodedUser.Valid == false {
		log.Println("Authentication at stake !!")
		return encodeJsonData(user)
	}
	log.Println("PDATE UUID: ", decodedUser.Uuid)
	log.Println("PDATE NAME: ", user.Name)
	log.Println("PDATE Password: ", user.Password)
	user.Uuid = decodedUser.Uuid
	user.Hmac = getCode(user.Uuid + user.Name)
	log.Println("PDATE HMAC: ", user.Hmac)
	return encodeJsonData(user)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("key"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func encodeJsonData(user User) string {
	jsonUser, errJsonMarshalError := json.Marshal(user)
	if errJsonMarshalError != nil {
		log.Println("Error: ", errJsonMarshalError)
	}
	return base64.StdEncoding.EncodeToString(jsonUser)
}

func decodeJsonData(cookie *http.Cookie) User {
	log.Println("Cookie", cookie.Value)
	decode, err := base64.StdEncoding.DecodeString(cookie.Value)
	if err != nil {
		log.Println("Error: ", err)
		var user User
		user.Valid = false
		return user
	}
	var user User
	json.Unmarshal(decode,&user)
	log.Println("USER NAME: ", user.Name)
	log.Println("USER UUID: ", user.Uuid)
	log.Println("USER PASSWORD: ", user.Password)
	log.Println("USER HMAC: ", user.Hmac)
	log.Println("DECODE JSON: ", getCode(user.Uuid + user.Name))
	log.Println("USER HMAC: ", []byte(user.Hmac))
	log.Println("DECODE JSON: ", []byte(getCode(user.Uuid + user.Name)))
	if user.Hmac == getCode(user.Uuid + user.Name) {
		log.Println("Fuck its true")
		user.Valid = true
		return user
	}
	log.Println("auth fails")
	user.Valid = false
	return user
}