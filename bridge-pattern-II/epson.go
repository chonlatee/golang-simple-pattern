package main

import (
	"fmt"
	"time"
)

type epson struct {

}

func (e *epson) printFile() {
	fmt.Println("Printing... file by EPSON")
	time.Sleep(time.Second * 3)
	fmt.Println("Print Completed")
}

