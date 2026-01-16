package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	APIBaseURL = "https://groupietrackers.herokuapp.com/api"
	Timeout    = 15 * time.Second
)

var (
	AllArtists   []Artist
	AllLocations []Location
	AllRelations []Relation
)

// FetchArtists récupère tous les artistes
func FetchArtists() error {
	fmt.Println("🔄 Chargement des artistes...")
	client := &http.Client{Timeout: Timeout}
	resp, err := client.Get(APIBaseURL + "/artists")
	if err != nil {
		fmt.Println("❌ Erreur connexion:", err)
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("erreur API: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &AllArtists)
	if err != nil {
		fmt.Println("❌ Erreur JSON:", err)
		return err
	}

	fmt.Printf("✅ %d artistes chargés\n", len(AllArtists))
	return nil
}

// FetchLocations récupère tous les lieux
func FetchLocations() error {
	fmt.Println("🔄 Chargement des locations...")
	client := &http.Client{Timeout: Timeout}
	resp, err := client.Get(APIBaseURL + "/locations")
	if err != nil {
		fmt.Println("❌ Erreur connexion:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &AllLocations)
	if err != nil {
		fmt.Println("❌ Erreur JSON:", err)
		return err
	}

	fmt.Printf("✅ %d locations chargées\n", len(AllLocations))
	return nil
}

// FetchRelations récupère les relations artistes/lieux/dates
func FetchRelations() error {
	fmt.Println("🔄 Chargement des relations...")
	client := &http.Client{Timeout: Timeout}
	resp, err := client.Get(APIBaseURL + "/relation")
	if err != nil {
		fmt.Println("❌ Erreur connexion:", err)
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &AllRelations)
	if err != nil {
		fmt.Println("❌ Erreur JSON:", err)
		return err
	}

	fmt.Printf("✅ %d relations chargées\n", len(AllRelations))
	return nil
}

// FetchAllData charge toutes les données de l'API
func FetchAllData() error {
	if err := FetchArtists(); err != nil {
		return fmt.Errorf("artistes: %w", err)
	}
	if err := FetchLocations(); err != nil {
		return fmt.Errorf("locations: %w", err)
	}
	if err := FetchRelations(); err != nil {
		return fmt.Errorf("relations: %w", err)
	}
	return nil
}

// GetArtistByID récupère un artiste par ID
func GetArtistByID(id int) *Artist {
	for i := range AllArtists {
		if AllArtists[i].ID == id {
			return &AllArtists[i]
		}
	}
	return nil
}

// GetRelationByArtistID récupère les relations d'un artiste
func GetRelationByArtistID(id int) *Relation {
	for i := range AllRelations {
		if AllRelations[i].ID == id {
			return &AllRelations[i]
		}
	}
	return nil
}

// SearchArtists cherche des artistes par nom ou membres
func SearchArtists(query string) []Artist {
	var results []Artist
	queryLower := string(query)

	for _, artist := range AllArtists {
		// Chercher dans le nom
		if contains(artist.Name, queryLower) {
			results = append(results, artist)
			continue
		}

		// Chercher dans les membres
		for _, member := range artist.Members {
			if contains(member, queryLower) {
				results = append(results, artist)
				break
			}
		}

		// Chercher dans l'album
		if contains(artist.FirstAlbum, queryLower) {
			results = append(results, artist)
		}
	}

	return results
}

// FilterByCreationDate filtre les artistes par date de création
func FilterByCreationDate(artists []Artist, minYear, maxYear int) []Artist {
	var results []Artist
	for _, artist := range artists {
		if artist.CreationDate >= minYear && artist.CreationDate <= maxYear {
			results = append(results, artist)
		}
	}
	return results
}

func contains(s, substr string) bool {
	return strings.Contains(strings.ToLower(s), strings.ToLower(substr))
}
