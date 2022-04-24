package api

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	version = "1.0.0"
)

type application struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", 8080),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	app.infoLogger.Println(fmt.Sprintf("Starting urlShortener Server on port %d", 8080))

	return srv.ListenAndServe()
}

func main() {

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		infoLogger:  infoLog,
		errorLogger: errorLog,
	}

	err := app.serve()

	if err != nil {
		log.Fatal(fmt.Sprintln(err))
	}
}