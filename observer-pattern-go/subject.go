package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	maxTemperature = 120
	minTemperature = -30
)

type Subject interface {
	NotifyAll()
}

type WeatherSensor struct {
	Subject
	weatherStations []*WeatherStation
	temperature     int
}

func createWeatherSensor() *WeatherSensor {
	return &WeatherSensor{
		temperature: getRandomTemperature(),
	}
}

func (ws *WeatherSensor) addStation(station *WeatherStation) {
	ws.weatherStations = append(ws.weatherStations, station)
}

func (ws *WeatherSensor) removeStation(removeStation *WeatherStation) {
	var w []*WeatherStation
	for _, s := range ws.weatherStations {
		if s.Name != removeStation.Name {
			w = append(w, s)
		}
	}

	ws.weatherStations = w
}

func (ws *WeatherSensor) NotifyAll() {
	for _, s := range ws.weatherStations {
		s.Update(ws.temperature)
	}
}

func (ws *WeatherSensor) ChangeTemperature() {
	fmt.Printf("\nIt's a new day!\n")
	ws.temperature = getRandomTemperature()
	ws.NotifyAll()
}

func getRandomTemperature() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxTemperature-minTemperature) + minTemperature
}
