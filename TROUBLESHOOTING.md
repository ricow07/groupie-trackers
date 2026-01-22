# Guide de Dépannage - Groupie Tracker

## 🔴 Erreurs Courantes et Solutions

### 1. Lignes Fyne en rouge dans l'éditeur

**Problème** : Toutes les lignes avec `fyne` apparaissent en rouge dans VSCode ou votre IDE.

**Solutions** :

#### Solution A : Installer les dépendances
```bash
# Dans le dossier du projet
go mod init groupie-tracker
go get fyne.io/fyne/v2
go mod tidy
```

#### Solution B : Redémarrer le serveur Go (VSCode)
1. Appuyez sur `Ctrl+Shift+P` (ou `Cmd+Shift+P` sur Mac)
2. Tapez "Go: Restart Language Server"
3. Appuyez sur Entrée

#### Solution C : Installer l'extension Go pour VSCode
1. Ouvrez VSCode
2. Allez dans Extensions (Ctrl+Shift+X)
3. Cherchez "Go" (par Go Team at Google)
4. Installez l'extension
5. Redémarrez VSCode

#### Solution D : Définir le GOPATH
```bash
# Windows (PowerShell)
$env:GOPATH = "$HOME\go"

# Linux/Mac
export GOPATH=$HOME/go
```

### 2. Erreur "cannot use scroll as *fyne.Container"

**Problème** : 
```
cannot use scroll (variable of type *container.Scroll) as *fyne.Container value in return statement
```

**Solution** : Les fichiers ont été corrigés. Assurez-vous d'avoir la dernière version de `ui/shazam_view.go` qui retourne :
```go
return container.NewBorder(nil, nil, nil, nil, scroll)
```
au lieu de :
```go
return scroll
```

### 3. Erreur "gcc: command not found"

**Problème** : Fyne nécessite GCC pour compiler.

**Solutions par OS** :

#### Windows
1. Téléchargez TDM-GCC : https://jmeubank.github.io/tdm-gcc/download/
2. Installez `tdm64-gcc-10.3.0-2.exe`
3. Ajoutez `C:\TDM-GCC-64\bin` au PATH :
   - Windows Key → "variables d'environnement"
   - Modifier la variable PATH
   - Ajouter `C:\TDM-GCC-64\bin`
4. Redémarrez votre terminal

#### Linux (Ubuntu/Debian)
```bash
sudo apt-get update
sudo apt-get install gcc libgl1-mesa-dev xorg-dev
```

#### macOS
```bash
xcode-select --install
```

### 4. Erreur "package not found"

**Problème** : 
```
package groupie-tracker/models is not in GOROOT
```

**Solution** :
```bash
# Vérifier la structure des dossiers
ls -la models/ api/ services/ ui/

# Si des dossiers manquent, les créer
mkdir -p models api services ui

# Vérifier que go.mod existe
cat go.mod

# Si go.mod n'existe pas
go mod init groupie-tracker

# Réinstaller les dépendances
go mod tidy
```

### 5. Erreur de compilation au lancement

**Problème** : Erreurs de syntaxe ou imports manquants.

**Solution** :
```bash
# Vérifier les erreurs
go build

# Si erreur d'imports manquants
go get fyne.io/fyne/v2
go mod tidy

# Nettoyer le cache si nécessaire
go clean -modcache
go mod download
```

### 6. L'application se lance mais ne charge pas les données

**Problème** : L'écran reste sur "Chargement des données..."

**Solutions** :

#### Vérifier la connexion internet
```bash
# Tester l'API
curl https://groupietrackers.herokuapp.com/api/artists
```

#### Vérifier les logs
- Regardez le terminal où vous avez lancé l'application
- Recherchez les messages d'erreur

#### Augmenter le timeout
Dans `api/client.go`, ajoutez un timeout plus long :
```go
client:  &http.Client{
    Timeout: 30 * time.Second,
},
```

### 7. Erreur "go: go.mod file not found"

**Problème** : Vous n'êtes pas dans le bon dossier ou go.mod n'existe pas.

**Solution** :
```bash
# Vérifier où vous êtes
pwd

# Aller dans le bon dossier
cd groupie-tracker

# Initialiser go.mod si nécessaire
go mod init groupie-tracker
```

### 8. Import manquant "strings" dans map_view.go

**Problème** :
```
undefined: strings
```

**Solution** : Ajoutez l'import dans `ui/map_view.go` :
```go
import (
    "fmt"
    "groupie-tracker/models"
    "groupie-tracker/services"
    "strings"  // ← Ajouter cette ligne
    "time"
    ...
)
```

### 9. L'application compile mais crashe au démarrage

**Problème** : Panic ou erreur au lancement.

**Solutions** :

#### Vérifier les logs complets
```bash
go run main.go 2>&1 | tee error.log
```

#### Activer le mode debug
Dans `main.go`, ajoutez après l'import :
```go
import (
    "log"
    "os"
)

func main() {
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    log.SetOutput(os.Stdout)
    // ... reste du code
}
```

#### Vérifier les permissions réseau
- Sur Windows : Autoriser l'application dans le pare-feu
- Sur Linux : Vérifier les permissions réseau

### 10. Erreurs de type "undefined: xxx"

**Problème** : Variables ou fonctions non définies.

**Solution** : Vérifier que tous les imports sont corrects :

#### main.go
```go
import (
    "groupie-tracker/api"
    "groupie-tracker/models"
    "groupie-tracker/services"
    "groupie-tracker/ui"
    // ...
)
```

#### ui/spotify_view.go, map_view.go, shazam_view.go
```go
import (
    "groupie-tracker/models"
    "groupie-tracker/services"
    // ...
)
```

## 🛠️ Commandes Utiles de Diagnostic

### Vérifier la version de Go
```bash
go version
# Devrait afficher: go version go1.21.x ou supérieur
```

### Vérifier GCC
```bash
gcc --version
# Devrait afficher la version de GCC
```

### Lister les dépendances
```bash
go list -m all
```

### Vérifier les imports manquants
```bash
go mod tidy
```

### Nettoyer et reconstruire
```bash
go clean
go build
```

### Compiler avec informations détaillées
```bash
go build -v
```

### Tester la compilation sans créer d'exécutable
```bash
go build -o /dev/null   # Linux/Mac
go build -o NUL         # Windows
```

## 📁 Vérification de la Structure

Votre structure devrait ressembler à ça :

```
groupie-tracker/
├── main.go
├── go.mod
├── go.sum
├── README.md
├── models/
│   └── models.go
├── api/
│   └── client.go
├── services/
│   └── search.go
└── ui/
    ├── spotify_view.go
    ├── map_view.go
    └── shazam_view.go
```

### Vérifier rapidement (Linux/Mac)
```bash
find . -name "*.go" -type f
```

### Vérifier rapidement (Windows)
```cmd
dir /s /b *.go
```

## 🔄 Réinitialisation Complète

Si rien ne fonctionne, réinitialisez complètement :

```bash
# 1. Sauvegarder vos fichiers .go

# 2. Supprimer go.mod et go.sum
rm go.mod go.sum

# 3. Nettoyer le cache
go clean -modcache

# 4. Réinitialiser
go mod init groupie-tracker
go get fyne.io/fyne/v2
go mod tidy

# 5. Tester la compilation
go build
```

## 📞 Obtenir de l'Aide

Si le problème persiste :

1. **Vérifier les logs** : Lisez attentivement les messages d'erreur
2. **Vérifier la version de Go** : `go version` (minimum 1.21)
3. **Vérifier GCC** : `gcc --version`
4. **Copier l'erreur exacte** : Notez le message d'erreur complet
5. **Vérifier la structure** : Tous les fichiers sont-ils au bon endroit ?

## ✅ Checklist de Vérification Complète

- [ ] Go version 1.21+ installé
- [ ] GCC installé et dans le PATH
- [ ] Tous les fichiers .go créés dans les bons dossiers
- [ ] go.mod existe et contient `module groupie-tracker`
- [ ] `go mod tidy` exécuté sans erreur
- [ ] `go build` compile sans erreur
- [ ] Connexion internet disponible
- [ ] Pare-feu autorise l'application

Si tout est coché ✅, l'application devrait fonctionner parfaitement !