package storage

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	storageLog "google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"html/template"
	"log"
	"net/http"
)

const BUCKET_NAME = "deepthi-thotakura.appspot.com"

func init() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("./css"))))
	http.HandleFunc("/", handler)
}

func handler(res http.ResponseWriter, req *http.Request) {

	// Creating new context and client.
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	logStorageError(ctx, "Could not create a new client", err)
	defer client.Close()

	//Parsing the template
	tpl := template.Must(template.ParseFiles("index.html"))
	err = tpl.Execute(res, getPhotoNames(ctx, client))
	logError(err)
}

// Returns the name of the photos uploaded in google storage
func getPhotoNames(ctx context.Context, client *storage.Client) []string {

	query := &storage.Query{
		MaxResults: 4, // Limiting the end result to 4 images.
		Delimiter:  "/",
		Prefix:     "photos/",
	}
	objs, err := client.Bucket(BUCKET_NAME).List(ctx, query)
	logError(err)

	var names []string
	for _, result := range objs.Results {
		names = append(names, result.Name)
	}
	return names
}

// Logs the error given into log
func logError(err error) {
	if err != nil {
		log.Println(err)
	}
}

// Logs the error given into storage log
func logStorageError(ctx context.Context, errMessage string, err error) {
	if err != nil {
		storageLog.Errorf(ctx, errMessage, err)
	}
}
