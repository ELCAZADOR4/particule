package particles

import (
	"container/list"
	"math"
	"math/rand/v2"
	"project-particles/config"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {

	// Liste chaînée qui contiendra toutes les particules
	l := list.New()

	// Dimensions d’une particule (utilisées pour éviter qu’elle spawn hors écran)
	const particleW float64 = 20.0
	const particleH float64 = 20.0

	// Dimensions de la fenêtre
	windowW := float64(config.General.WindowSizeX)
	windowH := float64(config.General.WindowSizeY)

	// Calcul du rayon de la particule (diagonale / 2) pour empêcher qu’elle dépasse des bords lors du spawn
	diag := math.Sqrt(particleW*particleW + particleH*particleH)
	radius := diag / 2

	// Limites de spawn autorisées
	minX := radius
	maxX := windowW - radius
	minY := radius
	maxY := windowH - radius

	// Boucle de création des particules
	for i := 0; i < config.General.InitNumParticles; i++ {

		var positionX, positionY float64

		// Spawn aléatoire si activé dans la config, si RandomSpawn = True c'est en aléatoire
		if config.General.RandomSpawn {

			// Position aléatoire dans les limites autorisées
			positionX = minX + rand.Float64()*(maxX-minX)
			positionY = minY + rand.Float64()*(maxY-minY)

		} else {
			// Spawn fixe défini dans la configuration
			positionX = float64(config.General.SpawnX)
			positionY = float64(config.General.SpawnY)

		}

		// Ajout d’une nouvelle particule dans la liste
		l.PushBack(&Particle{
			PositionX: positionX,
			PositionY: positionY,
			ScaleX:    1.0,
			ScaleY:    1.0,

			// Couleur aléatoire
			ColorRed:   rand.Float64(),
			ColorGreen: rand.Float64(),
			ColorBlue:  rand.Float64(),

			Opacity:  1,
			Rotation: rand.Float64() * 2 * math.Pi, // rotation aléatoire

			// Vitesse aléatoire entre -1 et 1
			vitesseX: rand.Float64()*2 - 1,
			vitesseY: rand.Float64()*2 - 1,
		})
	}

	return System{Content: l}
}
