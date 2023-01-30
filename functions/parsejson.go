package groupie

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// main struct
var All API

func ParseJson() {
	urlArtists := "https://groupietrackers.herokuapp.com/api/artists"
	urlRelations := "https://groupietrackers.herokuapp.com/api/relation"
	ParseInfo(urlArtists, &All.Artists)
	ParseInfo(urlRelations, &All.Relation)
	for i, v := range All.Relation.Index {
		All.Artists[i].Rel = v.DatesLocations
	}
}

func ParseInfo(url string, temp interface{}) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		return
	}
	text, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	json.Unmarshal(text, &temp)
}
