package main

type square struct {
	side int
}

func (s *square) getType() string {
	return "Square"
}

func (s *square) accept(v visitor) {
	v.visitorForSquare(s)
}