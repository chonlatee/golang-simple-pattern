package main

import "fmt"

type Director struct {
	builder BuildProcess
}

func (d *Director) SetBuilder(b BuildProcess) {
	d.builder = b
}

func (d *Director) Construct() CellPhone {
	d.builder.SetCamera().SetDualSim().SetTorch().SetColorDisplay()
	return d.builder.GetCellPhone()
}

func main() {
	diro := &Director{}
	nokia := &Nokia{}
	diro.SetBuilder(nokia)

	phone := diro.Construct()
	fmt.Printf("Nokia phone %+v\n", phone)

	samsung := &Samsung{}
	diro.SetBuilder(samsung)
	phone = diro.Construct()
	fmt.Printf("Samsung phone %+v\n", phone)
}