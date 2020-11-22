package main

type Nokia struct {
	Phone CellPhone
}

func (n *Nokia) SetCamera() BuildProcess {
	n.Phone.Camera = false
	return n
}

func (n *Nokia) SetDualSim() BuildProcess {
	n.Phone.DualSim = false
	return n
}

func (n *Nokia) SetTorch() BuildProcess {
	n.Phone.Torch = true
	return n
}

func (n *Nokia) SetColorDisplay() BuildProcess {
	n.Phone.ColorDisplay = false
	return n
}

func (n *Nokia) GetCellPhone() CellPhone {
	return n.Phone
}