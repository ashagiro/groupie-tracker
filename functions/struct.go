package groupie

type API struct {
	ID       int
	Artists  Artists
	Relation Relation
	// Dates     Dates
	// Locations Locations
}

type Artists []struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Rel          map[string]string
}

type Relation struct {
	Index []struct {
		ID             int64             `json:"id"`
		DatesLocations map[string]string `json:"datesLocations"`
	} `json:"index"`
}

// type Dates struct {
// 	Index []struct {
// 		ID    int      `json:"id"`
// 		Dates []string `json:"dates"`
// 	} `json:"index"`
// }

// type Locations struct {
// 	Index []struct {
// 		ID        int      `json:"id"`
// 		Locations []string `json:"locations"`
// 	} `json:"index"`
// }
