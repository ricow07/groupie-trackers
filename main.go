package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	fmt.Println("🚀 Démarrage de Groupie Trackers...")

	myApp := app.New()
	mainWindow := myApp.NewWindow("🎵 Groupie Trackers - Concert Explorer")
	mainWindow.Resize(fyne.NewSize(1000, 700))

	// Variables pour l'UI
	artistsContainer := container.NewVBox()

	// ===== ZONE DE CHARGEMENT =====
	artistsContainer.Add(CreateLoadingWidget())

	// ===== ZONE DÉTAILS =====
	detailsContainer := container.NewVBox()
	detailsScroll := container.NewVScroll(detailsContainer)
	detailsScroll.SetMinSize(fyne.NewSize(400, 600))

	// ===== ZONE ARTISTES =====
	artistsScroll := container.NewVScroll(artistsContainer)
	artistsScroll.SetMinSize(fyne.NewSize(550, 600))

	// Fonction pour rafraîchir la liste des artistes
	refreshArtists := func(artists []Artist) {
		artistsContainer.Objects = nil

		if len(artists) == 0 {
			artistsContainer.Add(widget.NewLabel("❌ Aucun artiste trouvé"))
			return
		}

		for _, a := range artists {
			artist := a // Copie pour éviter les problèmes de pointeurs
			card := CreateArtistCard(&artist, func(selected *Artist) {
				// Mettre à jour la zone de détails
				detailsContainer.Objects = nil
				detailsContainer.Add(CreateDetailWindow(selected, myApp))
				detailsContainer.Refresh()
			})
			artistsContainer.Add(card)
		}
	}

	// ===== BARRE DE RECHERCHE & FILTRES =====
	search, periodSelect, searchBar := CreateSearchBar()

	// Fonction de recherche
	doSearch := func() {
		query := search.Text
		results := AllArtists

		// Filtrer par recherche
		if query != "" {
			results = SearchArtists(query)
		}

		// Filtrer par période
		period := periodSelect.Selected
		switch period {
		case "Avant 1980":
			results = FilterByCreationDate(results, 0, 1979)
		case "1980-1999":
			results = FilterByCreationDate(results, 1980, 1999)
		case "2000 et après":
			results = FilterByCreationDate(results, 2000, 2100)
		}

		refreshArtists(results)
	}

	search.OnChanged = func(s string) {
		doSearch()
	}

	periodSelect.OnChanged = func(s string) {
		doSearch()
	}

	// ===== BOUTON RÉINITIALISER =====
	resetBtn := widget.NewButton("🔄 Réinitialiser", func() {
		search.SetText("")
		periodSelect.SetSelected("Toutes les périodes")
		refreshArtists(AllArtists)
	})

	searchContainer := container.NewVBox(
		searchBar,
		resetBtn,
	)

	// ===== LAYOUT PRINCIPAL =====
	topLayout := container.NewVBox(
		widget.NewLabel("🎵 GROUPIE TRACKERS - Découvrez vos artistes préférés"),
		widget.NewSeparator(),
		searchContainer,
	)

	mainLayout := container.NewHBox(
		artistsScroll,
		widget.NewSeparator(),
		detailsScroll,
	)

	fullLayout := container.NewBorder(
		topLayout,
		nil,
		nil,
		nil,
		mainLayout,
	)

	mainWindow.SetContent(fullLayout)

	// ===== CHARGEMENT DES DONNÉES =====
	go func() {
		fmt.Println("📡 Chargement de l'API...")
		err := FetchAllData()

		if err != nil {
			fmt.Println("❌ Erreur API:", err)
			artistsContainer.Objects = nil
			artistsContainer.Add(CreateErrorWidget(err.Error()))
			artistsContainer.Refresh()
		} else {
			fmt.Printf("✅ %d artistes chargés avec succès!\n", len(AllArtists))
			refreshArtists(AllArtists)
		}
	}()

	mainWindow.ShowAndRun()
}
