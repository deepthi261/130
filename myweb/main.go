package webAssign1

import (
	"golang.org/x/net/context"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/cloud/storage"
	"html/template"
	"net/http"
	"os"
	"io"
	"crypto/sha1"
	"fmt"
)

var tpl *template.Template

const gcsBucket = "deepthi-thotakura.appspot.com"

type gcsPhotos struct {
	ctx    context.Context
	res    http.ResponseWriter
	bucket *storage.BucketHandle
	client *storage.Client
}

// set everything up.
func init() {
	tpl = template.Must(template.ParseGlob("*.html"))
	resourceHandler("css")
	resourceHandler("img")
	resourceHandler("photos")
	http.HandleFunc("/", mainPage)
}

//present main page.
func mainPage(res http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(res, req)
		return
	}
	gcs := configureCloud(res, req)
	photos := gcs.retrievePhotos()
	if len(photos) == 0 {
		photos = gcs.uploadPhotos()						// upload files from local directory.
	}
	tpl.Execute(res, photos)
}

// hangle css / img / photo resource directories to point to required resouce files.
func resourceHandler(rsrcDir string) {
	fs := http.FileServer(http.Dir(rsrcDir))
	fs = http.StripPrefix("/"+rsrcDir, fs)
	http.Handle("/"+rsrcDir+"/", fs)
}

// configure cloud context, bucket.
func configureCloud(res http.ResponseWriter, req *http.Request) (gcs *gcsPhotos) {
	ctx := appengine.NewContext(req)
	client, err := storage.NewClient(ctx)
	if err != nil {
		log.Errorf(ctx, "ERROR handler NewClient: ", err)
		return
	}
	defer client.Close()
	
	gcs = &gcsPhotos{
		ctx:    ctx,
		res:    res,
		client: client,
		bucket:   client.Bucket(gcsBucket),
	}
	return
}

// load photos from the cloud to memory.
func (gcs *gcsPhotos) retrievePhotos() (photos []string) {
        qry := &storage.Query {
            MaxResults: 6,
        }
	files, err := gcs.bucket.List(gcs.ctx, qry)
		if err != nil {
			log.Errorf(gcs.ctx, "listBucketDirMode: unable to list bucket %q: %v", gcsBucket, err)
			return
		}
	for _, name := range files.Results {
		photos = append(photos, name.Name)
	}
	return
}


// up load photos from local location to the cloud.
func (gcs *gcsPhotos) uploadPhotos() []string {
    ext := ".jpg"
    subDir := "photos/"
    testPhotos := []string{ "0","2","3","4","5","6","7","8","10","11","12","13","14","15","16","17","18","19","20","21","22","23","24"}
	 
    for _, name := range testPhotos {
        ffile := subDir + name + ext
        srcFile, fName, err := gcs.read(ffile)
        if err != nil {
            log.Errorf(gcs.ctx, "Open / read file error %v", err)
            return nil
        }
        fName = fName + ext
        gcs.write(fName, srcFile)
        log.Infof(gcs.ctx, "In File: %s, Out File: %s\n", ffile, fName)
    }
	 return gcs.retrievePhotos()
}

// read a file from the list.
func (gcs *gcsPhotos) read(filename string) (buf []byte, outFilename string, err error) {
    src, err := os.Open(filename)
    if err != nil {
        return
    }
    defer src.Close()
    size, err := src.Stat()
    if err != nil {
        return
    }
    buf = make([]byte, size.Size())
    _, err = src.Read(buf)
    if err != nil {
        return
    }
    outFilename = getSha(buf)
    return
}

func (gcs *gcsPhotos) write(filename string, data []byte) {
    writer := gcs.bucket.Object(filename).NewWriter(gcs.ctx)
    writer.ACL = []storage.ACLRule{
        {storage.AllUsers, storage.RoleReader},
        }
        defer writer.Close()
        writer.Write(data)
}

// encode filename.
func getSha(data []byte) string {
	h := sha1.New()
        io.WriteString(h, string(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}