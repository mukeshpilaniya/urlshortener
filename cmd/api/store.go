package main

import "time"

// URLShortener is a type of all URLs
type URLShortener struct {
	Hash         string    `json:"hash"`
	ShortenerURL string    `json:"shortener_url"`
	OriginalURL  string    `json:"original_url"`
	CreationDate time.Time `-`
}

var db URLStore

type URLStore struct {
	store map[string]URLShortener
}

func init() {
	db.store = make(map[string]URLShortener)
}

// dbStore func store URLShortener into the database
func dbStore(shortener URLShortener) bool {
	shortener.CreationDate = time.Now()
	db.store[shortener.Hash] = shortener
	return true
}
