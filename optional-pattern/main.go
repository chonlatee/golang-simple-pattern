package main

import "fmt"

func main() {

	h := NewHouse()
	fmt.Printf("House %+v\n", h)
	options := []HouseOption{
		WithConcrete(),
		WithoutFireplace(),
	}
	nh := NewHouse(options...)
	fmt.Printf("House %+v\n", nh)


}
