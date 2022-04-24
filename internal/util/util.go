package util

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// ReadJSON reads json from request body into data. only support a single json value in the body
func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})

	if err != io.EOF {
		return errors.New("invalid json body")
	}
	return nil
}

// WriteJSON write arbitrary data as Json and set json content type
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
	return nil
}
