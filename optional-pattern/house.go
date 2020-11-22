package main

type House struct {
	material string
	hasFireplace bool
	floors int
}

const defaultFloors = 2
const defaultHasFireplace = true
const defaultMaterial = "wood"

type HouseOption func(house *House)

func NewHouse(options ...HouseOption) *House {
	h := &House{
		material:     defaultMaterial,
		hasFireplace: defaultHasFireplace,
		floors:       defaultFloors,
	}

	for _, o := range options {
		o(h)
	}

	return h
}

func WithConcrete() HouseOption {
	return func(h *House) {
		h.material = "concrete"
	}
}

func WithoutFireplace() HouseOption {
	return func(h *House) {
		h.hasFireplace = false
	}
}

func WithFloors(floors int) HouseOption {
	return func(h *House) {
		h.floors = floors
	}
}
