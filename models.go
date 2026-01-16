package main

// Artist représente un artiste de l'API
type Artist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// Location représente une location (lieu de concert)
type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}

// Relation représente les relations entre artistes et lieux/dates
type Relation struct {
	ID             int                 `json:"id"`
	ArtistID       int                 `json:"artistId"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// ConcertDate représente une date de concert
type ConcertDate struct {
	Date   string
	Artist string
}
