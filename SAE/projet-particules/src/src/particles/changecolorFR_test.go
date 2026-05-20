package particles

import (
	"testing"

	"project-particles/config"
)

// Test pour la couleur bleue
func TestChangecolorBLUEFR(t *testing.T) {
	config.General.WindowSizeX = 900

	p := &Particle{PositionX: 10, ColorRed: 0, ColorGreen: 0, ColorBlue: 0}

	p.ChangecolorFR()

	if p.PositionX <= float64(config.General.WindowSizeX)/3 {
		if p.ColorRed != 0 || p.ColorGreen != 0 || p.ColorBlue != 100 {
			t.Fatal("la couleur ne correspond pas à la couleur")
		}
	}

}

// Test pour la couleur blanche
func TestChangecolorWHITEFR(t *testing.T) {
	config.General.WindowSizeX = 900

	p := &Particle{PositionX: 500, ColorRed: 0, ColorGreen: 0, ColorBlue: 0}

	p.ChangecolorFR()

	if p.PositionX <= float64(config.General.WindowSizeX)/3 {
		if p.ColorRed != 100 || p.ColorGreen != 100 || p.ColorBlue != 100 {
			t.Fatal("la couleur ne correspond pas à la couleur")
		}
	}

}

// Test pour la couleur rouge
func TestChangecolorREDFR(t *testing.T) {

	config.General.WindowSizeX = 900

	p := &Particle{PositionX: 800, ColorRed: 0, ColorGreen: 0, ColorBlue: 0}

	p.ChangecolorFR()

	if p.PositionX <= float64(config.General.WindowSizeX)/3 {
		if p.ColorRed != 100 || p.ColorGreen != 0 || p.ColorBlue != 0 {
			t.Fatal("la couleur ne correspond pas à la couleur")
		}
	}

}
