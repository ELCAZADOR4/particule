package particles

import (
	"project-particles/config"
)

func (p *Particle) Changecolor() {

	windowW := float64(config.General.WindowSizeX)

	p.ColorRed = float64((p.PositionX / windowW) * 1)
	p.ColorBlue = 0
	p.ColorGreen = float64(1 - (p.PositionX / windowW))
}
