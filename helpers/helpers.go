package helpers

import (
	"html/template"
	"log"
	"net/http"
)

func RenderTemplates(w http.ResponseWriter, tmpl string) {

	ts, err := template.ParseFiles(tmpl)
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
