package particles

import "project-particles/config"

func (p *Particle) particule() {
	// Applique les différentes fonctionnalités selon la configuration du config.json

	if config.General.Gravite {
		p.Gravite()
	}

	if config.General.Durreedevie {
		p.Durreedevie()
	}

	if config.General.Lampetorche {
		p.Lampetorche()
	}

	if config.General.Para {
		p.Para()
	}

	if config.General.Interface {
		p.Interface()
	}

	if config.General.Changecolor {
		p.Changecolor()
	}

	if config.General.ChangecolorFR {
		p.ChangecolorFR()
	}

	if config.General.Durreedevie {
		p.Durreedevie()
	}

	if config.General.Dessin {
		p.Dessin()

	}
}
