package main

import "fmt"

type middleCoordinates struct {
	x int
	y int
}


func (m *middleCoordinates) visitorForSquare(s *square) {
	m.x = 2
	m.y = 2
	fmt.Println("Coordinates x, y", m.x, m.y)
}