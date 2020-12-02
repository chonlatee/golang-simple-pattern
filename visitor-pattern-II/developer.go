package main

import "fmt"

type developer struct {
	FirstName string
	LastName string
	Rank int
}

func (d developer) FullName() {
	fmt.Println("Developer: ", d.FirstName, d.LastName)
}

func (d developer) Accept(v visitor) {
	v.VisitorDeveloper(d)
}