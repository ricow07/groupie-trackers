package models

// Artist représente un artiste ou groupe
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// Location représente les lieux de concerts
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Date représente les dates de concerts
type Date struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}

// Relation lie les artistes, dates et lieux
type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// APIData contient toutes les données de l'API
type APIData struct {
	Artists   []Artist
	Locations []Location
	Dates     []Date
	Relations []Relation
}

// Concert représente un concert avec toutes ses informations
type Concert struct {
	ArtistID   int
	ArtistName string
	Location   string
	Dates      []string
}

// SearchResult représente un résultat de recherche
type SearchResult struct {
	Type   string // "artist", "member", "location", "date"
	Value  string
	Artist *Artist
}
