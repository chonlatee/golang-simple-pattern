package main

import "fmt"

type DialogueReciter interface {
	Recite()
}

type spiderMan struct{}

func (s spiderMan) Recite() {
	fmt.Println("No Man Can With Every Battle, But No Man Should Fall Without A Struggle")
}

type superMan struct{}

func (s superMan) Recite() {
	fmt.Println("There is a superhero in all of us, we just need the courage to put on the cape")
}

type batMan struct{}

func (b batMan) Recite() {
	fmt.Println("It's not who I am underneath, but what I do that defines me!")
}

type toy struct {
	DialogueReciter DialogueReciter
}

func NewToy(dr DialogueReciter) *toy {
	return &toy{
		DialogueReciter: dr,
	}
}

func (t *toy) PerformDialogue() {
	t.DialogueReciter.Recite()
}

func (t *toy) SetSuperHero(dr DialogueReciter) {
	t.DialogueReciter = dr
}

func main() {
	t := NewToy(spiderMan{})
	t.PerformDialogue()
	t.SetSuperHero(superMan{})
	t.PerformDialogue()
	t.SetSuperHero(batMan{})
	t.PerformDialogue()
}

