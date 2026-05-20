package particles

func (p *Particle) Durreedevie() {
	p.ScaleX -= 0.004
	p.ScaleY -= 0.004
	p.Age++

	if p.Age > 240 {
		p.Dead = true
	}
}
