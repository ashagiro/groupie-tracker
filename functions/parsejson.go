package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// main struct
var All API

func ParseJson() {
	urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlRelations := "https://groupietrackers.herokuapp.com/api/relation"
	// urlLocations := "https://groupietrackers.herokuapp.com/api/locations"
	// urlDates := "https://groupietrackers.herokuapp.com/api/dates"
	ParseInfo(urlArtists, &All.Artists)
	ParseInfo(urlRelations, &All.Relation)
	// ParseInfo(urlLocations, &All.Locations)
	// ParseInfo(urlDates, &All.Dates)
	for i, v := range All.Relation.Index {
		All.Artists[i].Rel = v.DatesLocations
	}
	// fmt.Println(All.Artists[0])
}

func ParseInfo(url string, temp interface{}) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	text, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(text, &temp)
}
