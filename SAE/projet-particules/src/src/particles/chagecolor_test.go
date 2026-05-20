package particles

import (
	"project-particles/config"
	"testing"
)

func TestChangecolor(t *testing.T) {
	config.General.WindowSizeX = 900
	windowW := float64(config.General.WindowSizeX)

	p := &Particle{PositionX: windowW, PositionY: 0}
	p.Changecolor()

	// verifie que à droite ColorRed est 1 et ColorGreen est 0
	if p.ColorRed != 1 || p.ColorGreen != 0 {
		t.Errorf("Erreur à droite: attendu R:1 G:0")
	}

	p.PositionX = 0
	p.Changecolor()

	// Verifie que à gauche ColorRed est 0 et ColorGreen est 1
	if p.ColorRed != 0 || p.ColorGreen != 1 {
		t.Errorf("Erreur à gauche: attendu R:0 G:1,")
	}

}
