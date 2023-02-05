// create a route  for greeting
// create a handler greeting that will output welcome to my page
package handlers

import (
	"fmt" // display to screen
	"html/template"
	"log"      // display to terminal
	"net/http" // create a multiplexer and create routes/endpoints
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

	//w.Write([]byte("Welcome to polly."))
	ts, err := template.ParseFiles("./templates/home.page.tmpl")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "internal server error", 500)
		return
	}
	err = ts.Execute(w, nil)

	if err != nil {
		log.Print(err.Error())
		http.Error(w, "internal sefer error", 500)
		//no return because we are at the end of the code
	}

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
