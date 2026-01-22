package ui

import (
	"fmt"
	"groupie-tracker/models"
	"groupie-tracker/services"
	"math/rand"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// ShazamView gère la vue Shazam
type ShazamView struct {
	window        fyne.Window
	searchService *services.SearchService
	data          *models.APIData
	history       []ShazamResult
}

// ShazamResult représente un résultat de reconnaissance
type ShazamResult struct {
	Artist    models.Artist
	Timestamp time.Time
}

// NewShazamView crée une nouvelle vue Shazam
func NewShazamView(window fyne.Window, searchService *services.SearchService, data *models.APIData) *ShazamView {
	return &ShazamView{
		window:        window,
		searchService: searchService,
		data:          data,
		history:       []ShazamResult{},
	}
}

// Render affiche la vue Shazam
func (v *ShazamView) Render() *fyne.Container {
	header := widget.NewLabelWithStyle("🎤 Shazam - Reconnaissance Musicale", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})

	// Zone de résultat
	resultLabel := widget.NewLabel("Appuyez sur le bouton pour identifier une chanson")
	resultLabel.Wrapping = fyne.TextWrapWord
	resultLabel.Alignment = fyne.TextAlignCenter

	// Image de visualisation (animation simulée)
	visualizationLabel := widget.NewLabel("🎵 ♪ ♫ ♬ 🎶")
	visualizationLabel.Alignment = fyne.TextAlignCenter
	visualizationLabel.TextStyle = fyne.TextStyle{Bold: true}
	visualizationLabel.Hide()

	// Bouton d'écoute
	listenBtn := widget.NewButton("🎧 Écouter et Identifier", func() {
		v.performRecognition(resultLabel, visualizationLabel)
	})
	listenBtn.Importance = widget.HighImportance

	// Historique
	historyContainer := v.createHistoryView()

	// Statistiques
	statsBtn := widget.NewButton("📊 Mes statistiques", func() {
		v.showUserStats()
	})

	// Layout principal
	mainContent := container.NewVBox(
		header,
		layout.NewSpacer(),
		container.NewCenter(visualizationLabel),
		container.NewCenter(listenBtn),
		layout.NewSpacer(),
		resultLabel,
		layout.NewSpacer(),
		widget.NewSeparator(),
		container.NewHBox(
			widget.NewLabelWithStyle("📜 Historique", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}),
			layout.NewSpacer(),
			statsBtn,
		),
		historyContainer,
	)

	return container.NewBorder(nil, nil, nil, nil, container.NewVScroll(mainContent))
}

// performRecognition simule la reconnaissance d'une chanson
func (v *ShazamView) performRecognition(resultLabel *widget.Label, visualizationLabel *widget.Label) {
	resultLabel.SetText("🎧 Écoute en cours...")
	visualizationLabel.Show()

	// Animation de visualisation
	go func() {
		animations := []string{
			"🎵 ♪ ♫ ♬ 🎶",
			"♪ 🎵 ♬ 🎶 ♫",
			"♫ ♪ 🎵 ♬ 🎶",
			"♬ ♫ ♪ 🎵 🎶",
			"🎶 ♬ ♫ ♪ 🎵",
		}

		for i := 0; i < 8; i++ {
			visualizationLabel.SetText(animations[i%len(animations)])
			time.Sleep(250 * time.Millisecond)
		}

		// Simulation de reconnaissance
		if v.data != nil && len(v.data.Artists) > 0 {
			// Sélection aléatoire d'un artiste
			rand.Seed(time.Now().UnixNano())
			artist := v.data.Artists[rand.Intn(len(v.data.Artists))]

			// Ajouter à l'historique
			v.history = append(v.history, ShazamResult{
				Artist:    artist,
				Timestamp: time.Now(),
			})

			// Afficher le résultat
			result := fmt.Sprintf("✅ Chanson identifiée!\n\n"+
				"🎸 Artiste: %s\n"+
				"💿 Album: %s\n"+
				"📅 Année de création: %d\n"+
				"👥 Membres: %s\n\n"+
				"🎵 Cette chanson fait partie de leur répertoire!",
				artist.Name,
				artist.FirstAlbum,
				artist.CreationDate,
				strings.Join(artist.Members, ", "))

			resultLabel.SetText(result)
		} else {
			resultLabel.SetText("❌ Aucune chanson détectée. Réessayez.")
		}

		visualizationLabel.Hide()
	}()
}

// createHistoryView crée la vue de l'historique
func (v *ShazamView) createHistoryView() *fyne.Container {
	historyList := container.NewVBox()

	// Mise à jour de l'historique
	updateHistory := func() {
		historyList.Objects = nil

		if len(v.history) == 0 {
			historyList.Add(widget.NewLabel("Aucune reconnaissance effectuée"))
		} else {
			// Afficher les 5 derniers résultats
			start := 0
			if len(v.history) > 5 {
				start = len(v.history) - 5
			}

			for i := len(v.history) - 1; i >= start; i-- {
				result := v.history[i]
				historyCard := v.createHistoryCard(result)
				historyList.Add(historyCard)
			}
		}

		historyList.Refresh()
	}

	// Timer pour mettre à jour l'historique
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			updateHistory()
		}
	}()

	updateHistory()

	scroll := container.NewVScroll(historyList)
	scroll.SetMinSize(fyne.NewSize(0, 300))

	return container.NewBorder(nil, nil, nil, nil, scroll)
}

// createHistoryCard crée une carte pour l'historique
func (v *ShazamView) createHistoryCard(result ShazamResult) *fyne.Container {
	timeAgo := v.formatTimeAgo(result.Timestamp)

	nameLabel := widget.NewLabelWithStyle(
		fmt.Sprintf("🎵 %s", result.Artist.Name),
		fyne.TextAlignLeading,
		fyne.TextStyle{Bold: true},
	)

	infoLabel := widget.NewLabel(fmt.Sprintf("⏰ %s | 💿 %s", timeAgo, result.Artist.FirstAlbum))

	detailsBtn := widget.NewButton("📋 Détails", func() {
		v.showArtistDetails(result.Artist)
	})

	concertsBtn := widget.NewButton("🎤 Concerts", func() {
		v.showConcerts(result.Artist)
	})

	buttonsContainer := container.NewHBox(detailsBtn, concertsBtn)

	card := container.NewVBox(
		nameLabel,
		infoLabel,
		buttonsContainer,
		widget.NewSeparator(),
	)

	return container.NewPadded(card)
}

// formatTimeAgo formate le temps écoulé
func (v *ShazamView) formatTimeAgo(t time.Time) string {
	duration := time.Since(t)

	if duration < time.Minute {
		return "À l'instant"
	} else if duration < time.Hour {
		minutes := int(duration.Minutes())
		if minutes == 1 {
			return "Il y a 1 minute"
		}
		return fmt.Sprintf("Il y a %d minutes", minutes)
	} else if duration < 24*time.Hour {
		hours := int(duration.Hours())
		if hours == 1 {
			return "Il y a 1 heure"
		}
		return fmt.Sprintf("Il y a %d heures", hours)
	} else {
		days := int(duration.Hours() / 24)
		if days == 1 {
			return "Il y a 1 jour"
		}
		return fmt.Sprintf("Il y a %d jours", days)
	}
}

// showArtistDetails affiche les détails d'un artiste
func (v *ShazamView) showArtistDetails(artist models.Artist) {
	content := container.NewVBox(
		widget.NewLabelWithStyle(artist.Name, fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel(fmt.Sprintf("🎸 Année de création: %d", artist.CreationDate)),
		widget.NewLabel(fmt.Sprintf("💿 Premier album: %s", artist.FirstAlbum)),
		widget.NewLabel(fmt.Sprintf("👥 Nombre de membres: %d", len(artist.Members))),
		widget.NewSeparator(),
	)

	membersLabel := widget.NewLabelWithStyle("Membres:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true})
	content.Add(membersLabel)

	for i, member := range artist.Members {
		content.Add(widget.NewLabel(fmt.Sprintf("  %d. %s", i+1, member)))
	}

	closeBtn := widget.NewButton("Fermer", func() {})

	scroll := container.NewVScroll(content)
	scroll.SetMinSize(fyne.NewSize(400, 300))

	dialog := widget.NewModalPopUp(
		container.NewBorder(nil, container.NewCenter(closeBtn), nil, nil, scroll),
		v.window.Canvas(),
	)

	closeBtn.OnTapped = func() {
		dialog.Hide()
	}

	dialog.Show()
}

// showConcerts affiche les concerts d'un artiste
func (v *ShazamView) showConcerts(artist models.Artist) {
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

// showUserStats affiche les statistiques de l'utilisateur
func (v *ShazamView) showUserStats() {
	content := container.NewVBox(
		widget.NewLabelWithStyle("📊 Mes Statistiques Shazam", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewSeparator(),
		widget.NewLabel(fmt.Sprintf("🎵 Nombre de reconnaissances: %d", len(v.history))),
	)

	if len(v.history) > 0 {
		// Compter les artistes les plus identifiés
		artistCount := make(map[string]int)
		for _, result := range v.history {
			artistCount[result.Artist.Name]++
		}

		content.Add(widget.NewSeparator())
		content.Add(widget.NewLabelWithStyle("🏆 Artistes les plus identifiés:", fyne.TextAlignLeading, fyne.TextStyle{Bold: true}))

		for artist, count := range artistCount {
			content.Add(widget.NewLabel(fmt.Sprintf("  • %s: %d fois", artist, count)))
		}
	}

	closeBtn := widget.NewButton("Fermer", func() {})

	scroll := container.NewVScroll(content)
	scroll.SetMinSize(fyne.NewSize(400, 300))

	dialog := widget.NewModalPopUp(
		container.NewBorder(nil, container.NewCenter(closeBtn), nil, nil, scroll),
		v.window.Canvas(),
	)

	closeBtn.OnTapped = func() {
		dialog.Hide()
	}

	dialog.Show()
}
