package main

const (
	pepperoniPrice = 1.75
	artichokePrice = 2.00
	anchoviePrice  = 2.50
	tomatoePrice   = 0.75

	pepperoniCalories = 200
	artichokeCalories = 150
	anchovieCalories  = 350
	tomatoeCalories   = 50
)

type pepperoniTopping struct {
	pizza pizza
}

type artichokeTopping struct {
	pizza pizza
}

type archoiveTopping struct {
	pizza pizza
}

type tomatoeTopping struct {
	pizza pizza
}

func (p *pepperoniTopping) getPrice() float64 {
	return p.pizza.getPrice() + pepperoniPrice
}

func (a *artichokeTopping) getPrice() float64 {
	return a.pizza.getPrice() + anchoviePrice
}

func (a *archoiveTopping) getPrice() float64 {
	return a.pizza.getPrice() + anchoviePrice
}

func (t *tomatoeTopping) getPrice() float64 {
	return t.pizza.getPrice() + tomatoePrice
}

func (p *pepperoniTopping) getCalories() int {
	return p.pizza.getCalories() + pepperoniCalories
}

func (a *artichokeTopping) getCalories() int {
	return a.pizza.getCalories() + artichokeCalories
}

func (a *archoiveTopping) getCalories() int {
	return a.pizza.getCalories() + artichokeCalories
}

func (t *tomatoeTopping) getCalories() int {
	return t.pizza.getCalories() + tomatoeCalories
}
