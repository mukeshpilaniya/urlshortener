package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetMD5Hash(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])
}

func (app *application) generateShortenerUrl(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	var url URLShortener
	err = json.Unmarshal(body, &url)

	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	urlString := url.OriginalURL
	if urlString == "" {
		app.infoLogger.Println("url string is empty ")
		data := Payload{
			Error:   true,
			Message: fmt.Sprintf("url string is empty"),
		}
		w.WriteHeader(http.StatusBadRequest)
		out, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			app.errorLogger.Println(err)
			return
		}
		w.Write(out)
		return
	}
	hash := GetMD5Hash(urlString)
	hash = hash[:6]

	url.Hash = hash
	url.ShortenerURL = fmt.Sprintf("http://localhost:8080/api/v1/short_url?=%s", hash)
	//store in DB
	if !dbStore(url) {
		app.infoLogger.Println("Error while storing url in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	data := Payload{
		Error:   false,
		Message: fmt.Sprintf("visit short url http://localhost:8080/short_url?=%s", hash),
	}

	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	w.Write(out)
}

func (app *application) fetchShortenerUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Query().Get("short_url")

	if _, ok := db.store[shortUrl]; !ok {
		data := Payload{
			Error:   true,
			Message: fmt.Sprintf("short url link is either expired or does not exits"),
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		out, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			app.errorLogger.Println(err)
		}
		w.Write(out)
		return
	}
	fmt.Println(db.store[shortUrl].OriginalURL)
	w.Header().Set("location", db.store[shortUrl].OriginalURL)
	w.WriteHeader(http.StatusFound)
	return
}

func (app *application) defaultPage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	data := Payload{
		Error:   false,
		Message: "Please use /api/v1/generate_shortener_url to generate a shorter url",
	}

	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	w.Write(out)
}
