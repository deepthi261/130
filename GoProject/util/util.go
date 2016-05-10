package util

import (
	"github.com/deepthi261/GoProject/datastore"
	"github.com/deepthi261/GoProject/log"
	"github.com/deepthi261/GoProject/memcache"
	"github.com/deepthi261/GoProject/session"
	"net/http"
)

// Stores the user in memcache and data store
func SaveUser(req *http.Request, u datastore.User) {
	// Storing into memcache
	err := memcache.Store(u.Email, u, req)
	log.LogErrorWithMsg("Cannot store the user into memcache:", err)
	// Storing into datastore
	err = datastore.Store(req, datastore.KIND_USER, u)
	log.LogErrorWithMsg("Cannot store the user into datastore:", err)
}

// Get's user's information from memcache, if it does not exists, it will look into datastore.
func GetUserWithEmail(email string, req *http.Request) datastore.User {
	// Getting the data from memcache
	var u datastore.User
	err := memcache.Retrieve(email, req, &u)
	if err != nil {
		log.Println("Cannot get the user from memcache", err)
		// Getting the user from datastore
		u, err = datastore.Retrieve(req, datastore.KIND_USER, email)
		if err == nil {
			// Trying to store the data into memcache
			err := memcache.Store(email, u, req)
			log.LogErrorWithMsg("Cannot store the data retreived from datastore into memcache", err)
		} else {
			log.LogErrorWithMsg("Cannot retreive the data from datastore", err)
		}
	}
	return u
}

// Gets the user based on the logged in user in session
func GetUser(req *http.Request) datastore.User {
	// Getting user's email from session
	email := session.GetUser(req).Email
	return GetUserWithEmail(email, req)
}
