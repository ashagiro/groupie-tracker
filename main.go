package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
)

type Artists []struct {
	// ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

type Locations struct {
	Index []struct {
		Location []string `json:"locations"`
	} `json:"index"`
}

func main() {
	res0, err0 := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err0 != nil {
		fmt.Print(err0.Error())
		os.Exit(1)
		// log.Fatal(err)
	}
	// defer res.Body.Close()
	res1, err0 := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err0 != nil {
		fmt.Print(err0.Error())
		os.Exit(1)
	}
	text0, err := ioutil.ReadAll(res0.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artists Artists
	json.Unmarshal(text0, &artists)

	text1, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		log.Fatal(err)
	}
	var location Locations
	json.Unmarshal(text1, &location)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/home.html")
		if err != nil {
			return
		}
		t.Execute(w, artists)
		// for _, v := range artists {
		// 	t.Execute(w, v)
		// }
		// for _, k := range location.Index {
		// 	t.Execute(w, k)
		// }
	})
	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fmt.Println("Link -->   " + "http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
