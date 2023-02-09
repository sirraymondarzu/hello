// create a route  for greeting
// create a handler greeting that will output welcome to my page
package main

import (
	"flag"
	"log"      // display to terminal
	"net/http" // create a multiplexer and create routes/endpoints
)

// create a new type 
type application struct{}

func main() {
	// create a flag for specifing the port number when starting the server

	addr := flag.String("port", ":4000", "HTTP Network address") //if no port assigned // stored as a pointer
	flag.Parse()                                                 // this should be called only once
	//create instance of application type
	app := &application{} // creating app type

	//create multiplexer
	mux := http.NewServeMux()
	//hello world
	// create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // if the route does not exist use aternative

	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/about", app.About)
	mux.HandleFunc("/poll", app.HandlePoll) // register the handler function

	log.Printf("Server is active on port %s", *addr)

	err := http.ListenAndServe(*addr, mux) // using the pointer of addr -
	log.Fatal(err)

}
