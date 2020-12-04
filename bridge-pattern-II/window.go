package main

import (
	"fmt"
	"time"
)

type window struct {
	printer printer
}

func (w *window) print() {
	fmt.Println("Print request from window")
	time.Sleep(time.Second)
	w.printer.printFile()
}

func (w *window) setPrinter(p printer) {
	w.printer = p
}

