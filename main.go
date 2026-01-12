package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

// --- MODELES ---
type Artist struct {
	ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

// --- LOGIQUE API ---
func fetchArtists() ([]Artist, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var artists []Artist
	err = json.NewDecoder(resp.Body).Decode(&artists)
	return artists, err
}

// --- INTERFACE ---
func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Groupie Tracker Pro")
	myWindow.Resize(fyne.NewSize(1000, 800))

	// Chargement des données au démarrage
	artists, err := fetchArtists()
	if err != nil {
		fmt.Println("Erreur API:", err)
		return
	}

	// Conteneur principal pour la grille d'artistes
	grid := container.New(layout.NewGridLayout(3))

	// Fonction pour remplir la grille
	renderArtists := func(filter string) {
		grid.Objects = nil // On nettoie
		filter = strings.ToLower(filter)

		for _, a := range artists {
			// Filtrage simple par nom ou membre
			match := strings.Contains(strings.ToLower(a.Name), filter)
			for _, m := range a.Members {
				if strings.Contains(strings.ToLower(m), filter) {
					match = true
				}
			}

			if match || filter == "" {
				// Image de l'artiste
				u, _ := storage.ParseURI(a.Image)
				img := canvas.NewImageFromURI(u)
				img.FillMode = canvas.ImageFillContain
				img.SetMinSize(fyne.NewSize(200, 200))

				// Carte avec infos
				info := fmt.Sprintf("Album: %s\nAnnée: %d", a.FirstAlbum, a.CreationDate)
				card := widget.NewCard(a.Name, info, img)
				
				// Bouton pour voir les détails (Action Client-Serveur)
				btn := widget.NewButton("Voir Concerts", func() {
					fmt.Println("Chargement des concerts pour", a.Name)
					// Ici vous pourriez ouvrir une nouvelle fenêtre pour la Map
				})
				
				item := container.NewVBox(card, btn)
				grid.Add(item)
			}
		}
		grid.Refresh()
	}

	// Barre de recherche
	search := widget.NewEntry()
	search.SetPlaceHolder("Rechercher un artiste ou un membre...")
	search.OnChanged = func(s string) {
		renderArtists(s)
	}

	// Layout final : Recherche en haut, Grille scrollable au centre
	scroll := container.NewVScroll(grid)
	content := container.NewBorder(
		container.NewVBox(widget.NewLabelWithStyle("GROUPIE TRACKER", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}), search),
		nil, nil, nil,
		scroll,
	)

	renderArtists("") // Premier affichage
	myWindow.SetContent(content)
	myWindow.ShowAndRun()
}