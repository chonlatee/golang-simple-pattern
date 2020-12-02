package main

import "fmt"

type director struct {
	FirstName string
	LastName string
	Rank int
}

func (d director) FullName() {
	fmt.Println("Developer: ", d.FirstName, d.LastName)
}

func (d director) Accept(v visitor) {
	v.VisitorDirector(d)
}