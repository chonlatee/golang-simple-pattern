package main

import (
	"fmt"
	"time"
)

type hp struct {
}

func (h *hp) printFile() {
	fmt.Println("Printing... file by HP")
	time.Sleep(time.Second * 3)
	fmt.Println("Print completed")
}

