package main

import (
	"fmt"
	"net/url"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// CreateArtistCard crée une card pour afficher un artiste
func CreateArtistCard(artist *Artist, onSelected func(*Artist)) *fyne.Container {
	// Titre avec nombre de membres
	title := widget.NewRichTextFromMarkdown("## " + artist.Name)

	// Infos principales
	info := widget.NewLabel(
		fmt.Sprintf("📅 Créé en %d\n💿 Premier album: %s\n👥 %d membres",
			artist.CreationDate, artist.FirstAlbum, len(artist.Members)),
	)
	info.Wrapping = fyne.TextWrapWord

	// Liste des membres
	membersText := strings.Join(artist.Members, ", ")
	if len(membersText) > 60 {
		membersText = membersText[:60] + "..."
	}
	members := widget.NewLabel("🎤 " + membersText)
	members.Wrapping = fyne.TextWrapWord

	// Bouton pour plus de détails
	detailBtn := widget.NewButton("📋 Détails & Concerts", func() {
		onSelected(artist)
	})

	// Card
	card := container.NewVBox(
		title,
		info,
		members,
		detailBtn,
	)

	return card
}

// CreateDetailWindow crée une fenêtre de détail pour un artiste
func CreateDetailWindow(artist *Artist, app fyne.App) fyne.CanvasObject {
	// En-tête
	title := canvas.NewText(artist.Name, nil)
	title.TextSize = 28
	title.TextStyle.Bold = true

	// Informations
	info := widget.NewLabel(
		fmt.Sprintf("🎵 Groupe créé en %d\n💿 Premier album: %s\n\n👥 Membres (%d):\n%s",
			artist.CreationDate,
			artist.FirstAlbum,
			len(artist.Members),
			strings.Join(artist.Members, "\n"),
		),
	)
	info.Wrapping = fyne.TextWrapWord

	// Bouton pour voir sur Google Maps
	mapBtn := widget.NewButton("🗺️ Voir les lieux de concert", func() {
		target, err := url.Parse("https://www.google.com/maps/search/" + strings.ReplaceAll(artist.Name, " ", "+"))
		if err == nil {
			app.OpenURL(target)
		}
	})

	// Bouton pour voir les relations (dates et lieux)
	relBtn := widget.NewButton("🎤 Voir les concerts", func() {
		relation := GetRelationByArtistID(artist.ID)
		if relation != nil {
			showConcertsWindow(artist, relation, app)
		}
	})

	// Layout
	content := container.NewVBox(
		title,
		widget.NewSeparator(),
		info,
		widget.NewSeparator(),
		mapBtn,
		relBtn,
	)

	return container.NewVScroll(content)
}

// showConcertsWindow affiche les concerts d'un artiste
func showConcertsWindow(artist *Artist, relation *Relation, app fyne.App) {
	w := app.NewWindow(artist.Name + " - Concerts")
	w.Resize(fyne.NewSize(600, 400))

	content := container.NewVBox()

	// Ajouter un titre
	title := widget.NewLabel("🎤 Dates de concert pour " + artist.Name)
	title.Wrapping = fyne.TextWrapWord
	content.Add(title)
	content.Add(widget.NewSeparator())

	// Afficher les dates et lieux
	if len(relation.DatesLocations) == 0 {
		content.Add(widget.NewLabel("Pas de concerts prévus"))
	} else {
		for location, dates := range relation.DatesLocations {
			// Titre du lieu
			locTitle := widget.NewLabel("📍 " + location)
			locTitle.Wrapping = fyne.TextWrapWord
			content.Add(locTitle)

			// Dates pour ce lieu
			for _, date := range dates {
				dateLabel := widget.NewLabel("  • " + date)
				dateLabel.Wrapping = fyne.TextWrapWord
				content.Add(dateLabel)
			}

			content.Add(widget.NewSeparator())
		}
	}

	w.SetContent(container.NewVScroll(content))
	w.Show()
}

// CreateSearchBar crée une barre de recherche avec filtres
func CreateSearchBar() (*widget.Entry, *widget.Select, fyne.CanvasObject) {
	// Champ de recherche
	search := widget.NewEntry()
	search.SetPlaceHolder("🔍 Rechercher un artiste, membre...")

	// Filtre par période
	periodSelect := widget.NewSelect([]string{
		"Toutes les périodes",
		"Avant 1980",
		"1980-1999",
		"2000 et après",
	}, func(s string) {})
	periodSelect.SetSelected("Toutes les périodes")

	// Layout
	container := container.NewVBox(
		search,
		periodSelect,
	)

	return search, periodSelect, container
}

// CreateLoadingWidget crée un widget de chargement
func CreateLoadingWidget() *widget.Label {
	label := widget.NewLabel("⏳ Chargement des données...\n\n🎵 Connexion à l'API Groupie Trackers\n👥 Récupération des artistes\n🗺️ Récupération des lieux\n🎤 Récupération des concerts")
	label.Wrapping = fyne.TextWrapWord
	return label
}

// CreateErrorWidget crée un widget d'erreur
func CreateErrorWidget(errMsg string) *widget.Label {
	label := widget.NewLabel("❌ Erreur: " + errMsg + "\n\nVérifiez votre connexion internet.")
	label.Wrapping = fyne.TextWrapWord
	return label
}
