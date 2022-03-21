package config

var SensorBaseUrl string

func init() {
	SensorBaseUrl = "https://temperature-sensor-service.herokuapp.com/sensor"
}
