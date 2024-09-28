package groupie

import (
	"bytes"
	"html/template"
	"net/http"
)

// HandleHome serves the home page by fetching a list of artists from an API and rendering it with a template, handling errors appropriately.
func HandleHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusBadRequest)
		return
	}

	if r.URL.Path != "/" {
		HandleError(w, http.StatusNotFound)
		return
	}

	var artists []Artists
	err := fetch("https://groupietrackers.herokuapp.com/api/", "artists", &artists)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, artists)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
	_, err = w.Write(buf.Bytes())
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
}
