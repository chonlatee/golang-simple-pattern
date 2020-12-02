package main

type visitor interface {
	VisitorDeveloper(d developer)
	VisitorDirector(d director)
}


