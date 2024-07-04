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
	slider   *Slider
)

func main() {
	// Demander à l'utilisateur d'entrer les dimensions de la grille
	fmt.Print("Entrez la taille de la grille (carré) : ")
	fmt.Scan(&width)

	height = width

	// Vérifier les entrées de l'utilisateur
	if width < 15 || width > 60 {
		fmt.Println("Erreur de taille : la taille doit être comprise entre 15 et 60.")
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

func adjustTPS(value float64) {
	newTPS := 10 + int(value*30) // Ajuste le TPS entre 10 et 40
	ebiten.SetTPS(newTPS)
	fmt.Printf("Vitesse ajustée à: %d TPS\n", newTPS)
}

func getIndex(x, y, width int) int {
	return y*width + x
}
