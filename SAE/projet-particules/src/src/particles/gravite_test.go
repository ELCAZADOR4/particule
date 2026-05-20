package particles

import (
	"project-particles/config"
	"testing"
)

func TestGravite(t *testing.T) {
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600
	config.General.Age = 0

	t.Run("Chute_Libre", func(t *testing.T) {
		p := &Particle{
			PositionX: 100,
			PositionY: 100,
			vitesseX:  0,
			vitesseY:  0,
		}

		p.Gravite()

		// Vérifie l'augmentation de la vitesse
		if p.vitesseY <= 0 {
			t.Errorf("La vitesseY devrait être positive en chute libre %f", p.vitesseY)
		}

		// Vérifie que l'Age a bien été incrémenté
		if config.General.Age != 1 {
			t.Errorf("L'Age de la config devrait être 1, got %d", config.General.Age)
		}

	})

	t.Run("Rebond_Au_Sol", func(t *testing.T) {

		p := &Particle{
			PositionX: 100,
			PositionY: 585, // Juste au dessus du maxY
			vitesseY:  10,  // Vitesse vers le bas
		}

		p.Gravite()

		// Vérifie que la vitesse s'est inversée
		if p.vitesseY >= 0 {
			t.Errorf("La vitesseY après rebond doit être négative (remontée), got %f", p.vitesseY)
		}

	})

}
