package particles

import (
	"project-particles/config"
)

func (p *Particle) ChangecolorFR() {

	windowW := float64(config.General.WindowSizeX)
	// Change la couleur en fonction de la position de la particule
	if p.PositionX <= windowW/3 {
		p.ColorBlue = 100
		p.ColorGreen = 0
		p.ColorRed = 0
	} else if p.PositionX > windowW/3 && p.PositionX <= 2*windowW/3 {
		p.ColorBlue = 100
		p.ColorGreen = 100
		p.ColorRed = 100
	} else {
		p.ColorBlue = 0
		p.ColorGreen = 0
		p.ColorRed = 1
	}

}
