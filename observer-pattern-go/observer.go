package main

import "fmt"

type Observer interface {
	Update(value int)
}

type WeatherStation struct {
	Name string
}

func createWeatherStation(name string) *WeatherStation {
	return &WeatherStation{
		Name: name,
	}
}

func (ws *WeatherStation) Update(temperature int) {
	fmt.Printf("%s reporting: \n", ws.Name)

	if temperature >= 85 {
		fmt.Printf("\tThe temperature is %d F, it will be a hot today!\n", temperature)
	} else if temperature >= 55 && temperature < 85 {
		fmt.Printf("\tThe temperature is %d F, it will be a mild day.\n", temperature)
	} else {
		fmt.Printf("\tThe temperature is %d F, it will be a cold day.\n", temperature)
	}
}
