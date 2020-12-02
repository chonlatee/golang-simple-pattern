package main


type bonus struct {
	val float32
}

func (b *bonus) VisitorDirector(d director) {
	b.val = float32(d.Rank) * 0.8
}

func (b *bonus) VisitorDeveloper(d developer) {
	b.val = float32(d.Rank) * 0.5
}

func (b *bonus) GetBonus() float32 {
	return b.val
}
