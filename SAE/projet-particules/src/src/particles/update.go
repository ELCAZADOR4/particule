package particles

import (
	"math"
	"math/rand/v2"
	"project-particles/config"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.

func (s *System) Update() {

	// Dimensions d’une particule (utilisées pour ne pas qu’elle dépasse de l'écran
	const particleW float64 = 20.0
	const particleH float64 = 20.0

	// Dimensions de la fenêtre
	windowW := float64(config.General.WindowSizeX)
	windowH := float64(config.General.WindowSizeY)

	// Calcul du rayon de la particule pour empêcher qu’elle dépasse des bords lors du spawn
	diag := math.Sqrt(particleW*particleW + particleH*particleH)
	moitier := diag / 2

	// Limites de déplacement autorisées
	minX := moitier
	maxX := windowW - moitier
	minY := moitier
	maxY := windowH - moitier

	for e := s.Content.Front(); e != nil; e = e.Next() {

		// Récupération de la particule
		p := e.Value.(*Particle)
		next := e.Next()
		// mise a jour de la vitesse et de la rotation
		p.PositionX += p.vitesseX
		p.PositionY += p.vitesseY
		p.Rotation += 0.05

		// Si la particule sort de l'écran
		if p.PositionX < minX || p.PositionX > maxX ||
			p.PositionY < minY || p.PositionY > maxY {
			p.Opacity = 1
			// Respawn aléatoire si RandomSpawn = true
			if config.General.RandomSpawn {
				p.PositionX = minX + rand.Float64()*(maxX-minX)
				p.PositionY = minY + rand.Float64()*(maxY-minY)

			} else {

				// Respawn fixe au point d'apparition si RandomSpawn = false
				spawnX := float64(config.General.SpawnX)
				spawnY := float64(config.General.SpawnY)
				p.PositionX = spawnX
				p.PositionY = spawnY
			}
		}
		p.particule()

		if p.Dead {
			s.Content.Remove(e)
			e = next
			continue
		}
	}

	for i := 0; i < int(config.General.SpawnRate); i++ {

		var x, y float64

		// Position de spawn aléatoire
		if config.General.RandomSpawn {
			x = minX + rand.Float64()*(maxX-minX)
			y = minY + rand.Float64()*(maxY-minY)
		} else { // Position de spawn fixe
			x = float64(config.General.SpawnX)
			y = float64(config.General.SpawnY)
		}

		if config.General.Durreedevie {
			nouvelleParticule := &Particle{
				PositionX: x,
				PositionY: y,

				// Vitesse aléatoire entre -1 et 1
				vitesseX: rand.Float64()*2 - 1,
				vitesseY: rand.Float64()*2 - 1,

				Opacity:  1.0,
				ScaleX:   1,
				ScaleY:   1,
				ColorRed: 1,
				// Rotation aléatoire entre 0 et 2π
				Rotation: rand.Float64() * 2 * math.Pi,
			}
			s.Content.PushBack(nouvelleParticule)

		} else {

			nouvelleParticule := &Particle{
				PositionX: x,
				PositionY: y,

				// Vitesse aléatoire entre -1 et 1
				vitesseX: rand.Float64()*2 - 1,
				vitesseY: rand.Float64()*2 - 1,

				Opacity: 1.0,
				ScaleX:  1,
				ScaleY:  1,
				// Rotation aléatoire entre 0 et 2π
				Rotation: rand.Float64() * 2 * math.Pi,
			}

			s.Content.PushBack(nouvelleParticule)
		}
	}
}
