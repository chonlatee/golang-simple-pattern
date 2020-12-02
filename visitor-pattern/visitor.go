package main

type visitor interface {
	visitorForSquare(*square)
}