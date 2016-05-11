package handler

import (
	"encoding/json"
	"github.com/GoProject/GoProject/log"
	"github.com/GoProject/GoProject/session"
	"github.com/GoProject/GoProject/storage"
	"google.golang.org/appengine"
	"google.golang.org/appengine/urlfetch"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Giphy struct {
	Data []GiphyData
}

type GiphyData struct {
	Images map[string]GiphyImage
}

type GiphyImage struct {
	Url    string
	Width  string
	Height string
}

// Download handler
func DownloadHandler(res http.ResponseWriter, req *http.Request) {
	fileName := req.URL.Query().Get("fileName")
	rc, err := storage.Retrieve(req, session.GetUser(req).Email, fileName)
	log.LogErrorWithMsg("Cannot retrieve the file name given", err)
	if err == nil {
		res.Header().Add("content-type", rc.ContentType())
		io.Copy(res, rc)
		defer rc.Close()
	}
}

// Handler for downloading GIF from Giphy api
func DownloadGiphyHandler(res http.ResponseWriter, req *http.Request) {
	log.Println("DownloadGiphy Handler...")
	// Calling the API
	name := req.FormValue("name")
	c := appengine.NewContext(req)
	client := urlfetch.Client(c)
	resp, err := client.Get(getGiphyUrl(name))
	defer resp.Body.Close()
	log.LogErrorWithMsg("Cannot call URL:", err)
	body, err := ioutil.ReadAll(resp.Body)
	log.LogErrorWithMsg("Cannot read response:", err)
	var data Giphy
	err = json.Unmarshal(body, &data)
	log.LogErrorWithMsg("Cannot unmarshal", err)
	log.Println(data)

	// Downloading the image
	if len(data.Data) > 0 {
		resp, err = client.Get(data.Data[0].Images["fixed_width"].Url)
		log.LogErrorWithMsg("Cannot download image:", err)
		io.Copy(res, resp.Body)
	}
}

//Creates URL to call Giphy API
func getGiphyUrl(name string) string {
	return "http://api.giphy.com/v1/gifs/search?q=" + url.QueryEscape(name) + "&api_key=dc6zaTOxFJmzC"
}
