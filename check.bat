@echo off
echo ====================================
echo Verification de Groupie Tracker
echo ====================================
echo.

echo [1/5] Verification de Go...
go version
if errorlevel 1 (
    echo ERREUR: Go n'est pas installe
    pause
    exit /b 1
)
echo OK
echo.

echo [2/5] Verification de GCC...
gcc --version
if errorlevel 1 (
    echo ERREUR: GCC n'est pas installe
    echo Installez TDM-GCC depuis https://jmeubank.github.io/tdm-gcc/download/
    pause
    exit /b 1
)
echo OK
echo.

echo [3/5] Verification de la structure des dossiers...
if not exist "models" (
    echo Creation du dossier models...
    mkdir models
)
if not exist "api" (
    echo Creation du dossier api...
    mkdir api
)
if not exist "services" (
    echo Creation du dossier services...
    mkdir services
)
if not exist "ui" (
    echo Creation du dossier ui...
    mkdir ui
)
echo OK
echo.

echo [4/5] Verification des fichiers requis...
set MISSING=0

if not exist "main.go" (
    echo MANQUANT: main.go
    set MISSING=1
)
if not exist "models\models.go" (
    echo MANQUANT: models\models.go
    set MISSING=1
)
if not exist "api\client.go" (
    echo MANQUANT: api\client.go
    set MISSING=1
)
if not exist "services\search.go" (
    echo MANQUANT: services\search.go
    set MISSING=1
)
if not exist "ui\spotify_view.go" (
    echo MANQUANT: ui\spotify_view.go
    set MISSING=1
)
if not exist "ui\map_view.go" (
    echo MANQUANT: ui\map_view.go
    set MISSING=1
)
if not exist "ui\shazam_view.go" (
    echo MANQUANT: ui\shazam_view.go
    set MISSING=1
)

if %MISSING%==1 (
    echo.
    echo ERREUR: Certains fichiers sont manquants
    pause
    exit /b 1
)
echo OK - Tous les fichiers sont presents
echo.

echo [5/5] Initialisation et installation des dependances...
if not exist "go.mod" (
    echo Initialisation du module Go...
    go mod init groupie-tracker
)

echo Installation de Fyne...
go get fyne.io/fyne/v2

echo Telechargement des dependances...
go mod tidy

echo.
echo ====================================
echo Verification terminee avec succes!
echo ====================================
echo.
echo Pour compiler: go build -o groupie-tracker.exe
echo Pour lancer: go run main.go
echo.
pause