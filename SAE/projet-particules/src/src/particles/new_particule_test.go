package particles

import (
	"project-particles/config"
	"testing"
)

// Test pour une particule
// Test que la particule soit créer
func TestPresenceParticule(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = true
	config.General.SpawnRate = 0
	config.General.SpawnX = 400
	config.General.SpawnY = 300
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	system := NewSystem()

	if system.Content.Len() != 1 {
		t.Fatalf("il n'y a  pas une particules")
	}
}

// Test que la particule soit bien dans l'écran
func TestPositionsInWindowParticule(t *testing.T) {
	for i := -2000; i < 1000; i = i + 500 {
		config.General.InitNumParticles = 1
		config.General.RandomSpawn = true
		config.General.SpawnX = i
		config.General.SpawnY = i
		config.General.WindowSizeX = 800
		config.General.WindowSizeY = 600

		system := NewSystem()

		for e := system.Content.Front(); e != nil; e = e.Next() {
			p, _ := e.Value.(*Particle)
			if p.Rotation == 0 {
				if p.PositionX < 0+10 || p.PositionX > float64(config.General.WindowSizeX)-10 {
					t.Fatalf("PositionX hors limites: %f", p.PositionX)
				}
				if p.PositionY < 0+10 || p.PositionY > float64(config.General.WindowSizeY)-10 {
					t.Fatalf("PositionY hors limites: %f", p.PositionY)
				}
			}

		}
	}
}

// teste que les couleur on bien été rentré
func TestParticuleColors(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = true
	config.General.SpawnX = 400
	config.General.SpawnY = 300
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
