package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
