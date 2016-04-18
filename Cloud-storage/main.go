package main

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	storageLog "google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"log"
	"net/http"
	"os"
)

func init() {
	http.HandleFunc("/", handler)
}

func handler(res http.ResponseWriter, req *http.Request) {
	// Saving the file
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	storageLog.Errorf(ctx, "Error: ", err)
	defer client.Close()

	writer := client.Bucket("learning-csci130.appspot.com").Object("helloworld.txt").NewWriter(ctx)
	writer.ACL = []storage.ACLRule{{
		storage.AllUsers,
		storage.RoleReader}}

	// Reading the file from disk
	reader, err := os.Open("helloworld.txt")
	log.Println(err)
	io.Copy(writer, reader)
	writer.Close()

	// Reading the file
	rd, err := client.Bucket("learning-csci130.appspot.com").Object("helloworld.txt").NewReader(ctx)
	log.Println(err)
	io.Copy(res, rd)
	defer rd.Close()
}