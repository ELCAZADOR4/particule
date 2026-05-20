# R1.01 | Système de Particules en Go (SAE TP1-2)

## Fait par
| Rôle | Nom |
| :--- | :--- |
| Développeur | AILLET Gabriel |
| Développeur | BAUBINEAU Baptiste |

---

## Description du Projet


L’objectif était de créer un système de particules :

- La création et l’initialisation des particules.
- Le mouvement continu.
- La mise à jour régulière de la particule (position, rotation, couleur, etc.. ) 
- Le contrôle du spawn et la gestion des limites de la fenêtre.

---

## Fonctionnalités de la Première Partie

| Fonctionnalité | Description | Fichier(s) Impliqué(s) |
| :--- | :--- | :--- |
| **Contrôle du Spawn** | Les particules apparaissent soit à une position fixe (`SpawnX`, `SpawnY`), soit à une position aléatoire (`RandomSpawn`). | `system.go`, `update.go` |
| **Mouvement et Vitesse** | Chaque particule possède une vitesse aléatoire et met à jour sa position à chaque pas de temps. | `update.go` |
| **Gestion des Limites** | Les particules qui sortent de l'écran sont repositionnées dans les limites (respawn). | `update.go` |
| **Taux de Génération** | De nouvelles particules sont créées à chaque appel à `Update()` selon `SpawnRate`. | `update.go` |
| **Couleurs et Opacité** | Les particules sont créées avec des couleurs aléatoires et une opacité initiale de 1. | `system.go`, `update.go` |
| **Test** | Le code est accompagné de test. | `new_test.go` |

---

## Contrôle du Spawn

Dès le début du projet, nous avons mis en place la gestion de l’apparition des particules.

Si `RandomSpawn = true`, alors les particules apparaissent à une position aléatoire dans la fenêtre, tout en respectant les limites pour éviter qu’elles spawnent en dehors de l’écran.

Sinon, chaque particule apparaît au même point, défini par les paramètres :

```go
SpawnX
SpawnY
```

## Mouvement et vitesse 

Nous avons fait en sorte que chaque particule possède une vitesse initiale aléatoire : 

```go 
vitesseX = rand.Float64()*2 - 1
vitesseY = rand.Float64()*2 - 1
```

De plus nous avons rajouté une rotation à la particule pour que celle-ci ne reste pas indefiniment dans la même potition. 

```go 
p.Rotation +=0.05
```

La rotation a une valeur faible pour que la particule continue à avancer et ne se concentre pas seulement sur la rotation 

## Gestion des limites 

Pour que chaque particule puissent rester dans la limite qui est la taille de la fenetre, nous avons créé des limites lors du spawn. Quand `RandomSpawn = true` chaque particule va apparaitre dans la fenetre et lors qu'elle ont une vitesse si une particule va venir toucher le bord alors elle sera redigiré vers un spawn aleatoire 

```go 
    const particleW float64 = 20.0
	const particleH float64 = 20.0

	windowW := float64(config.General.WindowSizeX)
	windowH := float64(config.General.WindowSizeY)

	diag := math.Sqrt(particleW*particleW + particleH*particleH)
	radius := diag / 2

	minX := radius
	maxX := windowW - radius
	minY := radius
	maxY := windowH - radius

	for i := 0; i < config.General.InitNumParticles; i++ {

		var positionX, positionY float64

		if config.General.RandomSpawn {

			positionX = minX + rand.Float64()*(maxX-minX)
			positionY = minY + rand.Float64()*(maxY-minY)
        }
    }
```

Quand le `RandomSpawn = false` les particules vont apparaitres dans les limites de l'écran mais quand elles vont toucher les limites on va les laissé filer car le SawnRate va servir de generateur ( on en parle plus loin )


## Taux de Génération

Nous avons créé un SpawnRate qui perment de créé des particules `"SpawnRate": 10`, les particules vont soit apparaitre aléatoirement soit sur le lieux du `SpawnX` et `SpawnY` on fonction du `RandomSpawn`

```go 
for i := 0; i < int(config.General.SpawnRate); i++ {
		var x, y float64
		if config.General.RandomSpawn {
			x = minX + rand.Float64()*(maxX-minX)
			y = minY + rand.Float64()*(maxY-minY)
		} else {
			x = float64(config.General.SpawnX)
			y = float64(config.General.SpawnY)
        }
        nouvelleParticule := &Particle{
			PositionX:  x,
			PositionY:  y,
			vitesseX:   rand.Float64()*2 - 1,
			vitesseY:   rand.Float64()*2 - 1,
			ColorRed:   rand.Float64(),
			ColorGreen: rand.Float64(),
			ColorBlue:  rand.Float64(),
			Opacity:    1,
			ScaleX:     1,
			ScaleY:     1,
			Rotation:   rand.Float64() * 2 * math.Pi,
		}
}
```

## Couleurs et Opacité

Pour les couleurs nous avons choisis de la mettre en aléatoire, c'est à dire que chaque particule va avoir une couleur differente 

```go 
	ColorRed:   rand.Float64(),
	ColorGreen: rand.Float64(),
	ColorBlue:  rand.Float64(),
    Opacity = 1 
```

Pour l'opacité nous n'y avons pas touché, chaque particule a la même de 1, c'est à dire au maximum 

## test 

Nous testons donc les deux fichiers que nous avons modifié : `new.go` et `update.go` : 
Pour cela nous avons d'abord avec un fichier qui va venir tester le `new.go` pour 1 particule puis un autre qui va tester le `new.go` pour un ensemble de particule ( test de creation )
Puis le principe est le même pour le teste du `update.go` qui va tester la vitesse, le SpawnRate etc...