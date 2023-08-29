package main

import (
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	// Flag parameters
	addr := flag.String("addr", ":4000", "HTTP Network Address:Port to listen")
	flag.Parse()

	// InfoLog custom info Logger with output to OS Stdout
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// ErrorLog custom error logger with output to OS Stderr
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new template cache...
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize the application struct
	// for application config
	app := application{
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: templateCache,
	}

	// Initialize the Server struct (GO standard library)
	// for server config
	server := &http.Server{
		Handler:      app.routes(), // routes for the app
		Addr:         *addr,
		ErrorLog:     errorLog,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)

	// start the server
	err = server.ListenAndServe()
	errorLog.Fatal(err)
}
