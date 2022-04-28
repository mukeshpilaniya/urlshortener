package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/mukeshpilaniya/urlshortener/internal/util"
	"net/http"
)

// Payload is a struct for sending custom message
type Payload struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

// GetMD5Hash generate MD5 hash for the given string
func GetMD5Hash(url string) string {
	hash := md5.Sum([]byte(url))
	return hex.EncodeToString(hash[:])
}

// generateShortenerUrl will generate a short url for the given string
func (app *application) generateShortenerUrl(w http.ResponseWriter, r *http.Request) {
	var url URLShortener
	var data Payload

	err := util.ReadJSON(w, r, &url)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	urlString := url.OriginalURL

	if urlString == "" {
		app.infoLogger.Println("url string is empty ")
		data = Payload{
			Error:   true,
			Message: fmt.Sprintf("url string is empty"),
		}
		err = util.WriteJSON(w, http.StatusBadRequest, data)
		if err != nil {
			app.errorLogger.Println(err)
		}
		return
	}

	hash := GetMD5Hash(urlString)
	hash = hash[:6]
	url.Hash = hash
	url.ShortenerURL = fmt.Sprintf("http://localhost:8080/api/v1/short_url?=%s", hash)

	//store hash url in DB
	if !dbStore(url) {
		app.infoLogger.Println("Error while storing url in database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data = Payload{
		Error:   false,
		Message: fmt.Sprintf("visit short url http://localhost:8080/api/v1/url?short_url=%s", hash),
	}

	err = util.WriteJSON(w, http.StatusOK, &data)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	app.infoLogger.Println("short url is generated for ", url.OriginalURL, " url")
}

// fetchShortenerUrl fetch long url from DB based on hash value
func (app *application) fetchShortenerUrl(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Query().Get("short_url")
	var data Payload
	if shortUrl == "" {
		data = Payload{
			Error:   true,
			Message: fmt.Sprintf("URL is not a valid one"),
		}
		w.WriteHeader(http.StatusBadRequest)
		err := util.WriteJSON(w, http.StatusBadRequest, &data)
		if err != nil {
			app.errorLogger.Println(err)
		}
		return
	}

	if _, ok := db.store[shortUrl]; !ok {
		data = Payload{
			Error:   true,
			Message: fmt.Sprintf("short url link is either expired or does not exits"),
		}
		err := util.WriteJSON(w, http.StatusBadRequest, &data)
		if err != nil {
			app.errorLogger.Println(err)
		}
		return
	}
	app.infoLogger.Println("Redirecting", r.URL, "request to ", db.store[shortUrl].OriginalURL, " url")
	w.Header().Set("location", db.store[shortUrl].OriginalURL)
	w.WriteHeader(http.StatusFound)
	return
}
