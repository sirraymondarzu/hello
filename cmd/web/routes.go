package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	//create multiplexer
	mux := http.NewServeMux()
	//hello world
	// create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // if the route does not exist use aternative

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)
	mux.HandleFunc("/poll", app.HandlePoll) // register the handler function

	return mux
}
