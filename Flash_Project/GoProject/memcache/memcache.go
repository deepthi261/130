package memcache

import (
	"encoding/json"
	"google.golang.org/appengine"
	"google.golang.org/appengine/memcache"
	"net/http"
)

// Stores any given interface using the key into memcache.
func Store(key string, value interface{}, req *http.Request) error {
	ctx := appengine.NewContext(req)
	bs, err := json.Marshal(value)
	if err != nil {
		return err
	}
	item := memcache.Item{
		Key:   key,
		Value: bs,
	}
	err = memcache.Set(ctx, &item)
	return err
}

// Retrieves an item from memcache from the given id and puts it into value
func Retrieve(id string, req *http.Request, value interface{}) error {
	ctx := appengine.NewContext(req)
	item, err := memcache.Get(ctx, id)
	if err != nil {
		return err
	}
	// unmarshal from JSON
	err = json.Unmarshal(item.Value, &value)
	if err != nil {
		return err
	}
	return nil
}

// Deletes a record from memcache for the given id
func Delete(id string, req *http.Request) error {
	ctx := appengine.NewContext(req)
	err := memcache.Delete(ctx, id)
	if err != memcache.ErrCacheMiss {
		return err
	}
	return nil
}
