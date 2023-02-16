package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	//create multiplexer
	mux := httprouter.New()
	//hello world
	// create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handler(http.MethodGet, "/static/*flepath", http.StripPrefix("/static", fileServer)) // if the route does not exist use aternative

	mux.HandlerFunc(http.MethodGet, "/", app.Home)
	mux.HandlerFunc(http.MethodGet, "/about", app.About)
	mux.HandlerFunc(http.MethodPost, "/poll", app.HandlePoll) // register the handler function

	return mux
}
