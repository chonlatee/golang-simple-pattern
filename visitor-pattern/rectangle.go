package main

type rectangle struct {
	l int
	b int
}

func (t *rectangle) getType() string {
	return "Rectangle"
}
