package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	width    = 100 // Taille de la grille en largeur
	height   = 100 // Taille de la grille en hauteur
	cellSize = 10  // Taille initiale de chaque cellule pour l'affichage
	offsetX  = 0.0
	offsetY  = 0.0
	slider  *Slider
)

type Game struct {
	Grid             []bool
	IsGameOver       bool
	LivingCellsCount int
	GenerationCount  int
}

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

	if height < 15 || height > 60 &&  width < 15 || width > 60 {
		fmt.Println("Erreur de taille : la taille doit être inférieur à 60 ou supérieur à .")
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

func NewGame() *Game {
	gridSize := width * height
	game := &Game{
		Grid:             make([]bool, gridSize),
		IsGameOver:       false,
		LivingCellsCount: 0,
		GenerationCount:  0,
	}
	return game
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)

	// Calculer la taille de la grille dans l'écran
	gridWidth := width * cellSize
	gridHeight := height * cellSize

	// Calculer le centre de l'écran
	centerX := float64(screen.Bounds().Dx()) / 2
	centerY := float64(screen.Bounds().Dy()-60) / 2 // Soustraire la hauteur de la barre des statistiques

	// Calculer le coin supérieur gauche de la grille dans l'écran
	gridStartX := centerX - float64(gridWidth)/2
	gridStartY := centerY - float64(gridHeight)/2

	// Dessiner la grille statique en arrière-plan
	for y := 0; y <= height; y++ {
		yPos := gridStartY + float64(y*cellSize)
		ebitenutil.DrawLine(screen, gridStartX, yPos, gridStartX+float64(gridWidth), yPos, color.Gray16{0x9999})
	}

	for x := 0; x <= width; x++ {
		xPos := gridStartX + float64(x*cellSize)
		ebitenutil.DrawLine(screen, xPos, gridStartY, xPos, gridStartY+float64(gridHeight), color.Gray16{0x9999})
	}

	// Dessiner une barre noire en bas de l'écran pour les statistiques
	statsBarHeight := 60
	statsBarY := screen.Bounds().Dy() - statsBarHeight
	ebitenutil.DrawRect(screen, 0, float64(statsBarY), float64(screen.Bounds().Dx()), float64(statsBarHeight), color.Black)

	// Dessiner les statistiques sur la barre noire
	stats := fmt.Sprintf("Génération: %d\nCellules vivantes: %d", g.GenerationCount, g.LivingCellsCount)
	ebitenutil.DebugPrintAt(screen, stats, 10, statsBarY+20)

	// Dessiner le slider
	slider.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// La taille de la fenêtre de jeu est la taille de la grille plus la barre de statistiques
	return width * cellSize, height * cellSize + 60
}

func (g *Game) Update() error {
	slider.Update()

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		offsetY += 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		offsetY -= 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		offsetX += 10
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		offsetX -= 10
	}

	return nil
}

func getIndex(x, y, width int) int {
	return y*width + x
}

func adjustTPS(value float64) {
	newTPS := 10 + int(value*30) // Ajuste le TPS entre 10 et 40
	ebiten.SetTPS(newTPS)
}

func (s *Slider) Draw(screen *ebiten.Image) {
	// Dessiner la barre du slider
	barColor := color.RGBA{0x80, 0x80, 0x80, 0xff}
	barRect := ebiten.NewImage(s.width, s.height)
	barRect.Fill(barColor)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.x), float64(s.y))
	screen.DrawImage(barRect, op)

	// Dessiner le curseur du slider
	cursorWidth := 10
	cursorHeight := s.height + 10
	cursorX := s.x + int(s.value*float64(s.width)) - cursorWidth/2
	cursorY := s.y - (cursorHeight-s.height)/2
	cursorColor := color.RGBA{0xff, 0x00, 0x00, 0xff}
	cursorRect := ebiten.NewImage(cursorWidth, cursorHeight)
	cursorRect.Fill(cursorColor)
	op = &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(cursorX), float64(cursorY))
	screen.DrawImage(cursorRect, op)
}

func (s *Slider) Update() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		mx, my := ebiten.CursorPosition()
		if mx >= s.x && mx < s.x+s.width && my >= s.y && my < s.y+s.height {
			s.value = math.Max(0, math.Min(1, float64(mx-s.x)/float64(s.width)))
			s.onChange(s.value)
		}
	}
}

// Slider struct and related methods
type Slider struct {
	x, y, width, height int
	value               float64
	onChange            func(float64)
}

func NewSlider(x int, y int, onChange func(float64)) *Slider {
	return &Slider{
		x:        x,
		y:        y,
		width:    250,
		height:   10,
		value:    0.5, // Valeur initiale du slider au milieu
		onChange: onChange,
	}
}
