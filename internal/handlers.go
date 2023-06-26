package funcs

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	files := []string{
		"./ui/html/index.page.html",
		"./ui/html/base.layout.html",
	}
	if r.Method != "GET" {
		ErrorHandler(w, 405)
		return
	}
	temp, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, 500)
		return
	}
	err = temp.Execute(w, All)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, 500)
		return
	}
}

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || !(id > 0 && id < len(All.Artists)+1) {
		ErrorHandler(w, 404)
		return
	}
	files := []string{
		"./ui/html/group.page.html",
		"./ui/html/base.layout.html",
	}
	switch r.Method {
	case http.MethodGet:
		id -= 1
		artist := &Artist{
			ID:           All.Artists[id].ID,
			Name:         All.Artists[id].Name,
			Members:      All.Artists[id].Members,
			CreationDate: All.Artists[id].CreationDate,
			FirstAlbum:   All.Artists[id].FirstAlbum,
			Image:        All.Artists[id].Image,
			Rel:          All.Artists[id].Rel,
		}
		temp, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, 500)
			return
		}
		err = temp.Execute(w, artist)
		if err != nil {
			log.Println(err.Error())
			ErrorHandler(w, 500)
			return
		}
	default:
		ErrorHandler(w, 400)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	files := []string{
		"./ui/html/error.page.html",
		"./ui/html/base.layout.html",
	}
	temp, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, 500)
		return
	}
	type errorText struct {
		Status int
		Text   string
	}
	msg := &errorText{
		Status: status,
		Text:   http.StatusText(status),
	}
	err = temp.Execute(w, msg)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, 500)
		return
	}
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/search" {
		ErrorHandler(w, http.StatusNotFound)
		return
	}
	files := []string{
		"./ui/html/search.page.html",
		"./ui/html/base.layout.html",
	}
	if r.Method != "POST" {
		ErrorHandler(w, 405)
		return
	}
	temp, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, 500)
		return
	}
	Search(r.FormValue("searching"))
	err = temp.Execute(w, FoundArtists)

	if err != nil {
		log.Println(err.Error())
		ErrorHandler(w, 500)
		return
	}
}
