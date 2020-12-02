package main

type circle struct {
	radius int
}

func (c *circle) getType () string {
	return "Circle"
}