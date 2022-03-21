package beer

import (
	"encoding/json"
	"net/http"
	"server/sensor"
)

type Beer struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	MinimumTemperature int8   `json:"minimumTemperature"`
	MaximumTemperature int8   `json:"maximumTemperature"`
	Temperature        int8   `json:"temperature"`
	TemperatureStatus  string `json:"temperatureStatus"`
}

var AvailableBeers = []Beer{
	{
		Id:                 "1",
		Name:               "Pilsner",
		MinimumTemperature: 4,
		MaximumTemperature: 6,
		Temperature:        0,
		TemperatureStatus:  "",
	},
	{
		Id:                 "2",
		Name:               "IPA",
		MinimumTemperature: 5,
		MaximumTemperature: 6,
		Temperature:        0,
		TemperatureStatus:  "",
	},
	{
		Id:                 "3",
		Name:               "Lager",
		MinimumTemperature: 4,
		MaximumTemperature: 7,
		Temperature:        0,
		TemperatureStatus:  "",
	},
	{
		Id:                 "4",
		Name:               "Stout",
		MinimumTemperature: 6,
		MaximumTemperature: 8,
		Temperature:        0,
		TemperatureStatus:  "",
	},
	{
		Id:                 "5",
		Name:               "Wheat beer",
		MinimumTemperature: 3,
		MaximumTemperature: 5,
		Temperature:        0,
		TemperatureStatus:  "",
	},
	{
		Id:                 "6",
		Name:               "Pale Ale",
		MinimumTemperature: 4,
		MaximumTemperature: 6,
		Temperature:        0,
		TemperatureStatus:  "",
	},
}

func GetAllProducts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	resultChannel := make(chan sensor.ResultsItem)

	for _, beer := range AvailableBeers {
		go sensor.SendGetSensorAsync(beer.Id, resultChannel)
	}

	result := make(map[string]int8)

	for range AvailableBeers {
		item := <-resultChannel
		result[item.Id] = item.Res.Temperature
	}

	beers := AvailableBeers

	for j, b := range beers {
		beers[j].Temperature = result[b.Id]
		beers[j].TemperatureStatus = GetBeerStatus(result[b.Id], b.MinimumTemperature, b.MaximumTemperature)
	}

	json.NewEncoder(w).Encode(beers)
}

func GetBeerStatus(id int8, minTemperature int8, maxTemperature int8) string {
	if id < minTemperature {
		return "too low"
	}

	if id > maxTemperature {
		return "too high"
	}

	return "all good"
}
