package main

import (
	"fmt"
	"time"
)

type mac struct {
	printer printer
}

func (m *mac) print() {
	fmt.Println("Print request for mac")
	time.Sleep(time.Second)
	m.printer.printFile()
}

func (m *mac) setPrinter(p printer) {
	m.printer = p
}
