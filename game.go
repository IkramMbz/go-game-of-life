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
	Zoom             float64
	OffsetX          float64
	OffsetY          float64
}

func NewGame() *Game {
	game := &Game{
		Grid:             make([]bool, width*height),
		IsGameOver:       false,
		LivingCellsCount: 0,
		GenerationCount:  0,
		Zoom:             1.0,
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

	// Calculer le nombre de cellules visibles
	cellsInWidth := float64(screen.Bounds().Dx()) / (float64(cellSize) * g.Zoom)
	cellsInHeight := float64(screen.Bounds().Dy()-60) / (float64(cellSize) * g.Zoom) // Soustraire la hauteur de la barre des statistiques

	// Calculer le centre de l'écran
	centerX := float64(screen.Bounds().Dx()) / 2
	centerY := float64(screen.Bounds().Dy()-60) / 2 // Soustraire la hauteur de la barre des statistiques

	// Calculer le coin supérieur gauche de la grille dans l'écran
	gridStartX := centerX - (cellsInWidth * float64(cellSize) * g.Zoom) / 2 + g.OffsetX
	gridStartY := centerY - (cellsInHeight * float64(cellSize) * g.Zoom) / 2 + g.OffsetY

	// Dessiner la grille statique en arrière-plan
	for y := 0; y <= int(cellsInHeight); y++ {
		yPos := gridStartY + float64(y*cellSize)*g.Zoom
		ebitenutil.DrawLine(screen, gridStartX, yPos, gridStartX+float64(cellsInWidth*float64(cellSize))*g.Zoom, yPos, color.Gray16{0x9999})
	}

	for x := 0; x <= int(cellsInWidth); x++ {
		xPos := gridStartX + float64(x*cellSize)*g.Zoom
		ebitenutil.DrawLine(screen, xPos, gridStartY, xPos, gridStartY+float64(cellsInHeight*float64(cellSize))*g.Zoom, color.Gray16{0x9999})
	}

	// Dessiner les cellules vivantes
	for i := 0; i < width*height; i++ {
		if g.Grid[i] {
			x := i % width
			y := i / width
			cellX := gridStartX + float64(x*cellSize)*g.Zoom
			cellY := gridStartY + float64(y*cellSize)*g.Zoom
			if cellX >= 0 && cellX < float64(screen.Bounds().Dx()) && cellY >= 0 && cellY < float64(screen.Bounds().Dy()) {
				ebitenutil.DrawRect(screen, cellX, cellY, float64(cellSize)*g.Zoom, float64(cellSize)*g.Zoom, color.Black)
			}
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
	// La taille de la fenêtre de jeu est fixe
	return outsideWidth, outsideHeight
}

func (g *Game) Update() error {
	slider.Update()

	// Gestion du zoom avec la molette de la souris
	_, dy := ebiten.Wheel()
	if dy != 0 {
		newZoom := g.Zoom + dy*0.1
		if newZoom >= 0.5 && newZoom <= 2.0 {
			g.Zoom = newZoom
		}
	}

	g.handleDrawing()

	g.Grid = nextGeneration(g)

	// Mettre à jour les statistiques
	g.GenerationCount++
	g.LivingCellsCount = countLivingCells(g.Grid)

	return nil
}

func (g *Game) handleDrawing() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// On ajuste les coordonnées de la souris en fonction du zoom et de sa direction
		x = int((float64(x) - g.OffsetX) / g.Zoom)
		y = int((float64(y) - g.OffsetY) / g.Zoom)

		// Calculer les coordonnées de la grille
		gridX := x / cellSize
		gridY := y / cellSize

		// Vérifier que les coordonnées sont dans les limites de la grille
		if gridX >= 0 && gridX < width && gridY >= 0 && gridY < height {
			//Cela à un impact direct sur ses voisins
			index := gridY*width + gridX
			if !g.Grid[index] {
				g.Grid[index] = true
				g.LivingCellsCount++
			}
		}
	}
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