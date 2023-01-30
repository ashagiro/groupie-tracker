package groupie

import (
	"html/template"
	"net/http"
	"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound)
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
		ErrorHandler(w, 405)
		return
	}
}

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	id := ""
	if len(r.URL.Path) > 7 {
		id = r.URL.Path[7:]
	}
	index, err := strconv.Atoi(id)
	if err != nil || !(index > 0 && index < len(All.Artists)+1) {
		ErrorHandler(w, 404)
		return
	}
	if r.URL.Path == "/group"+id {
		ErrorHandler(w, 404)
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
		temp, err := template.ParseFiles("./templates/group.html")
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
	w.WriteHeader(status)
	temp, err := template.ParseFiles("./templates/error.html")
	if err != nil {
		ErrorHandler(w, 500)
		return
	}
	temp.Execute(w, http.StatusText(status))
}
