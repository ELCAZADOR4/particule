package particles

import (
	"testing"
)

func TestLampeTorche1(t *testing.T) {
	p := &Particle{
		PositionX: 100,
		PositionY: 100,
		vitesseX:  2,
		vitesseY:  -3,
		Opacity:   0.5,
		ScaleX:    1.0,
		ScaleY:    1.0,
	}

	p.Lampetorche()

	// Verifie me déplacement
	if p.PositionX != 102 || p.PositionY != 97 {
		t.Errorf("Le déplacement est incorrect. ")
	}

	// Vérifie la baisse de l'opacité (0.5 - 0.02 = 0.48)
	if p.Opacity != 0.48 {
		t.Errorf("L'opacité aurait dû baisser à 0.48, obtenu: %f", p.Opacity)
	}

	// Vérifie la baisse de la taille
	if p.ScaleX >= 1.0 || p.ScaleY >= 1.0 {
		t.Error("La taille (Scale) n'a pas diminué")
	}
}

func TestLampeTorche2(t *testing.T) {
	p := &Particle{
		Opacity:   1.0,
		PositionX: 500, // Position éloignée
		PositionY: 500,
	}

	p.Lampetorche()

	// Vérifie que la particule a rejoint la position de la souris (0,0 dans le test)
	if p.PositionX > 50 || p.PositionY > 50 {
		t.Errorf("La particule ne semble pas avoir rejoint la souris (0,0)")
	}
}
