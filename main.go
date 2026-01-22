package main

import (
	"fmt"
	"groupie-tracker/api"
	"groupie-tracker/models"
	"groupie-tracker/services"
	"groupie-tracker/ui"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// App représente l'application principale
type App struct {
	window        fyne.Window
	apiClient     *api.Client
	searchService *services.SearchService
	data          *models.APIData
	currentView   string

	// Vues
	spotifyView *ui.SpotifyView
	mapView     *ui.MapView
	shazamView  *ui.ShazamView
}

func main() {
	myApp := app.New()
	myApp.Settings().SetTheme(theme.DarkTheme())

	window := myApp.NewWindow("Groupie Tracker - Instagram Style")
	window.Resize(fyne.NewSize(1200, 800))

	application := &App{
		window:      window,
		apiClient:   api.NewClient(),
		currentView: "spotify",
	}

	// Créer l'interface principale
	mainUI := application.createMainUI()
	window.SetContent(mainUI)

	// Charger les données en arrière-plan
	go application.loadData()

	window.ShowAndRun()
}

// loadData charge toutes les données de l'API
func (a *App) loadData() {
	log.Println("🔄 Chargement des données...")

	data, err := a.apiClient.LoadAllData()
	if err != nil {
		log.Printf("❌ Erreur lors du chargement: %v\n", err)
		a.showError("Erreur de chargement des données. Veuillez vérifier votre connexion.")
		return
	}

	a.data = data
	a.searchService = services.NewSearchService(data)

	// Initialiser les vues
	a.spotifyView = ui.NewSpotifyView(a.window, a.searchService, a.data)
	a.mapView = ui.NewMapView(a.window, a.searchService, a.data)
	a.shazamView = ui.NewShazamView(a.window, a.searchService, a.data)

	log.Printf("✅ Données chargées: %d artistes\n", len(data.Artists))
}

// createMainUI crée l'interface principale
func (a *App) createMainUI() *fyne.Container {
	// Container pour le contenu principal
	mainContent := container.NewStack()

	// Navigation Instagram-style
	navigation := a.createNavigation(mainContent)

	// Vue par défaut (Spotify)
	loadingLabel := widget.NewLabel("⏳ Chargement des données de l'API...")
	loadingLabel.Alignment = fyne.TextAlignCenter
	mainContent.Objects = []fyne.CanvasObject{
		container.NewCenter(loadingLabel),
	}

	// Layout principal avec navigation à gauche
	return container.NewBorder(
		nil, nil,
		navigation,
		nil,
		mainContent,
	)
}

// createNavigation crée la barre de navigation latérale
func (a *App) createNavigation(mainContent *fyne.Container) *fyne.Container {
	// Logo/Titre
	title := widget.NewLabelWithStyle(
		"Groupie Tracker",
		fyne.TextAlignCenter,
		fyne.TextStyle{Bold: true},
	)

	separator1 := widget.NewSeparator()

	// Boutons de navigation
	spotifyBtn := widget.NewButton("", func() {
		a.switchView("spotify", mainContent)
	})
	spotifyBtn.Icon = theme.MediaMusicIcon()
	spotifyBtn.Importance = widget.HighImportance

	spotifyLabel := widget.NewLabel("Spotify")
	spotifyLabel.Alignment = fyne.TextAlignCenter

	mapBtn := widget.NewButton("", func() {
		a.switchView("map", mainContent)
	})
	mapBtn.Icon = theme.HomeIcon()
	mapBtn.Importance = widget.HighImportance

	mapLabel := widget.NewLabel("Carte")
	mapLabel.Alignment = fyne.TextAlignCenter

	shazamBtn := widget.NewButton("", func() {
		a.switchView("shazam", mainContent)
	})
	shazamBtn.Icon = theme.MediaRecordIcon()
	shazamBtn.Importance = widget.HighImportance

	shazamLabel := widget.NewLabel("Shazam")
	shazamLabel.Alignment = fyne.TextAlignCenter

	// Organisation verticale des boutons
	spotifyContainer := container.NewVBox(
		container.NewCenter(spotifyBtn),
		spotifyLabel,
	)

	mapContainer := container.NewVBox(
		container.NewCenter(mapBtn),
		mapLabel,
	)

	shazamContainer := container.NewVBox(
		container.NewCenter(shazamBtn),
		shazamLabel,
	)

	separator2 := widget.NewSeparator()

	// Informations en bas
	infoLabel := widget.NewLabel("API: Groupie Tracker")
	infoLabel.Alignment = fyne.TextAlignCenter
	infoLabel.TextStyle = fyne.TextStyle{Italic: true}

	// Container de navigation
	navContent := container.NewVBox(
		container.NewPadded(title),
		separator1,
		layout.NewSpacer(),
		container.NewPadded(spotifyContainer),
		widget.NewSeparator(),
		container.NewPadded(mapContainer),
		widget.NewSeparator(),
		container.NewPadded(shazamContainer),
		layout.NewSpacer(),
		separator2,
		container.NewPadded(infoLabel),
	)

	// Définir une largeur fixe pour la navigation
	navContainer := container.NewBorder(nil, nil, nil, nil, navContent)
	navContainer.Resize(fyne.NewSize(200, 0))

	return container.NewPadded(navContainer)
}

// switchView change de vue
func (a *App) switchView(view string, mainContent *fyne.Container) {
	a.currentView = view

	// Vérifier si les données sont chargées
	if a.data == nil {
		loadingLabel := widget.NewLabel("⏳ Chargement des données...")
		loadingLabel.Alignment = fyne.TextAlignCenter
		mainContent.Objects = []fyne.CanvasObject{
			container.NewCenter(loadingLabel),
		}
		mainContent.Refresh()
		return
	}

	var newView fyne.CanvasObject

	switch view {
	case "spotify":
		if a.spotifyView == nil {
			a.spotifyView = ui.NewSpotifyView(a.window, a.searchService, a.data)
		}
		newView = a.spotifyView.Render()

	case "map":
		if a.mapView == nil {
			a.mapView = ui.NewMapView(a.window, a.searchService, a.data)
		}
		newView = a.mapView.Render()

	case "shazam":
		if a.shazamView == nil {
			a.shazamView = ui.NewShazamView(a.window, a.searchService, a.data)
		}
		newView = a.shazamView.Render()

	default:
		newView = container.NewCenter(widget.NewLabel("Vue non disponible"))
	}

	mainContent.Objects = []fyne.CanvasObject{newView}
	mainContent.Refresh()

	log.Printf("📱 Vue changée: %s\n", view)
}

// showError affiche un message d'erreur
func (a *App) showError(message string) {
	dialog := widget.NewModalPopUp(
		container.NewVBox(
			widget.NewLabel("❌ Erreur"),
			widget.NewSeparator(),
			widget.NewLabel(message),
			widget.NewButton("OK", func() {}),
		),
		a.window.Canvas(),
	)

	dialog.Show()
	fmt.Println(message)
}
