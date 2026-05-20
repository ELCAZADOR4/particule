package particles

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func (p *Particle) Para() {
	// Attire les particules vers la position de la souris
	sourisX, sourisY := ebiten.CursorPosition()

	// Calcul de la distance entre la particule et la souris
	deltaX := float64(sourisX) - p.PositionX
	deltaY := float64(sourisY) - p.PositionY

	distance := math.Sqrt(deltaX*deltaX + deltaY*deltaY)

	seuil := 100.0
	// Si la particule est trop proche de la souris, elle devient noire pour ne pas la voir
	if distance < seuil {
		p.ColorBlue = 0
		p.ColorRed = 0
		p.ColorGreen = 0
		return
	}

	if distance >= seuil && distance <= seuil+5 {
		p.vitesseX = -p.vitesseX * 0.01
		p.vitesseY = -p.vitesseY * 0.01
	}

}
