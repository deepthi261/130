package handler

import (
	"github.com/GoProject/GoProject/log"
	"github.com/GoProject/GoProject/memcache"
	"github.com/GoProject/GoProject/session"
	"net/http"
)

// Logout handler
func LogoutHandler(res http.ResponseWriter, req *http.Request) {
	// Clearing the session ID if its present and redirecting the user to font page.
	cookie, err := req.Cookie(session.SESSION_ID)
	if err == nil {
		// Clearing the cookie
		cookie.MaxAge = -1
		http.SetCookie(res, cookie)
		// Clears the sessionId given from sessions
		err := memcache.Delete(cookie.Value, req)
		log.LogErrorWithMsg("Cannot logout the user", err)

	} else {
		log.Println("The session is not set, skipping the logic.")
	}
	// Redirecting the user to front page.
	http.Redirect(res, req, URL_ROOT, http.StatusFound)
}
