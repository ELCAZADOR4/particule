package particles

import (
	"project-particles/config"

	"github.com/hajimehoshi/ebiten/v2"
)

// Interface permet de modifier les paramètres des particules via le clavier
func (p *Particle) Interface() {
	if !config.General.Interface {
		return
	}

	// Couleur
	if ebiten.IsKeyPressed(ebiten.KeyR) { // rouge +
		p.ColorRed += 0.05
		if p.ColorRed > 1 {
			p.ColorRed = 1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyO) { // rouge -
		p.ColorRed -= 0.05
		if p.ColorRed < 0 {
			p.ColorRed = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyV) { // vert +
		p.ColorGreen += 0.05
		if p.ColorGreen > 1 {
			p.ColorGreen = 1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyH) { // vert -
		p.ColorGreen -= 0.05
		if p.ColorGreen < 0 {
			p.ColorGreen = 0
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyB) { // bleu +
		p.ColorBlue += 0.05
		if p.ColorBlue > 1 {
			p.ColorBlue = 1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyN) { // bleu -
		p.ColorBlue -= 0.05
		if p.ColorBlue < 0 {
			p.ColorBlue = 0
		}
	}

	// Taille
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.ScaleY += 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.ScaleY -= 0.1
		if p.ScaleY < 0.1 {
			p.ScaleY = 0.1
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.ScaleX += 0.1
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.ScaleX -= 0.1
		if p.ScaleX < 0.1 {
			p.ScaleX = 0.1
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyG) {
		config.General.Gravite = !config.General.Gravite
	}

	if ebiten.IsKeyPressed(ebiten.KeyP) {
		config.General.Para = !config.General.Para
	}

	if ebiten.IsKeyPressed(ebiten.KeyC) {
		config.General.Changecolor = !config.General.Changecolor
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		config.General.Durreedevie = !config.General.Durreedevie
	}

	if ebiten.IsKeyPressed(ebiten.KeyL) {
		config.General.Lampetorche = !config.General.Lampetorche
	}

	if ebiten.IsKeyPressed(ebiten.KeyF) {
		config.General.ChangecolorFR = !config.General.ChangecolorFR
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		config.General.Changecolor = !config.General.Changecolor
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		config.General.SpawnRate += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyQ) {
		config.General.SpawnRate -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		config.General.Changecolor = !config.General.Changecolor
	}

	if ebiten.IsKeyPressed(ebiten.KeyE) {
		config.General.Dessin = !config.General.Dessin
	}
}
