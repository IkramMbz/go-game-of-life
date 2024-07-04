package main

type Game struct {
	Grid             []bool
	IsGameOver       bool
	LivingCellsCount int
	GenerationCount  int
}



func (g *Game) InitializeGrid() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			g.Grid[y][x] = rand.Intn(2) == 1
			if g.Grid[y][x] {
				g.LivingCellsCount++
			}
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
	centerY := float64(screen.Bounds().Dy()) / 2

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

	// Dessiner les cellules vivantes
	for y, row := range g.Grid {
		for x, alive := range row {
			if alive {
				ebitenutil.DrawRect(screen, gridStartX+float64(x*cellSize), gridStartY+float64(y*cellSize), float64(cellSize), float64(cellSize), color.Black)
			}
		}
	}

	// Dessiner les statistiques
	stats := fmt.Sprintf("Génération: %d   Cellules vivantes: %d", g.GenerationCount, g.LivingCellsCount)
	ebitenutil.DebugPrint(screen, stats)

	// Dessiner le slider
	slider.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// La taille de la fenêtre de jeu est la taille de la grille
	return width * cellSize, height * cellSize
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

	g.Grid = nextGeneration(g)

	// Mettre à jour les statistiques
	g.GenerationCount++
	g.LivingCellsCount = countLivingCells(g.Grid)

	return nil
}

func countLivingCells(grid [][]bool) int {
	count := 0
	for _, row := range grid {
		for _, alive := range row {
			if alive {
				count++
			}
		}
	}
	return count
}

func (g *Game) ResetGame() {
	g.IsGameOver = false
	g.GenerationCount = 0
	g.InitializeGrid()
}

func nextGeneration(g *Game) [][]bool {
	next := make([][]bool, height)
	for i := range next {
		next[i] = make([]bool, width)
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			count := countNeighbors(g.Grid, x, y)
			if g.Grid[y][x] && (count == 2 || count == 3) {
				next[y][x] = true
			} else if !g.Grid[y][x] && count == 3 {
				next[y][x] = true
			}
		}
	}

	return next
}

func countNeighbors(grid [][]bool, x, y int) int {
	count := 0
	for yOffset := -1; yOffset <= 1; yOffset++ {
		for xOffset := -1; xOffset <= 1; xOffset++ {
			if !(yOffset == 0 && xOffset == 0) {
				neighborX := (x + xOffset + width) % width
				neighborY := (y + yOffset + height) % height
				if grid[neighborY][neighborX] {
					count++
				}
			}
		}
	}
	return count
}

