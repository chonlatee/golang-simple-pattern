package main

type areaCalculator struct {
	area int
}

func (a *areaCalculator) visitorForSquare(s *square) {
	a.area = s.side * s.side
}

func (a *areaCalculator) GetArea() int {
	return a.area
}
