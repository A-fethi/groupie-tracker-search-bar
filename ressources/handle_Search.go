package groupie

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

type SearchResult struct {
	Value string
	Type  string
}

func HandleSearch(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		HandleError(w, http.StatusMethodNotAllowed)
		return
	}

	var artists []Artists

	err := fetch("https://groupietrackers.herokuapp.com/api/", "artists", &artists)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	var searchResults []SearchResult
	search := r.FormValue("search")

	for _, artist := range artists {
		if strings.Contains(strings.ToLower(artist.Name), strings.ToLower(search)) {
			searchResults = append(searchResults, SearchResult{Value: artist.Name, Type: "artist/band"})
		}
		if strings.Contains(strings.ToLower(strconv.Itoa(artist.CreationDate)), strings.ToLower(search)) {
			searchResults = append(searchResults, SearchResult{Value: strconv.Itoa(artist.CreationDate), Type: "creation date"})
		}
		if strings.Contains(strings.ToLower(artist.FirstAlbum), strings.ToLower(search)) {
			searchResults = append(searchResults, SearchResult{Value: artist.FirstAlbum, Type: "first album"})
		}
		for _, member := range artist.Members {
			if strings.Contains(strings.ToLower(member), strings.ToLower(search)) {
				searchResults = append(searchResults, SearchResult{Value: member, Type: "member"})
			}
		}
	}

	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}

	data := struct {
		Artists       []Artists
		SearchResults []SearchResult
	}{
		Artists:       artists,
		SearchResults: searchResults,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		HandleError(w, http.StatusInternalServerError)
		return
	}
}
