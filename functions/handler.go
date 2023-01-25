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
	// fmt.Println(len(All.Artists))
	switch r.Method {
	case "GET":
		temp, err := template.ParseFiles("./templates/home.html")
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
	ID := ""
	if len(r.URL.Path) > 7 {
		ID = r.URL.Path[7:]
	}
	if r.URL.Path != "/group/"+ID {
		ErrorHandler(w, 404)
		return
	}
	if r.URL.Path == "/" && r.Method != "GET" || r.URL.Path == "artist" && r.Method != "POST" {
		ErrorHandler(w, 405)
		return
	}
	type numb struct {
		// Index  int
		Total interface{}
		// Total struct {
		// 	ID           int      "json:\"id\""
		// 	Image        string   "json:\"image\""
		// 	Name         string   "json:\"name\""
		// 	Members      []string "json:\"members\""
		// 	CreationDate int      "json:\"creationDate\""
		// 	FirstAlbum   string   "json:\"firstAlbum\""
		// 	Rel          map[string]string
		// }
	}
	switch r.Method {
	case http.MethodGet:
		index, _ := strconv.Atoi(ID)
		if index < 0 && index > 52 {
			ErrorHandler(w, 400)
			return
		}

		// fmt.Print(All.Artists[index].ID)

		artist := &numb{
			Total: All.Artists[index],
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
