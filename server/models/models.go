package models

type Beer struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	MinimumTemperature int8   `json:"minimumTemperature"`
	MaximumTemperature int8   `json:"maximumTemperature"`
	Temperature        int8   `json:"temperature"`
	TemperatureStatus  string `json:"temperatureStatus"`
}

type Sensor struct {
	Id          string `json:"id"`
	Temperature int8   `json:"temperature"`
}

type ResultsItem struct {
	Id  string
	Res Sensor
}
