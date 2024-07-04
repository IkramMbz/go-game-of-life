package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	width    = 100 // Taille de la grille en largeur
	height   = 100 // Taille de la grille en hauteur
	cellSize = 10  // Taille initiale de chaque cellule pour l'affichage
	offsetX  = 0.0
	offsetY  = 0.0
	slider  *Slider
)

func main() {
	// Demander à l'utilisateur d'entrer les dimensions de la grille
	fmt.Print("Entrez la largeur de la grille: ")
	fmt.Scan(&width)
	fmt.Print("Entrez la hauteur de la grille: ")
	fmt.Scan(&height)

	//Vérifier les entrées de l'utilisateur
	if height != width {
		fmt.Println("Erreur de taille : la largeur et la hauteur doivent être égales.")
		return
	}

	if height < 15 || height > 60 &&  width < 15 || width > 90 {
		fmt.Println("Erreur de taille : la taille doit être inférieur à 90 ou supérieur à 15.")
		return
	}

	// Définir la taille de la fenêtre en fonction des dimensions de la grille et de la taille des cellules
	windowWidth := width * cellSize
	windowHeight := height * cellSize + 60 // Ajouter 60 pixels pour la barre des statistiques

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("Game of Life")

	rand.Seed(time.Now().UnixNano())
	game := NewGame()

	slider = NewSlider(25, 25, adjustTPS)

	if err := ebiten.RunGame(game); err != nil {
		fmt.Println("error:", err)
	}
}
