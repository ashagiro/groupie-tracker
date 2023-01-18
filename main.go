package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
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

	res1, err1 := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err1 != nil {
		fmt.Print(err1.Error())
		os.Exit(1)
	}

	text0, err := ioutil.ReadAll(res0.Body)
	if err != nil {
		log.Fatal(err)
	}

	text1, err := ioutil.ReadAll(res1.Body)
	if err != nil {
		log.Fatal(err)
	}
	var artists Artists
	json.Unmarshal(text0, &artists)

	var location Locations
	json.Unmarshal(text1, &location)

	// for i, p := range artists {
	// 	fmt.Println("Name", (i + 1), ":", p.Name)
	// 	fmt.Println("---------------------------")
	// 	fmt.Println("Image", p.Image)
	// 	fmt.Println("Members", p.Members)
	// 	fmt.Println()
	// }
	// fmt.Println(location.Index)
	for _, p := range location.Index {
		fmt.Println("Locations", ":", p.Location)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := template.ParseFiles("templates/home.html")
		if err != nil {
			return
		}
		// t.Execute(w, "hello")
		fmt.Fprintf(w, "Welcome to my website!")
	})

	fs := http.FileServer(http.Dir("static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
	fmt.Println("Link -->   " + "http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
