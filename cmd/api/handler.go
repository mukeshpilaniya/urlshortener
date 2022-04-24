package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Payload struct {
	Error   bool   `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func (app *application) generateShortenerUrl(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	fmt.Println(body)
}

func (app *application) fetchShortenerUrl(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		app.errorLogger.Println(err)
		return
	}
	fmt.Println(body)
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
