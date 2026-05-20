package particles

import (
	"testing"

	"project-particles/config"
)

// Teste pour plusieurs particules

// TestUpdateChangesPositions teste que Update() modifie les positions via la vitesse
func TestUpdateChangesPositionParticules(t *testing.T) {
	config.General.InitNumParticles = 10000
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

// TestUpdateCreateNewParticles teste que Update() crée de nouvelles particules avec SpawnRate
func TestUpdateCreateNewParticles(t *testing.T) {
	config.General.InitNumParticles = 1000000
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

	expectedCount := initialCount + int(config.General.SpawnRate) // fait le calcule du nombre de particules

	if finalCount != expectedCount {
		t.Fatalf(" Mauvais nombre de particules") // renvoie une erreur
	}
}

// test pour voir si les particules tourne
func TestUpdateRotateParticules(t *testing.T) {
	config.General.InitNumParticles = 1000000
	config.General.RandomSpawn = false
	config.General.SpawnX = 50
	config.General.SpawnY = 50
	config.General.SpawnRate = 100
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	system := NewSystem()
	p := system.Content.Front().Value.(*Particle)

	initialRot := p.Rotation
	system.Update()
	afterUpdate := p.Rotation

	if afterUpdate == initialRot {
		t.Fatalf(" Rotation ne bouge pas ")
	}
}

// TestMultipleParticlesScalePositiveFixed vérifie Scale pour RandomSpawn=false
func TestMultipleParticlesScalePositiveFixed(t *testing.T) {
	config.General.InitNumParticles = 5
	config.General.RandomSpawn = false
	config.General.SpawnX = 50
	config.General.SpawnY = 50
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	s := NewSystem()
	for e := s.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		if p.ScaleX <= 0 || p.ScaleY <= 0 {
			t.Fatalf("scale not positive (fixed): ScaleX=%f ScaleY=%f", p.ScaleX, p.ScaleY)
		}
	}
}

// TestInitialPositionsSameParticules vérifie que lorsque RandomSpawn=false
// toutes les particules initiales ont la même position (spawn fixe).
func TestInitialPositionsSameParticules(t *testing.T) {
	config.General.InitNumParticles = 100000
	config.General.RandomSpawn = false
	config.General.SpawnX = 200
	config.General.SpawnY = 180
	config.General.WindowSizeX = 800
	config.General.WindowSizeY = 600

	sys := NewSystem()

	first := sys.Content.Front().Value.(*Particle)
	fx := first.PositionX
	fy := first.PositionY

	for e := sys.Content.Front(); e != nil; e = e.Next() {
		p := e.Value.(*Particle)
		if p.PositionX != fx || p.PositionY != fy { // compare la position initial des deux particules
			t.Fatalf("Spawn différent")
		}
	}
}
