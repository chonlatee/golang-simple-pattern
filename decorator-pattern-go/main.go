package main

import "fmt"

func main() {

	theBombPizza := &archoiveTopping{
		pizza: &artichokeTopping{
			pizza: &tomatoeTopping{
				pizza: newCheesePizza(),
			},
		},
	}

	peperoniPizza := &pepperoniTopping{
		pizza: newCheesePizza(),
	}

	fmt.Printf("Price of theBombPizza is %.2f with %d calories\n", theBombPizza.getPrice(), theBombPizza.getCalories())
	fmt.Printf("Price of peperoniPizza is %.2f with %d calories\n", peperoniPizza.getPrice(), peperoniPizza.getCalories())

}
