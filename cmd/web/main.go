// create a route  for greeting
// create a handler greeting that will output welcome to my page
package main

import (
	"log"      // display to terminal
	"net/http" // create a multiplexer and create routes/endpoints

	"github.com/sirraymondarzu/hello/handlers" //separate local modules by a space
)

func main() {
	//create multiplexer
	mux := http.NewServeMux()
	//hello world
	// create file server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer)) // if the route does not exist use aternative

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/about", handlers.About)
	mux.HandleFunc("/poll", handlers.HandlePoll) // register the handler function

	log.Println("Server is active on port 4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)

}
