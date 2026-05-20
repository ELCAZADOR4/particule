package particles

import (
	"math"
	"project-particles/config"
)

func (p *Particle) Gravite() {

	const particleH float64 = 20.0
	const particleW float64 = 20.0
	windowH := float64(config.General.WindowSizeY)
	windowW := float64(config.General.WindowSizeX)

	diag := math.Sqrt(particleW*particleW + particleH*particleH)
	moitier := diag / 2

	minX := moitier
	maxX := windowW - moitier
	maxY := windowH - moitier

	p.vitesseY += 0.1
	p.Rotation = 0

	if p.PositionY+p.vitesseY > maxY {
		p.PositionY = maxY             // Recalage précis sur le sol
		p.vitesseY = -p.vitesseY * 0.4 // Rebond (inversion + amortissement)
		p.vitesseX *= 0.8
		p.ColorBlue = 1
		p.ColorGreen = 0
		p.ColorRed = 0
		// Arrêt du rebond si la vitesse est trop faible (évite les vibrations)
		if math.Abs(p.vitesseY) < 0.2 {
			p.vitesseY = 0
		}
	}

	if p.PositionX+p.vitesseX > maxX {
		p.PositionX = maxX
		p.vitesseX = -p.vitesseX * 0.6 // Rebond latéral amorti
	} else if p.PositionX+p.vitesseX < minX {
		p.PositionX = minX
		p.vitesseX = -p.vitesseX * 0.6 // Rebond latéral amorti
	}
	p.ScaleY = 3
	p.ScaleX = 0.2
	const g float64 = 9.81
	const frame float64 = 60
	var calcule float64

	calcule = g / frame

	config.General.Age++

	if config.General.Age > 0 {
		p.vitesseY += calcule
	}

}
