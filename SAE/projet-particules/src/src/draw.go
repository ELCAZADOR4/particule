package main

import (
	"fmt"
	"project-particles/assets"
	"project-particles/config"
	"project-particles/particles"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *game) Draw(screen *ebiten.Image) {
	for e := g.system.Content.Front(); e != nil; e = e.Next() {
		p, ok := e.Value.(*particles.Particle)
		if ok {
			options := ebiten.DrawImageOptions{}
			options.GeoM.Rotate(p.Rotation)
			options.GeoM.Scale(p.ScaleX, p.ScaleY)
			options.GeoM.Translate(p.PositionX, p.PositionY)
			options.ColorScale.Scale(
				float32(p.ColorRed),
				float32(p.ColorGreen),
				float32(p.ColorBlue),
				float32(p.Opacity),
			)
			screen.DrawImage(assets.ParticleImage, &options)
		}
	}

	// Permet d'afficher l'interface
	if config.General.Interface {
		y := 10 // Position verticale initiale
		line := func(text string) {
			ebitenutil.DebugPrintAt(screen, text, 10, y)
			y += 15
		}
		// Affichage des informations de l'interface
		line("INTERFACE PARTICULES")
		line(fmt.Sprintf("Nombre de particules : %d", g.system.Content.Len()))
		line("SpawnRate : A/Q +/-")
		line("Touches pour couleur : R/O = rouge +/- | V/H = vert +/- | B/N = bleu +/-")
		line("Touches pour taille  : flèches haut/bas/left/right")
		line(fmt.Sprintf("Gravite : %v (appuyez sur G )", config.General.Gravite))
		line(fmt.Sprintf("Changement de couleur : %v (appuyez sur C )", config.General.Changecolor))
		line(fmt.Sprintf("Durée de vie : %v (appuyez sur D )", config.General.Durreedevie))
		line(fmt.Sprintf("Lampe torche : %v (appuyez sur L )", config.General.Lampetorche))
		line(fmt.Sprintf("Parapluie : %v (appuyez sur P )", config.General.Para))
		line(fmt.Sprintf("Drapeau de la France : %v (appuyez sur F )", config.General.ChangecolorFR))
		line(fmt.Sprintf("Cercle : %v (appuyez sur E )", config.General.Dessin))
	}
}
