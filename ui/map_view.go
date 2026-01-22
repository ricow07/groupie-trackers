package ui

import (
	"fmt"
	"groupie-tracker/models"
	"groupie-tracker/services"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// MapView gère la vue Carte des concerts
type MapView struct {
	window        fyne.Window
	searchService *services.SearchService
	data          *models.APIData
}

// NewMapView crée une nouvelle vue Carte
func NewMapView(window fyne.Window, searchService *services.SearchService, data *models.APIData) *MapView {
	return &MapView{
		window:        window,
		searchService: searchService,
		data:          data,
	}
}

// Render affiche la vue Carte
func (v *MapView) Render() *fyne.Container {
	header := widget.NewLabelWithStyle("🗺️ Carte des Concerts", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Barre de recherche
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste pour voir ses concerts...")

	// Filtres
	filterContainer := v.createFilters()

	// Liste des concerts
	concertList := container.NewVBox()

	updateConcertList := func(filter string, locationFilter string) {
		concertList.Objects = nil

		if v.data == nil || len(v.data.Artists) == 0 {
			concertList.Add(widget.NewLabel("⏳ Chargement des données..."))
			concertList.Refresh()
			return
		}

		artists := v.searchService.SearchArtists(filter)
		hasResults := false

		for _, artist := range artists {
			concerts := v.searchService.GetConcertsByArtistID(artist.ID)

			// Filtrage par lieu si nécessaire
			var filteredConcerts []models.Concert
			for _, concert := range concerts {
				if locationFilter == "" || services.FormatLocation(concert.Location) == locationFilter {
					filteredConcerts = append(filteredConcerts, concert)
				}
			}

			if len(filteredConcerts) == 0 {
				continue
			}

			hasResults = true

			// En-tête artiste
			artistHeader := widget.NewLabelWithStyle(
				fmt.Sprintf("🎸 %s", artist.Name),
				fyne.TextAlignLeading,
				fyne.TextStyle{Bold: true},
			)
			concertList.Add(artistHeader)

			// Grouper par lieu
			locationMap := make(map[string][]string)
			for _, concert := range filteredConcerts {
				formattedLocation := services.FormatLocation(concert.Location)
				locationMap[formattedLocation] = append(locationMap[formattedLocation], concert.Dates...)
			}

			// Afficher les concerts par lieu
			for location, dates := range locationMap {
				locationCard := v.createConcertCard(artist.Name, location, dates)
				concertList.Add(locationCard)
			}

			concertList.Add(widget.NewSeparator())
		}

		if !hasResults {
			concertList.Add(widget.NewLabel("❌ Aucun concert trouvé"))
		}

		concertList.Refresh()
	}

	// Variable pour le filtre de localisation
	currentLocationFilter := ""

	searchEntry.OnChanged = func(query string) {
		updateConcertList(query, currentLocationFilter)
	}

	// Initialisation
	go func() {
		time.Sleep(500 * time.Millisecond)
		updateConcertList("", "")
	}()

	scrollList := container.NewVScroll(concertList)
	scrollList.SetMinSize(fyne.NewSize(800, 600))

	return container.NewBorder(
		container.NewVBox(header, searchEntry, filterContainer),
		nil, nil, nil,
		scrollList,
	)
}

// createFilters crée les filtres pour la vue carte
func (v *MapView) createFilters() *fyne.Container {
	// Bouton pour afficher tous les concerts
	allConcertsBtn := widget.NewButton("🌍 Tous les concerts", func() {
		// Action déjà gérée par la recherche vide
	})

	// Bouton pour voir les statistiques
	statsBtn := widget.NewButton("📊 Statistiques", func() {
		v.showStats()
	})

	return container.NewHBox(allConcertsBtn, statsBtn)
}

// createConcertCard crée une carte pour un concert
func (v *MapView) createConcertCard(artistName, location string, dates []string) *fyne.Container {
	locationLabel := widget.NewLabelWithStyle(
		fmt.Sprintf("📍 %s", location),
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	datesContainer := container.NewVBox()
	for _, date := range dates {
		dateLabel := widget.NewLabel(fmt.Sprintf("  📅 %s", date))
		datesContainer.Add(dateLabel)
	}

	// Bouton pour voir sur la carte (simulation)
	viewMapBtn := widget.NewButton("🗺️ Voir sur la carte", func() {
		v.showLocationOnMap(location, artistName, dates)
	})

	card := container.NewVBox(
		locationLabel,
		datesContainer,
		viewMapBtn,
	)

	return container.NewPadded(card)
}

// showLocationOnMap simule l'affichage sur une carte
func (v *MapView) showLocationOnMap(location, artistName string, dates []string) {
	content := container.NewVBox(
		widget.NewLabelWithStyle(
			fmt.Sprintf("📍 %s", location),
			fyne.TextAlignCenter,
			fyne.TextStyle{Bold: true},
		),
		widget.NewSeparator(),
		widget.NewLabel(fmt.Sprintf("🎸 Artiste: %s", artistName)),
		widget.NewLabel(fmt.Sprintf("📅 Nombre de concerts: %d", len(dates))),
		widget.NewSeparator(),
	)

	datesLabel := widget.NewLabelWithStyle("Dates des concerts:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	content.Add(datesLabel)

	for _, date := range dates {
		content.Add(widget.NewLabel(fmt.Sprintf("  • %s", date)))
	}

	// Informations géographiques (simulation)
	content.Add(widget.NewSeparator())
	content.Add(widget.NewLabel("🌐 Coordonnées géographiques (exemple):"))
	content.Add(widget.NewLabel("   Latitude: 48.8566"))
	content.Add(widget.NewLabel("   Longitude: 2.3522"))
	content.Add(widget.NewLabel(""))
	content.Add(widget.NewLabel("💡 Une vraie carte interactive pourrait être intégrée ici"))

	closeBtn := widget.NewButton("Fermer", func() {})

	scroll := container.NewVScroll(content)
	scroll.SetMinSize(fyne.NewSize(500, 400))

	dialogContent := container.NewBorder(
		nil,
		container.NewCenter(closeBtn),
		nil, nil,
		scroll,
	)

	dialog := widget.NewModalPopUp(dialogContent, v.window.Canvas())

	closeBtn.OnTapped = func() {
		dialog.Hide()
	}

	dialog.Show()
}

// showStats affiche les statistiques des concerts
func (v *MapView) showStats() {
	if v.data == nil {
		return
	}

	// Calculer les statistiques
	totalArtists := len(v.data.Artists)
	totalConcerts := 0
	locationMap := make(map[string]int)
	countryMap := make(map[string]int)

	for _, relation := range v.data.Relations {
		for location, dates := range relation.DatesLocations {
			totalConcerts += len(dates)
			locationMap[location]++

			// Extraire le pays (dernier élément après split)
			parts := strings.Split(location, "-")
			if len(parts) > 0 {
				country := parts[len(parts)-1]
				countryMap[country]++
			}
		}
	}

	content := container.NewVBox(
		widget.NewLabelWithStyle("📊 Statistiques des Concerts", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel(fmt.Sprintf("🎸 Nombre d'artistes: %d", totalArtists)),
		widget.NewLabel(fmt.Sprintf("🎤 Nombre total de concerts: %d", totalConcerts)),
		widget.NewLabel(fmt.Sprintf("📍 Nombre de lieux différents: %d", len(locationMap))),
		widget.NewLabel(fmt.Sprintf("🌍 Nombre de pays: %d", len(countryMap))),
		widget.NewSeparator(),
	)

	// Top 5 des pays avec le plus de concerts
	content.Add(widget.NewLabelWithStyle("🏆 Top 5 des pays:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))

	// Trier les pays par nombre de concerts (simple)
	topCountries := getTopCountries(countryMap, 5)
	for i, country := range topCountries {
		content.Add(widget.NewLabel(fmt.Sprintf("  %d. %s - %d concerts", i+1, country.name, country.count)))
	}

	closeBtn := widget.NewButton("Fermer", func() {})

	scroll := container.NewVScroll(content)
	scroll.SetMinSize(fyne.NewSize(500, 400))

	dialogContent := container.NewBorder(
		nil,
		container.NewCenter(closeBtn),
		nil, nil,
		scroll,
	)

	dialog := widget.NewModalPopUp(dialogContent, v.window.Canvas())

	closeBtn.OnTapped = func() {
		dialog.Hide()
	}

	dialog.Show()
}

type countryCount struct {
	name  string
	count int
}

func getTopCountries(countryMap map[string]int, n int) []countryCount {
	var countries []countryCount
	for name, count := range countryMap {
		countries = append(countries, countryCount{name, count})
	}

	// Tri simple par insertion
	for i := 1; i < len(countries); i++ {
		key := countries[i]
		j := i - 1
		for j >= 0 && countries[j].count < key.count {
			countries[j+1] = countries[j]
			j--
		}
		countries[j+1] = key
	}

	if len(countries) > n {
		return countries[:n]
	}
	return countries
}
