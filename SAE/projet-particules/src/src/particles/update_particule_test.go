package particles

import (
	"testing"

	"project-particles/config"
)

// Test pour une particule au départ
// TestUpdateChangesPositions teste que Update() modifie les positions via la vitesse
func TestUpdateChangesPositionsParticule(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = false
	config.General.SpawnX = 100
	config.General.SpawnY = 100
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	system := NewSystem()

	// Récupérer la première particule
	p := system.Content.Front().Value.(*Particle)

	// Sauvegarder l'état initial
	initialX := p.PositionX
	initialY := p.PositionY

	// Appeler Update()
	system.Update()

	// Vérifier si la position a changé
	if p.PositionX == initialX || p.PositionY == initialY {
		t.Fatalf("la position n'a pas bougé")
	}
}

// TestUpdateCreatesNewParticles teste que Update() crée de nouvelles particules avec SpawnRate
func TestUpdateCreatesNewParticle(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = true
	config.General.SpawnRate = 100000
	config.General.SpawnX = 400
	config.General.SpawnY = 300
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	system := NewSystem()

	initialCount := system.Content.Len()
	// Appeler Update()
	system.Update()

	finalCount := system.Content.Len()

	expectedCount := initialCount + int(config.General.SpawnRate) // ajoute manuellement le nombre de particule normalement ajouté
	if finalCount != expectedCount {
		t.Fatalf(" Mauvais nombre de particules")
	}
}

func TestUpdateRotatesParticule(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = false
	config.General.SpawnX = 50
	config.General.SpawnY = 50
	config.General.SpawnRate = 10000
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	system := NewSystem()
	p := system.Content.Front().Value.(*Particle)

	initialRot := p.Rotation // sens initial
	system.Update()
	afterUpdate := p.Rotation // sens après update

	if afterUpdate == initialRot {
		t.Fatalf(" Rotation ne bouge pas ")
	}
}

// TestSingleParticleScalePositive vérifie que la scale d'une particule est > 0
func TestSingleParticleScalePositive(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = false
	config.General.SpawnX = 100
	config.General.SpawnY = 100
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	s := NewSystem()
	if s.Content.Len() != 1 {
		t.Fatalf("expected 1 particle, got %d", s.Content.Len())
	}

	p := s.Content.Front().Value.(*Particle)
	if p.ScaleX <= 0 || p.ScaleY <= 0 {
		t.Fatalf("scale not positive: ScaleX=%f ScaleY=%f", p.ScaleX, p.ScaleY)
	}
}

// TestInitialPositionSameParticule vérifie que quand RandomSpawn=false et
// plusieurs particules doivent apparaître au même spawn, leur position
// initiale est identique.
func TestInitialPositionSameParticule(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = false
	config.General.SpawnX = 120
	config.General.SpawnY = 150
	config.General.SpawnRate = 1000
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	s := NewSystem()

	// prendre la première particule comme référence
	first := s.Content.Front().Value.(*Particle)
	refX := first.PositionX
	refY := first.PositionY
	// compare la position des autres avec la premières
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		if p.PositionX != refX || p.PositionY != refY {
			t.Fatalf("Spawn différent")
		}
	}
}
