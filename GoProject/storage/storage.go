package storage

import (
	"github.com/deepthi261/GoProject/log"
	"google.golang.org/appengine"
	"google.golang.org/cloud/storage"
	"io"
	"net/http"
	"strings"
)

const BUCKET_NAME = "gotraining-1271.appspot.com"

// Saves the given file for the given name and content under the username.
func Store(req *http.Request, userName, fileName string, file io.Reader) error {

	fileName = getFilePath(userName, fileName)

	// Creating new context and client.
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	defer client.Close()

	writer := client.Bucket(BUCKET_NAME).Object(fileName).NewWriter(ctx)
	writer.ACL = []storage.ACLRule{{
		storage.AllUsers,
		storage.RoleReader}}

	io.Copy(writer, file)
	writer.Close()
	return err
}

// Retrieve the file stored for the given name
func Retrieve(req *http.Request, userName, fileName string) (*storage.Reader, error) {
	// Creating new context and client.
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	log.LogErrorWithMsg("Cannot create a new client", err)
	defer client.Close()
	filePath := getFilePath(userName, fileName)
	return client.Bucket(BUCKET_NAME).Object(filePath).NewReader(ctx)
}

// Retrieves the list of files for the given username
func RetrieveFileList(req *http.Request, userName string) ([]string, error) {
	// Creating new context and client.
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	log.LogErrorWithMsg("Cannot create a new client", err)
	defer client.Close()

	query := &storage.Query{
		Delimiter: "/",
		Prefix:    userName + "/",
	}

	objs, err := client.Bucket(BUCKET_NAME).List(ctx, query)
	log.LogError(err)

	var names []string
	for _, result := range objs.Results {
		names = append(names, strings.Split(result.Name, "/")[1])
	}
	return names, err
}

// Creates file path based on the user and file names
func getFilePath(userName, fileName string) string {
	return strings.ToLower(userName + "/" + fileName)
}
