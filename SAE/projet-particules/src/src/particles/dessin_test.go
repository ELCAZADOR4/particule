package particles

import (
	"testing"
)

func TestDessin(t *testing.T) {
	p := &Particle{PositionX: 0, PositionY: 0, ColorRed: 0, ColorGreen: 0, ColorBlue: 0}
	p.Dessin()
	// On vérifie que la position est dans les limites attendues
	if p.PositionX == 0 && p.PositionY == 0 {
		t.Fatal("La position de la particule n'est pas dans les limites attendues")
	}

	// On vérifie que la couleur est verte
	if p.ColorRed != 0 || p.ColorBlue != 0 || p.ColorGreen != 1 {
		t.Fatal("La couleur de la particule n'est pas vertes")
	}
}
