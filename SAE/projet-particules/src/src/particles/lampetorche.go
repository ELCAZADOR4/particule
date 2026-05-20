package particles

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

func (p *Particle) Lampetorche() {
	// Position de la souris
	sourisX, sourisY := ebiten.CursorPosition()

	// Naissance de la particule au niveau de la souris
	if p.Opacity >= 1 {
		p.PositionX = float64(sourisX)
		p.PositionY = float64(sourisY)

		// Vitesse de flamme
		p.vitesseX = (rand.Float64() - 0.5) * 1.5
		p.vitesseY = -rand.Float64()*2 - 1

		// Taille initiale
		p.ScaleX = rand.Float64()*0.5 + 0.8
		p.ScaleY = rand.Float64()*0.5 + 0.8
	}

	// Déplacement
	p.PositionX += p.vitesseX
	p.PositionY += p.vitesseY

	// Réduction progressive de la taille
	p.ScaleX *= 0.99
	p.ScaleY *= 0.99

	// Perte d'opacité (durée de vie)
	p.Opacity -= 0.02
	if p.Opacity < 0 {
		p.Opacity = 0
	}
	p.ColorRed = 0.5
	p.ColorGreen = 0.8 * p.Opacity
}
