package main

import "fmt"

type shape interface {
	getType() string
	accept(visitor)
}




func main() {
	square := &square{side: 2}
	fmt.Println(square.getType())

	areaCalculator := &areaCalculator{}
	square.accept(areaCalculator)

	fmt.Println(areaCalculator.GetArea())

}