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

// SpotifyView gère la vue Spotify
type SpotifyView struct {
	window        fyne.Window
	searchService *services.SearchService
	data          *models.APIData
}

// NewSpotifyView crée une nouvelle vue Spotify
func NewSpotifyView(window fyne.Window, searchService *services.SearchService, data *models.APIData) *SpotifyView {
	return &SpotifyView{
		window:        window,
		searchService: searchService,
		data:          data,
	}
}

// Render affiche la vue Spotify
func (v *SpotifyView) Render() *fyne.Container {
	header := widget.NewLabelWithStyle("🎵 Artistes & Albums", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Barre de recherche
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Rechercher un artiste, membre, date...")

	// Container pour les suggestions
	suggestionsBox := container.NewVBox()
	suggestionsScroll := container.NewVScroll(suggestionsBox)
	suggestionsScroll.Hide()

	// Liste des artistes
	artistList := container.NewVBox()

	updateArtistList := func(filter string) {
		artistList.Objects = nil

		if v.data == nil || len(v.data.Artists) == 0 {
			artistList.Add(widget.NewLabel("⏳ Chargement des données..."))
			artistList.Refresh()
			return
		}

		artists := v.searchService.SearchArtists(filter)

		if len(artists) == 0 {
			artistList.Add(widget.NewLabel("❌ Aucun résultat trouvé"))
		} else {
			for _, artist := range artists {
				artistCard := v.createArtistCard(artist)
				artistList.Add(artistCard)
			}
		}

		artistList.Refresh()
	}

	// Mise à jour des suggestions en temps réel
	searchEntry.OnChanged = func(query string) {
		if query == "" {
			suggestionsBox.Objects = nil
			suggestionsScroll.Hide()
			updateArtistList("")
			return
		}

		// Recherche universelle
		results := v.searchService.UniversalSearch(query)

		suggestionsBox.Objects = nil
		if len(results) > 0 {
			// Limiter à 10 suggestions
			maxResults := 10
			if len(results) > maxResults {
				results = results[:maxResults]
			}

			for _, result := range results {
				r := result // Capture pour la closure
				suggestionBtn := widget.NewButton(
					fmt.Sprintf("%s - %s", getTypeIcon(r.Type), r.Value),
					func() {
						if r.Artist != nil {
							v.showArtistDetails(*r.Artist)
						}
					},
				)
				suggestionBtn.Alignment = widget.ButtonAlignLeading
				suggestionsBox.Add(suggestionBtn)
			}
			suggestionsScroll.Show()
		} else {
			suggestionsScroll.Hide()
		}

		suggestionsBox.Refresh()
		updateArtistList(query)
	}

	// Initialisation
	go func() {
		time.Sleep(500 * time.Millisecond)
		updateArtistList("")
	}()

	scrollList := container.NewVScroll(artistList)
	scrollList.SetMinSize(fyne.NewSize(800, 600))

	// Layout avec suggestions
	searchContainer := container.NewBorder(
		nil, suggestionsScroll, nil, nil,
		searchEntry,
	)

	return container.NewBorder(
		container.NewVBox(header, searchContainer),
		nil, nil, nil,
		scrollList,
	)
}

// createArtistCard crée une carte pour un artiste
func (v *SpotifyView) createArtistCard(artist models.Artist) *fyne.Container {
	nameLabel := widget.NewLabelWithStyle(artist.Name, fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	nameLabel.TextStyle.Bold = true

	membersText := "👥 Membres: " + strings.Join(artist.Members, ", ")
	membersLabel := widget.NewLabel(membersText)
	membersLabel.Wrapping = fyne.TextWrapWord

	infoLabel := widget.NewLabel(fmt.Sprintf("📅 Création: %d | 💿 Premier album: %s | 🎸 Membres: %d",
		artist.CreationDate, artist.FirstAlbum, len(artist.Members)))

	// Bouton pour voir les détails
	detailsBtn := widget.NewButton("📋 Détails", func() {
		v.showArtistDetails(artist)
	})

	concertBtn := widget.NewButton("🎤 Voir concerts", func() {
		v.showConcerts(artist)
	})

	buttonsContainer := container.NewHBox(detailsBtn, concertBtn)

	card := container.NewVBox(
		nameLabel,
		membersLabel,
		infoLabel,
		buttonsContainer,
		widget.NewSeparator(),
	)

	return container.NewPadded(card)
}

// showArtistDetails affiche les détails d'un artiste
func (v *SpotifyView) showArtistDetails(artist models.Artist) {
	content := container.NewVBox(
		widget.NewLabelWithStyle(artist.Name, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel(fmt.Sprintf("🎸 Année de création: %d", artist.CreationDate)),
		widget.NewLabel(fmt.Sprintf("💿 Premier album: %s", artist.FirstAlbum)),
		widget.NewLabel(fmt.Sprintf("👥 Nombre de membres: %d", len(artist.Members))),
	)

	membersLabel := widget.NewLabelWithStyle("Membres:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	content.Add(membersLabel)

	for i, member := range artist.Members {
		content.Add(widget.NewLabel(fmt.Sprintf("  %d. %s", i+1, member)))
	}

	closeBtn := widget.NewButton("Fermer", func() {})

	dialog := widget.NewModalPopUp(
		container.NewBorder(
			content,
			container.NewCenter(closeBtn),
			nil, nil,
		),
		v.window.Canvas(),
	)

	closeBtn.OnTapped = func() {
		dialog.Hide()
	}

	dialog.Show()
}

// showConcerts affiche les concerts d'un artiste
func (v *SpotifyView) showConcerts(artist models.Artist) {
	concerts := v.searchService.GetConcertsByArtistID(artist.ID)

	concertContent := container.NewVBox()

	if len(concerts) == 0 {
		concertContent.Add(widget.NewLabel("❌ Aucun concert programmé"))
	} else {
		for _, concert := range concerts {
			locationLabel := widget.NewLabelWithStyle(
				fmt.Sprintf("📍 %s", services.FormatLocation(concert.Location)),
				fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
			concertContent.Add(locationLabel)

			for _, date := range concert.Dates {
				dateLabel := widget.NewLabel(fmt.Sprintf("  📅 %s", date))
				concertContent.Add(dateLabel)
			}
			concertContent.Add(widget.NewSeparator())
		}
	}

	closeBtn := widget.NewButton("Fermer", func() {})

	scroll := container.NewVScroll(concertContent)
	scroll.SetMinSize(fyne.NewSize(500, 400))

	dialogContent := container.NewBorder(
		container.NewVBox(
			widget.NewLabelWithStyle(fmt.Sprintf("🎤 Concerts de %s", artist.Name),
				fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
			widget.NewSeparator(),
		),
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

// getTypeIcon retourne l'icône pour un type de résultat
func getTypeIcon(resultType string) string {
	switch resultType {
	case "artist":
		return "🎸"
	case "member":
		return "👤"
	case "location":
		return "📍"
	case "album":
		return "💿"
	case "date":
		return "📅"
	default:
		return "🔍"
	}
}
