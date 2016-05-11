package session

import (
	"github.com/GoProject/GoProject/log"
	"github.com/GoProject/GoProject/memcache"
	"github.com/nu7hatch/gouuid"
	"net/http"
)

const SESSION_ID = "SESSION-ID"

// Handles session related operation. Checks to see if a user is in session.
func Handle(res http.ResponseWriter, req *http.Request) bool {
	isUserInSession, _ := isUserInSession(req)
	if isUserInSession {
		log.Println("User is in session.")
	} else {
		log.Println("User is not in session, setting the SESSION-ID ...")
	}
	return isUserInSession
}

// Creates a session by creating a new UUID and setting it on cookie and sessions.
func CreateSession(res *http.ResponseWriter, req *http.Request, user User) {
	newUuid, err := uuid.NewV4()
	sessionId := newUuid.String()
	log.LogError(err)
	memcache.Store(sessionId, user, req)
	createCookie(res, SESSION_ID, sessionId)
}

// Creates a cookie for the given name and value and sets it on the response
func createCookie(res *http.ResponseWriter, cookieName, cookieValue string) {
	// Setting the cookie
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    cookieValue,
		HttpOnly: true,
	}
	// Setting the cookie on the response back to the client
	http.SetCookie(*res, cookie)
}

// Checks to see if the user is logged in by looking at the sessionID stored on the request cookie
func isUserInSession(req *http.Request) (bool, *User) {
	sessionIdCookie, err := req.Cookie(SESSION_ID)
	var user User
	if err != nil {
		log.Println("Error reading SESSIONID:" + err.Error())
		return false, &user
	}
	// Retrieve the item from memcache
	memcache.Retrieve(sessionIdCookie.Value, req, &user)
	if user.Email != "" {
		return true, &user
	}
	return false, &user
}

// Gets the current user logged in
func GetUser(req *http.Request) *User {
	_, user := isUserInSession(req)
	return user
}
