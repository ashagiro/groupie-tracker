package groupie

import (
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, 404)
		return
	}
	if r.URL.Path == "/" && r.Method != "GET" || r.URL.Path == "artist" && r.Method != "POST" {
		ErrorHandler(w, 405)
		return
	}
	switch r.Method {
	case "GET":
		temp, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			ErrorHandler(w, 500)
			return
		}
		temp.Execute(w, All)
	default:
		ErrorHandler(w, 400)
		return
	}
}

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	id := ""
	if len(r.URL.Path) > 7 {
		id = r.URL.Path[7:]
	}
	if r.URL.Path != "/artist/"+id {
		ErrorHandler(w, 404)
		return
	}
	index, err := strconv.Atoi(id)
	if err != nil || !(index >= 0 && index < 52) {
		ErrorHandler(w, 404)
		return
	}
	if r.URL.Path == "/" && r.Method != "GET" || r.URL.Path == "artist" && r.Method != "POST" {
		ErrorHandler(w, 405)
		return
	}
	switch r.Method {
	case http.MethodGet:
		index -= 1
		artist := &Artist{
			ID:           All.Artists[index].ID,
			Name:         All.Artists[index].Name,
			Members:      All.Artists[index].Members,
			CreationDate: All.Artists[index].CreationDate,
			FirstAlbum:   All.Artists[index].FirstAlbum,
			Image:        All.Artists[index].Image,
			Rel:          All.Artists[index].Rel,
		}
		temp, err := template.ParseFiles("./templates/artist.html")
		if err != nil {
			ErrorHandler(w, 500)
		}
		temp.Execute(w, artist)
	default:
		ErrorHandler(w, 400)
		return
	}
}

func ErrorHandler(w http.ResponseWriter, status int) {
	var title string
	switch status {
	case 400:
		title = "400 Bad Request"
	case 404:
		title = "404 Not Found"
	case 405:
		title = "405 Method Not Allowed"
	case 500:
		title = "500 Internal Server Error"
	}
	w.WriteHeader(status)
	temp, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		ErrorHandler(w, 500)
		return
	}
	temp.Execute(w, title)
}
