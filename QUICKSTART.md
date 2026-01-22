# 🚀 Démarrage Rapide - Groupie Tracker

## ⚡ Installation en 5 Minutes

### Étape 1 : Vérifier les Prérequis (2 min)

#### Vérifier Go
```bash
go version
```
✅ Devrait afficher `go1.21.x` ou supérieur  
❌ Si non installé → https://go.dev/dl/

#### Vérifier GCC
```bash
gcc --version
```
✅ Devrait afficher la version de GCC  
❌ Si non installé → https://jmeubank.github.io/tdm-gcc/download/ (Windows)

### Étape 2 : Créer la Structure (1 min)

```bash
# Créer le dossier principal
mkdir groupie-tracker
cd groupie-tracker

# Créer les sous-dossiers
mkdir models api services ui
```

### Étape 3 : Copier les Fichiers (1 min)

Copiez chaque fichier dans son dossier :

```
groupie-tracker/
├── main.go                    ← À la racine
├── go.mod                     ← À la racine
├── models/
│   └── models.go              ← Dans models/
├── api/
│   └── client.go              ← Dans api/
├── services/
│   └── search.go              ← Dans services/
└── ui/
    ├── spotify_view.go        ← Dans ui/
    ├── map_view.go            ← Dans ui/
    └── shazam_view.go         ← Dans ui/
```

### Étape 4 : Installer les Dépendances (1 min)

```bash
# Initialiser le module
go mod init groupie-tracker

# Installer Fyne
go get fyne.io/fyne/v2

# Télécharger toutes les dépendances
go mod tidy
```

### Étape 5 : Lancer l'Application (< 1 min)

#### Option A : Mode Développement
```bash
go run main.go
```

#### Option B : Compiler puis Exécuter
```bash
# Compiler
go build -o groupie-tracker.exe

# Lancer
./groupie-tracker.exe
```

#### Option C : Script Automatique (Windows)
```bash
# Double-cliquer sur run.bat
# Ou dans le terminal :
run.bat
```

---

## 🎯 Navigation dans l'Application

Une fois l'application lancée :

### 🎵 Vue Spotify (Bouton gauche)
- **Recherche** : Tapez dans la barre de recherche
- **Suggestions** : Apparaissent automatiquement en temps réel
- **Détails** : Cliquez sur "📋 Détails" pour voir les infos complètes
- **Concerts** : Cliquez sur "🎤 Voir concerts" pour la liste des concerts

### 🗺️ Vue Carte (Bouton milieu)
- **Recherche** : Recherchez un artiste spécifique
- **Statistiques** : Cliquez sur "📊 Statistiques" pour les stats globales
- **Carte** : Cliquez sur "🗺️ Voir sur la carte" pour les détails du lieu

### 🎤 Vue Shazam (Bouton droit)
- **Reconnaissance** : Cliquez sur "🎧 Écouter et Identifier"
- **Historique** : Voir toutes vos reconnaissances précédentes
- **Statistiques** : Cliquez sur "📊 Mes statistiques" pour vos stats perso

---

## 🔧 Résolution de Problèmes Rapide

### Problème : Lignes Fyne en rouge

**Solution rapide** :
```bash
go get fyne.io/fyne/v2
go mod tidy
```
Puis redémarrez votre IDE (VSCode : Ctrl+Shift+P → "Reload Window")

### Problème : "gcc not found"

**Solution Windows** :
1. Télécharger : https://jmeubank.github.io/tdm-gcc/download/
2. Installer `tdm64-gcc-10.3.0-2.exe`
3. Ajouter au PATH : `C:\TDM-GCC-64\bin`
4. Redémarrer le terminal

### Problème : "cannot use scroll as *fyne.Container"

**Solution** : Les fichiers ont été corrigés. Assurez-vous d'avoir la dernière version de tous les fichiers .go

### Problème : Données ne se chargent pas

**Vérifications** :
1. Connexion internet active ?
2. L'API est accessible : https://groupietrackers.herokuapp.com/api/artists
3. Regardez les logs dans le terminal

---

## 📝 Commandes Essentielles

### Développement
```bash
go run main.go              # Lancer sans compiler
go build                    # Compiler
go build -o mon_app.exe     # Compiler avec nom personnalisé
```

### Maintenance
```bash
go mod tidy                 # Nettoyer les dépendances
go clean                    # Nettoyer les fichiers compilés
go fmt ./...                # Formater le code
```

### Diagnostic
```bash
go version                  # Version de Go
go env                      # Variables d'environnement Go
go list -m all              # Lister les modules
```

---

## 🎨 Personnalisation Rapide

### Changer le Thème
Dans `main.go`, ligne ~27 :
```go
myApp.Settings().SetTheme(theme.LightTheme())  // Thème clair
```

### Changer la Taille de Fenêtre
Dans `main.go`, ligne ~30 :
```go
window.Resize(fyne.NewSize(1400, 900))  // Plus grand
```

### Activer les Logs Détaillés
Dans `main.go`, après les imports :
```go
import (
    "log"
    // ... autres imports
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    // ... reste du code
}
```

---

## 📚 Ressources Utiles

- **Documentation Fyne** : https://developer.fyne.io/
- **API Groupie Tracker** : https://groupietrackers.herokuapp.com/api
- **Guide de dépannage complet** : Voir `TROUBLESHOOTING.md`
- **Documentation Go** : https://go.dev/doc/

---

## ✅ Checklist de Vérification

Avant de demander de l'aide, vérifiez :

- [ ] Go version 1.21+ installé (`go version`)
- [ ] GCC installé (`gcc --version`)
- [ ] Tous les fichiers .go dans les bons dossiers
- [ ] `go mod tidy` exécuté sans erreur
- [ ] `go build` compile sans erreur
- [ ] Connexion internet disponible
- [ ] Aucune erreur dans les logs du terminal

---

## 🎉 Première Utilisation

1. **Lancez l'application**
2. **Attendez le chargement** (5-10 secondes)
3. **Explorez la vue Spotify** (vue par défaut)
4. **Testez la recherche** : Tapez "Queen" ou "Michael"
5. **Cliquez sur un artiste** pour voir les détails
6. **Changez de vue** avec les boutons de navigation
7. **Testez Shazam** pour une simulation de reconnaissance

---

## 💡 Astuces Pro

### Raccourcis Utiles
- `Ctrl+C` dans le terminal pour arrêter l'app
- `Ctrl+Shift+P` dans VSCode → "Go: Restart Language Server"
- Utilisez `run.bat` (Windows) pour un lancement automatique

### Performance
- La première compilation peut prendre 30-60 secondes
- Les lancements suivants sont beaucoup plus rapides
- Les données de l'API se chargent en arrière-plan

### Développement
- Modifiez le code pendant que l'app tourne
- Relancez avec `go run main.go` pour voir les changements
- Utilisez `go fmt` pour formater automatiquement le code

---

**Vous êtes maintenant prêt ! 🚀**

Si vous rencontrez des problèmes, consultez `TROUBLESHOOTING.md` pour des solutions détaillées.