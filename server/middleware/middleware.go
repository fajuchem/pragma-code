package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"server/models"
)

func GetAllProducts(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var beers []models.Beer

	beer1 := models.Beer{
		Id:                 "1",
		Name:               "Pilsner",
		MinimumTemperature: 4,
		MaximumTemperature: 6,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beer2 := models.Beer{
		Id:                 "2",
		Name:               "IPA",
		MinimumTemperature: 5,
		MaximumTemperature: 6,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beer3 := models.Beer{
		Id:                 "3",
		Name:               "Lager",
		MinimumTemperature: 4,
		MaximumTemperature: 7,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beer4 := models.Beer{
		Id:                 "4",
		Name:               "Stout",
		MinimumTemperature: 6,
		MaximumTemperature: 8,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beer5 := models.Beer{
		Id:                 "5",
		Name:               "Wheat beer",
		MinimumTemperature: 3,
		MaximumTemperature: 5,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beer6 := models.Beer{
		Id:                 "6",
		Name:               "Pale Ale",
		MinimumTemperature: 4,
		MaximumTemperature: 6,
		Temperature:        0,
		TemperatureStatus:  "",
	}

	beers = append(beers, beer1, beer2, beer3, beer4, beer5, beer6)

	resultChannel := make(chan models.ResultsItem)

	for i, beer := range beers {
		fmt.Println(i, beer.Name)

		go SendGetSensorAsync(beer.Id, resultChannel)
	}

	result := make(map[string]int8)

	for range beers {
		item := <-resultChannel
		result[item.Id] = item.Res.Temperature
	}

	for j, beer := range beers {
		beers[j].Temperature = result[beer.Id]

		fmt.Println(result[beer.Id])
		fmt.Println(beer)

		if result[beer.Id] < beer.MinimumTemperature {
			beers[j].TemperatureStatus = "too low"
		}

		if result[beer.Id] > beer.MaximumTemperature {
			beers[j].TemperatureStatus = "too high"
		}

		if result[beer.Id] >= beer.MinimumTemperature && result[beer.Id] <= beer.MaximumTemperature {
			beers[j].TemperatureStatus = "all good"
		}
	}

	json.NewEncoder(w).Encode(beers)
}

func SendGetSensorAsync(id string, rc chan models.ResultsItem) {
	response, err := http.Get("https://temperature-sensor-service.herokuapp.com/sensor/" + id)
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)

	var sensor models.Sensor
	if err := json.Unmarshal(body, &sensor); err != nil {
		panic(err)
	}
	rc <- models.ResultsItem{Id: id, Res: sensor}
}
