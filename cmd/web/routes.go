package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"jesusmarin.dev/gostrap/ui"
)

// Routes, fileserver and middleware setup
func (app *application) routes() http.Handler {
	// Initiate a httproute router
	router := httprouter.New()

	// Define a fileserver. This use embed files in the binary app
	fileServer := http.FileServer(http.FS(ui.Files))
	// define the path for the fileserver
	router.Handler(http.MethodGet, "/static/*filepath", fileServer)

	// Routes definition
	router.HandlerFunc(http.MethodGet, "/", app.home)

	// Standard middleware managed by alice with some custom middlewares
	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	return standard.Then(router)
}
