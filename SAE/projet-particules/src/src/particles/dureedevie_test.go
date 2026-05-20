package particles

import (
	"testing"
)

func TestDurreedevie(t *testing.T) {
	p := &Particle{ScaleX: 1.0, ScaleY: 1.0, Age: 0, Dead: false}

	// Simuler plusieurs appels à Durreedevie
	for i := 0; i < 250; i++ {
		p.Durreedevie()
	}

	// Vérifie que la taille a diminué correctement
	if p.ScaleX >= 1.0 || p.ScaleY >= 1.0 {
		t.Fatal("La taille de la particule n'a pas diminué correctement")
	}

	// Vérifie que l'âge est correct
	if p.Age != 250 {
		t.Fatalf("L'âge de la particule est incorrect, attendu 250, obtenu %d", p.Age)
	}

	// Vérifie que la particule est marquée comme morte après un certain âge
	if !p.Dead {
		t.Fatal("La particule n'est pas marquée comme morte après avoir dépassé la durée de vie")
	}
}
