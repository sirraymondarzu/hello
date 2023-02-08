// create a route  for greeting
// create a handler greeting that will output welcome to my page
package handlers

import (
	"fmt" // display to screen
	// display to terminal
	"net/http" // create a multiplexer and create routes/endpoints

	"github.com/sirraymondarzu/hello/helpers"
	// extract current time formats
)

// handler function
func Home(w http.ResponseWriter, r *http.Request) { // signature has a variable and a pointer
	/*	if r.Method != "POST" {
		w.Header().Set("ALLOW", "POST")
		w.WriteHeader(405)
		w.Write([]byte("Method not allowed"))
		return
	} */
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	//new code

	helpers.RenderTemplates(w, "./static/html/home.page.tmpl")

}

// handler
func About(w http.ResponseWriter, r *http.Request) {
	poll := "polls"
	w.Write([]byte(fmt.Sprintf("We do %s", poll)))
	//w.Write([]byte("Welcome to my about page"))
}

// handler
func HandlePoll(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to my handle poll page"))
}
