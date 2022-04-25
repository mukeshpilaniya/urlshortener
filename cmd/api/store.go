package main

import "time"

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

func dbStore(shortener URLShortener) bool {
	shortener.CreationDate = time.Now()
	db.store[shortener.Hash] = shortener
	return true
}
