package particles

import (
	"project-particles/config"
	"testing"
)

// Test pour plusieurs particules
// Test que les particules soit créer
func TestPresencesParticules(t *testing.T) {
	config.General.InitNumParticles = 500000
	config.General.RandomSpawn = true
	config.General.SpawnRate = 5000
	config.General.SpawnX = 400
	config.General.SpawnY = 300
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	system := NewSystem()

	if system.Content.Len() != config.General.InitNumParticles {
		t.Fatalf("il y a toutes les particules") // renvoie si il y a une erreur
	}
}

// Test que les particules soit bien dans l'écran
func TestPositionInWindowParticules(t *testing.T) {
	for i := -2000; i < 1000; i = i + 500 {
		config.General.InitNumParticles = 1000000
		config.General.RandomSpawn = true
		config.General.SpawnX = i
		config.General.SpawnY = i
		config.General.WindowSizeX = 800
		config.General.WindowSizeY = 600

		system := NewSystem()
		for e := system.Content.Front(); e != nil; e = e.Next() {
			p, _ := e.Value.(*Particle)
			// Vérifier que toutes les positions sont dans les limites de la fenêtre
			if p.PositionX < 0 || p.PositionX > float64(config.General.WindowSizeX) {
				t.Fatalf("PositionX hors limites: %f (fenêtre: 0-%d)", p.PositionX, config.General.WindowSizeX)
			}
			if p.PositionY < 0 || p.PositionY > float64(config.General.WindowSizeY) {
				t.Fatalf("PositionY hors limites: %f (fenêtre: 0-%d)", p.PositionY, config.General.WindowSizeY)
			}
		}
	}
}

// teste que les couleur on bien été rentré
func TestParticulesColor(t *testing.T) {

	config.General.InitNumParticles = 100000
	config.General.RandomSpawn = true
	config.General.SpawnRate = 10000
	config.General.SpawnX = 400
	config.General.SpawnY = 500
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	system := NewSystem()

	for e := system.Content.Front(); e != nil; e = e.Next() {
		p, _ := e.Value.(*Particle)

		// Vérifier que les couleurs sont dans la plage [0, 1]
		if p.ColorRed < 0 || p.ColorRed > 1 {
			t.Fatalf("la couleur rouge n'est pas [0,1]: %f", p.ColorRed)
		}
		if p.ColorGreen < 0 || p.ColorGreen > 1 {
			t.Fatalf("la couleur verte n'est pas [0,1]: %f", p.ColorGreen)
		}
		if p.ColorBlue < 0 || p.ColorBlue > 1 {
			t.Fatalf("la couleur bleue n'est pas [0,1]: %f", p.ColorBlue)
		}
	}
}
