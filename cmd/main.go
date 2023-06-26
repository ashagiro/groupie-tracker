package main

import (
	"log"
	"net/http"

	funcs "01.alem.school/git/ashagiro/groupie-tracker-search-bar/internal"
)

func main() {
	// Parsing JSON
	funcs.ParseJson()

	mux := http.NewServeMux()

	// Serving static files
	mux.Handle("/ui/", http.StripPrefix("/ui", http.FileServer(http.Dir("./ui"))))

	// Handlers
	mux.HandleFunc("/", funcs.IndexHandler)
	mux.HandleFunc("/group/", funcs.GroupHandler)
	mux.HandleFunc("/search", funcs.SearchHandler)

	// Initiating port
	log.Println("Link -->   " + "http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
