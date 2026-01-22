# 📋 Liste de Vérification des Fichiers

## ✅ Vérification Complète de la Structure

Utilisez cette liste pour vous assurer que tous les fichiers sont présents et corrects.

---

## 📁 Structure des Dossiers

```
groupie-tracker/
├── 📄 main.go
├── 📄 go.mod
├── 📄 go.sum (généré automatiquement)
├── 📄 README.md
├── 📄 QUICKSTART.md
├── 📄 TROUBLESHOOTING.md
├── 📄 FILES_CHECKLIST.md (ce fichier)
├── 📄 run.bat (Windows)
├── 📄 check.bat (Windows)
│
├── 📂 models/
│   └── 📄 models.go
│
├── 📂 api/
│   └── 📄 client.go
│
├── 📂 services/
│   └── 📄 search.go
│
└── 📂 ui/
    ├── 📄 spotify_view.go
    ├── 📄 map_view.go
    └── 📄 shazam_view.go
```

---

## 📄 Fichiers Requis (Obligatoires)

### ✅ Racine du Projet

- [ ] **main.go** - Point d'entrée de l'application
  - Package : `package main`
  - Imports requis : `groupie-tracker/api`, `groupie-tracker/models`, `groupie-tracker/services`, `groupie-tracker/ui`
  - Fonction principale : `func main()`

- [ ] **go.mod** - Configuration du module
  - Première ligne : `module groupie-tracker`
  - Doit contenir : `require fyne.io/fyne/v2 v2.4.3`

### ✅ Dossier models/

- [ ] **models/models.go**
  - Package : `package models`
  - Structures : `Artist`, `Location`, `Date`, `Relation`, `APIData`, `Concert`, `SearchResult`
  - Aucun import externe requis

### ✅ Dossier api/

- [ ] **client.go**
  - Package : `package api`
  - Imports requis : `groupie-tracker/models`
  - Fonctions : `NewClient()`, `GetArtists()`, `GetLocations()`, `GetDates()`, `GetRelations()`, `LoadAllData()`
  - Constante : `BaseURL = "https://groupietrackers.herokuapp.com/api"`

### ✅ Dossier services/

- [ ] **search.go**
  - Package : `package services`
  - Imports requis : `groupie-tracker/models`
  - Fonctions : `NewSearchService()`, `SearchArtists()`, `SearchByMember()`, `UniversalSearch()`, `FormatLocation()`

### ✅ Dossier ui/

- [ ] **spotify_view.go**
  - Package : `package ui`
  - Imports requis : `groupie-tracker/models`, `groupie-tracker/services`, `fyne.io/fyne/v2`
  - Structure : `SpotifyView` avec méthode `Render()`

- [ ] **map_view.go**
  - Package : `package ui`
  - Imports requis : `groupie-tracker/models`, `groupie-tracker/services`, `fyne.io/fyne/v2`, `strings`
  - Structure : `MapView` avec méthode `Render()`

- [ ] **shazam_view.go**
  - Package : `package ui`
  - Imports requis : `groupie-tracker/models`, `groupie-tracker/services`, `fyne.io/fyne/v2`
  - Structure : `ShazamView` avec méthode `Render()`
  - **IMPORTANT** : Les fonctions doivent retourner `*fyne.Container`, pas `*container.Scroll`

---

## 📄 Fichiers Optionnels (Recommandés)

- [ ] **README.md** - Documentation complète
- [ ] **QUICKSTART.md** - Guide de démarrage rapide
- [ ] **TROUBLESHOOTING.md** - Guide de dépannage
- [ ] **run.bat** - Script de lancement Windows
- [ ] **check.bat** - Script de vérification Windows

---

## 🔍 Vérifications par Fichier

### main.go - Points Critiques

```go
// ✅ Imports corrects
import (
    "groupie-tracker/api"
    "groupie-tracker/models"
    "groupie-tracker/services"
    "groupie-tracker/ui"
    
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
)

// ✅ Structure App
type App struct {
    window        fyne.Window
    apiClient     *api.Client
    searchService *services.SearchService
    data          *models.APIData
    currentView   string
    spotifyView   *ui.SpotifyView
    mapView       *ui.MapView
    shazamView    *ui.ShazamView
}

// ✅ Fonction main
func main() {
    myApp := app.New()
    myApp.Settings().SetTheme(theme.DarkTheme())
    window := myApp.NewWindow("Groupie Tracker - Instagram Style")
    window.Resize(fyne.NewSize(1200, 800))
    // ...
}
```

### models/models.go - Points Critiques

```go
// ✅ Package correct
package models

// ✅ Structures principales
type Artist struct {
    ID           int      `json:"id"`
    Image        string   `json:"image"`
    Name         string   `json:"name"`
    Members      []string `json:"members"`
    CreationDate int      `json:"creationDate"`
    FirstAlbum   string   `json:"firstAlbum"`
}

type APIData struct {
    Artists   []Artist
    Locations []Location
    Dates     []Date
    Relations []Relation
}
```

### api/client.go - Points Critiques

```go
// ✅ Package et imports
package api

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "groupie-tracker/models"
)

// ✅ Constante API
const BaseURL = "https://groupietrackers.herokuapp.com/api"

// ✅ Fonction LoadAllData
func (c *Client) LoadAllData() (*models.APIData, error) {
    // Doit charger: Artists, Relations, Locations, Dates
}
```

### services/search.go - Points Critiques

```go
// ✅ Package et imports
package services

import (
    "fmt"
    "groupie-tracker/models"
    "strings"
)

// ✅ Fonction FormatLocation (IMPORTANTE)
func FormatLocation(location string) string {
    location = strings.ReplaceAll(location, "-", ", ")
    location = strings.ReplaceAll(location, "_", " ")
    // ...
    return strings.Join(parts, ", ")
}
```

### ui/shazam_view.go - Points Critiques ⚠️

```go
// ✅ Render() doit retourner *fyne.Container
func (v *ShazamView) Render() *fyne.Container {
    // ...
    mainContent := container.NewVBox(...)
    
    // ✅ CORRECT : Wrapper le scroll dans un container
    return container.NewBorder(nil, nil, nil, nil, 
        container.NewVScroll(mainContent))
    
    // ❌ INCORRECT : Retourner directement le scroll
    // return container.NewVScroll(mainContent)
}

// ✅ createHistoryView() doit retourner *fyne.Container
func (v *ShazamView) createHistoryView() *fyne.Container {
    // ...
    scroll := container.NewVScroll(historyList)
    
    // ✅ CORRECT
    return container.NewBorder(nil, nil, nil, nil, scroll)
    
    // ❌ INCORRECT
    // return scroll
}
```

---

## 🧪 Tests de Vérification

### Test 1 : Vérification des Imports

```bash
# Vérifier que tous les fichiers compilent
go build ./...

# Si erreur "package not found", vérifier :
# 1. go.mod existe et contient "module groupie-tracker"
# 2. Tous les imports utilisent "groupie-tracker/xxx"
# 3. go mod tidy a été exécuté
```

### Test 2 : Vérification de la Structure

```bash
# Windows
dir /s /b *.go

# Linux/Mac
find . -name "*.go" -type f

# Devrait afficher 8 fichiers .go :
# - main.go
# - models/models.go
# - api/client.go
# - services/search.go
# - ui/spotify_view.go
# - ui/map_view.go
# - ui/shazam_view.go
```

### Test 3 : Compilation Complète

```bash
# Compiler
go build -v

# Devrait afficher :
# groupie-tracker/models
# groupie-tracker/api
# groupie-tracker/services
# groupie-tracker/ui
# groupie-tracker

# Si succès : ✅
# Si erreur : Voir TROUBLESHOOTING.md
```

### Test 4 : Vérification des Dépendances

```bash
go list -m all | grep fyne

# Devrait afficher :
# fyne.io/fyne/v2 v2.4.3
```

---

## 🚨 Erreurs Communes à Vérifier

### ❌ Erreur : "cannot use scroll as *fyne.Container"

**Cause** : Fonction retourne `*container.Scroll` au lieu de `*fyne.Container`

**Fichiers à vérifier** :
- `ui/shazam_view.go` ligne ~89 et ~186

**Solution** :
```go
// ❌ INCORRECT
return scroll

// ✅ CORRECT
return container.NewBorder(nil, nil, nil, nil, scroll)
```

### ❌ Erreur : "undefined: strings"

**Cause** : Import manquant

**Fichier à vérifier** : `ui/map_view.go`

**Solution** : Ajouter `"strings"` dans les imports

### ❌ Erreur : "package groupie-tracker/xxx is not in GOROOT"

**Cause** : Module Go non initialisé

**Solution** :
```bash
go mod init groupie-tracker
go mod tidy
```

---

## ✅ Checklist Finale

Avant de lancer l'application, vérifiez :

- [ ] Tous les fichiers .go sont présents (8 fichiers)
- [ ] go.mod existe avec `module groupie-tracker`
- [ ] Tous les packages sont corrects (`package main`, `package models`, etc.)
- [ ] Tous les imports utilisent `groupie-tracker/xxx`
- [ ] `go mod tidy` exécuté sans erreur
- [ ] `go build` compile sans erreur
- [ ] GCC est installé (pour Fyne)
- [ ] Connexion internet disponible (pour l'API)

Si tous les points sont cochés → **Vous êtes prêt à lancer ! 🚀**

```bash
go run main.go
```

---

## 📞 Aide Supplémentaire

Si un élément de cette checklist échoue :
1. Consultez `TROUBLESHOOTING.md` pour le problème spécifique
2. Vérifiez que vous êtes dans le bon dossier (`pwd` ou `cd`)
3. Relancez `go mod tidy`
4. Redémarrez votre terminal/IDE

**Commande de diagnostic rapide** :
```bash
go version && gcc --version && go mod tidy && go build
```

Si cette commande réussit → Tout est OK ✅