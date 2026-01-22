@echo off
title Groupie Tracker
color 0A

echo.
echo  ========================================
echo   GROUPIE TRACKER - Instagram Style
echo  ========================================
echo.

REM Vérifier si go.mod existe
if not exist "go.mod" (
    echo [INFO] Initialisation du projet...
    go mod init groupie-tracker
    if errorlevel 1 (
        echo [ERREUR] Echec de l'initialisation
        pause
        exit /b 1
    )
    echo [OK] Module initialise
)

REM Installer/Mettre à jour les dépendances
echo [INFO] Verification des dependances...
go get fyne.io/fyne/v2
go mod tidy

if errorlevel 1 (
    echo.
    echo [ERREUR] Probleme avec les dependances
    echo.
    echo Verifiez que:
    echo  - Go est installe (go version)
    echo  - GCC est installe (gcc --version)
    echo  - Vous avez une connexion internet
    echo.
    pause
    exit /b 1
)

echo [OK] Dependances installees
echo.

REM Compiler l'application
echo [INFO] Compilation de l'application...
go build -o groupie-tracker.exe

if errorlevel 1 (
    echo.
    echo [ERREUR] Echec de la compilation
    echo.
    echo Verifiez les erreurs ci-dessus et consultez TROUBLESHOOTING.md
    echo.
    pause
    exit /b 1
)

echo [OK] Compilation reussie
echo.

REM Lancer l'application
echo [INFO] Lancement de Groupie Tracker...
echo.
echo =========================================
echo   Application demarree !
echo   Consultez la fenetre de l'application
echo =========================================
echo.

groupie-tracker.exe

if errorlevel 1 (
    echo.
    echo [ERREUR] L'application s'est arretee avec une erreur
    pause
)

echo.
echo Application fermee.
pause