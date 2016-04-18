package Storage-Query

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"io"
	"net/http"
	"strings"
)

const gcsBucket = "gorosave-project-1.appspot.com"

type Entity struct {
	Value string
}

type demo struct {
	context context.Context
	w       http.ResponseWriter
	bucket  *storage.BucketHandle
	client  *storage.Client
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	context := appengine.NewContext(r)

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	client, err := storage.NewClient(context)
	if err != nil {
		log.Errorf(context, "error handler newClient: ", err)
		return
	}

	defer client.Close()

	d := &demo{
		context: context,
		w:       w,
		client:  client,
		bucket:  client.Bucket(gcsBucket),
	}

	d.createFiles()
	d.listFiles()
	io.WriteString(d.w, "\nResult From ListDir - With Delimiter\n")
	d.listDir()

}

func (d *demo) listDir() {
	query := &storage.Query{
		Delimiter: "/",
	}

	objs, err := d.bucket.List(d.context, query)
	if err != nil {
		log.Errorf(d.context, "listBucketDirMode: unable to list bucket %q: %v", gcsBucket, err)
		return
	}

	for _, obj := range objs.Results {
		fmt.Fprintf(d.w, "%v/n", obj.Name)
	}
	io.WriteString(d.w, "\nPrefixes Found\n")
	fmt.Fprintf(d.w, "%v", objs.Prefixes)
}

func (d *demo) listFiles() {
	io.WriteString(d.w, "\nRetrieving File Names\n")

	client, err := storage.NewClient(d.context)
	if err != nil {
		log.Errorf(d.context, "%v", err)
		return
	}
	defer client.Close()

	objs, err := client.Bucket(gcsBucket).List(d.context, nil)
	if err != nil {
		log.Errorf(d.context, "%v", err)
		return
	}

	for _, obj := range objs.Results {
		io.WriteString(d.w, obj.Name+"\n")
	}
}

func (d *demo) createFiles() {
	io.WriteString(d.w, "\nCreating more files for listbucket...\n")
	for _, n := range []string{"foo1", "foo2", "bar", "bar/1", "bar/2", "boo/", "boo/yah", "compadre/amigo/diaz", "compadre/luego/hasta", "bar/nonce/1", "bar/nonce/2", "bar/nonce/compadre/1", "bar/nonce/compadre/2"} {
		d.createFile(n)
	}
}

func (d *demo) createFile(fileName string) {
	fmt.Fprintf(d.w, "creating file /%v/%v\n", gcsBucket, fileName)

	wc := d.bucket.Object(fileName).NewWriter(d.context)
	wc.ContentType = "text/plain"

	if _, err := wc.Write([]byte("abcde\n")); err != nil {
		log.Errorf(d.context, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}

	if _, err := wc.Write([]byte(strings.Repeat("f", 1024*4) + "\n")); err != nil {
		log.Errorf(d.context, "createFile: unable to write data to bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}

	if err := wc.Close(); err != nil {
		log.Errorf(d.context, "createFile: unable to close bucket %q, file %q: %v", gcsBucket, fileName, err)
		return
	}
}