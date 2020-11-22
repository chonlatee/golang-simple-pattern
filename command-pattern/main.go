package main

import "fmt"

func main() {
	r := NewRestaurant()

	tasks := []Command{
		r.MakePizza(2),
		r.MakeSalad(2),
		r.MakePizza(3),
		r.CleanDishes(),
		r.MakePizza(4),
		r.CleanDishes(),
	}

	cooks := []*Cook{
		&Cook{},
		&Cook{},
	}

	for i, task := range tasks {
		cook := cooks[i%len(cooks)]
		cook.Commands = append(cook.Commands, task)
	}

	for i, c := range cooks {
		fmt.Println("cook", i, ":")
		c.executeCommand()
	}
}