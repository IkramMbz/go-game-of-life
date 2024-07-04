package main

import (
	"fmt"
	"image/color"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)


type Game struct {
	Grid             []bool
	IsGameOver       bool
	LivingCellsCount int
	GenerationCount  int
}

func NewGame() *Game {
	game := &Game{
		Grid:             make([]bool, width*height),
		IsGameOver:       false,
		LivingCellsCount: 0,
		GenerationCount:  0,
	}
	game.InitializeGrid()
	return game
}

func (g *Game) InitializeGrid() {
	for i := 0; i < width*height; i++ {
		g.Grid[i] = rand.Intn(2) == 1
		if g.Grid[i] {
			g.LivingCellsCount++
		}
	}
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
		ebitenutil.DrawLine(screen, gridStartX, yPos, gridStartX+float64(width*cellSize), yPos, color.Gray16{0x9999})
	}

	for x := 0; x <= width; x++ {
		xPos := gridStartX + float64(x*cellSize)
		ebitenutil.DrawLine(screen, xPos, gridStartY, xPos, gridStartY+float64(height*cellSize), color.Gray16{0x9999})
	}

	// Dessiner les cellules vivantes
	for i := 0; i < width*height; i++ {
		if g.Grid[i] {
			x := i % width
			y := i / width
			ebitenutil.DrawRect(screen, gridStartX+float64(x*cellSize), gridStartY+float64(y*cellSize), float64(cellSize), float64(cellSize), color.Black)
		}
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// La taille de la fenêtre de jeu est la taille de la grille
	return width * cellSize, height * cellSize + 60
	
	//return screenWidth, screenHeight
}

func (g *Game) Update() error {
	slider.Update()

	g.Grid = nextGeneration(g)

	// Mettre à jour les statistiques
	g.GenerationCount++
	g.LivingCellsCount = countLivingCells(g.Grid)

	return nil
}

func countLivingCells(grid []bool) int {
	count := 0
	for _, alive := range grid {
		if alive {
			count++
		}
	}
	return count
}

func nextGeneration(g *Game) []bool {
	next := make([]bool, width*height)

	for i := 0; i < width*height; i++ {
		x := i % width
		y := i / width
		count := countNeighbors(g.Grid, x, y)
		if g.Grid[i] && (count == 2 || count == 3) {
			next[i] = true
		} else if !g.Grid[i] && count == 3 {
			next[i] = true
		}
	}

	return next
}

func countNeighbors(grid []bool, x, y int) int {
	count := 0
	for yOffset := -1; yOffset <= 1; yOffset++ {
		for xOffset := -1; xOffset <= 1; xOffset++ {
			if !(yOffset == 0 && xOffset == 0) {
				neighborX := (x + xOffset + width) % width
				neighborY := (y + yOffset + height) % height
				if grid[neighborY*width+neighborX] {
					count++
				}
			}
		}
	}
	return count
}
