package main

import (
	"image/color"
	"math"
	"github.com/hajimehoshi/ebiten/v2"
)

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
		height:   25,
		value:    0.5, // Valeur initiale du slider au milieu
		onChange: onChange,
	}
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

// Ajuster les Tick par seconds pour faire varier la vitesse
func adjustTPS(value float64) {
	newTPS := 10 + int(value*70)
	ebiten.SetTPS(newTPS)
}
