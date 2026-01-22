#!/bin/bash

# Script d'installation automatique pour Groupie Tracker
echo "🚀 Installation de Groupie Tracker"
echo "===================================="

# Vérifier que Go est installé
if ! command -v go &> /dev/null
then
    echo "❌ Go n'est pas installé. Veuillez l'installer depuis https://go.dev/dl/"
    exit 1
fi

echo "✅ Go est installé: $(go version)"

# Vérifier que GCC est installé (pour Fyne)
if ! command -v gcc &> /dev/null
then
    echo "⚠️  GCC n'est pas installé. Fyne nécessite GCC."
    echo "   Linux: sudo apt-get install gcc libgl1-mesa-dev xorg-dev"
    echo "   macOS: xcode-select --install"
    echo "   Windows: Télécharger depuis https://jmeubank.github.io/tdm-gcc/download/"
    exit 1
fi

echo "✅ GCC est installé: $(gcc --version | head -n 1)"

# Créer la structure des dossiers
echo ""
echo "📁 Création de la structure des dossiers..."
mkdir -p models api services ui

# Initialiser le module Go
echo ""
echo "📦 Initialisation du module Go..."
go mod init groupie-tracker 2>/dev/null || echo "   Module déjà initialisé"

# Installer les dépendances
echo ""
echo "📥 Installation des dépendances..."
go get fyne.io/fyne/v2

# Télécharger toutes les dépendances
echo ""
echo "🔄 Téléchargement des dépendances..."
go mod tidy

echo ""
echo "✅ Installation terminée!"
echo ""
echo "📝 Prochaines étapes:"
echo "   1. Copiez tous les fichiers .go dans leurs dossiers respectifs"
echo "   2. Lancez l'application avec: go run main.go"
echo "   3. Ou compilez avec: go build -o groupie-tracker"
echo ""
echo "📚 Structure des fichiers:"
echo "   ├── main.go"
echo "   ├── models/models.go"
echo "   ├── api/client.go"
echo "   ├── services/search.go"
echo "   ├── ui/spotify_view.go"
echo "   ├── ui/map_view.go"
echo "   └── ui/shazam_view.go"
echo ""
echo "🎉 Bon développement!"