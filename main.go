package main

import (
	"fmt"
	groups "groupie/functions"
	"log"
	"net/http"
)

func main() {
	// Parsing JSON
	groups.ParseJson()

	// Serving static files
	fs := http.FileServer(http.Dir("assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Mainpage Handler
	http.HandleFunc("/", groups.IndexHandler)

	// Grouppage Handler
	http.HandleFunc("/group/", groups.GroupHandler)

	// Initiating port
	fmt.Println("Link -->   " + "http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
