package particles

import (
	"math"
	"project-particles/config"
)

func (p *Particle) Dessin() {
	centerX := float64(config.General.WindowSizeX) / 2
	centerY := float64(config.General.WindowSizeY) / 2
	rayon := 200.0

	// fait tourner les particules dans le cercle
	angleR := float64(p.Age) * 0.02

	angleI := p.Rotation // On utilise la rotation aléatoire du spawn

	total := angleI + angleR

	p.PositionX = centerX + rayon*math.Cos(total)
	p.PositionY = centerY + rayon*math.Sin(total)

	p.ColorGreen = 1

}
