package datastore

import (
	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"net/http"
)

const (
	KIND_USER string = "USER"
)

// Stores the given model for the the kind in data store
func Store(req *http.Request, kind string, u User) error {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, kind, u.Email, 0, nil)
	_, err := datastore.Put(ctx, key, &u)
	return err
}

// Retrieves the model passed using the ID inside of it for the kind given.
func Retrieve(req *http.Request, kind, id string) (User, error) {
	ctx := appengine.NewContext(req)
	key := datastore.NewKey(ctx, kind, id, 0, nil)
	var u User
	err := datastore.Get(ctx, key, &u)
	return u, err
}
