package main

type pizza interface {
	getPrice() float64
	getCalories() int
}

type cheesePizza struct {
	calories int
	price    float64
}

func newCheesePizza() *cheesePizza {
	return &cheesePizza{
		calories: 500,
		price:    10.00,
	}
}

func (c *cheesePizza) getPrice() float64 {
	return c.price
}

func (c *cheesePizza) getCalories() int {
	return c.calories
}
