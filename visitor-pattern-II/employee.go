package main

type employee interface {
	FullName()
	Accept(visitor)
}
