package particles

import (
	"testing"
)

func TestPara1(t *testing.T) {
	p := &Particle{
		PositionX:  10,
		PositionY:  10,
		ColorRed:   1,
		ColorGreen: 1,
		ColorBlue:  1,
	}

	p.Para()
	// Vérification que la couleur est noire
	if p.ColorRed != 0 || p.ColorGreen != 0 || p.ColorBlue != 0 {
		t.Errorf("La particule devrait être noire (0,0,0) car elle est sous le seuil")
	}
}

func TestPara2(t *testing.T) {
	// On place la particule exactement à la limite du seuil (entre 100 et 105)
	p := &Particle{
		PositionX: 102,
		PositionY: 0,
		vitesseX:  10.0,
		vitesseY:  10.0,
		ColorRed:  1,
	}

	p.Para()

	// Vérification de la vitesse
	attendu := -0.1
	if p.vitesseX != attendu || p.vitesseY != attendu {
		t.Errorf("La vitesse n'a pas été réduite correctement. Attendu: %f, Obtenu: X:%f, Y:%f",
			attendu, p.vitesseX, p.vitesseY)
	}

	// Vérification de la couleur
	if p.ColorRed == 0 {
		t.Error("La particule ne devrait pas être noire car elle est au-dessus du seuil de 100")
	}
}
