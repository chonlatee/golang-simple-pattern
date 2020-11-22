package main

type CellPhone struct {
	Camera bool
	DualSim bool
	Torch bool
	ColorDisplay bool
}

type BuildProcess interface {
	SetCamera() BuildProcess
	SetDualSim() BuildProcess
	SetTorch() BuildProcess
	SetColorDisplay() BuildProcess
	GetCellPhone() CellPhone
}

