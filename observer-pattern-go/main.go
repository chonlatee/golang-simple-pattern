package main

func main() {
	coSensor := createWeatherSensor()
	wySensor := createWeatherSensor()
	denverStation := createWeatherStation("denver")
	vailStation := createWeatherStation("vail")
	cheyenneStation := createWeatherStation("cheyenne")

	coSensor.addStation(denverStation)
	coSensor.addStation(vailStation)
	wySensor.addStation(cheyenneStation)

	coSensor.NotifyAll()
	wySensor.NotifyAll()


	// change temperature for a new day
	coSensor.ChangeTemperature()
	wySensor.ChangeTemperature()

	coSensor.removeStation(vailStation)

	coSensor.ChangeTemperature()
}
