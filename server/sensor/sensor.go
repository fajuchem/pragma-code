package sensor

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"server/config"
)

type Sensor struct {
	Id          string `json:"id"`
	Temperature int8   `json:"temperature"`
}

type ResultsItem struct {
	Id  string
	Res Sensor
	Err *error
}

func SendGetSensorAsync(id string, rc chan ResultsItem) {
	var err error
	var sensor Sensor

	defer func() {
		rc <- ResultsItem{
			Id:  id,
			Res: sensor,
			Err: &err,
		}
	}()

	response, err := http.Get(config.SensorBaseUrl + "/" + id)

	if err != nil {
		log.Printf("Internal error: %v", err)
		return
	}

	body, _ := ioutil.ReadAll(response.Body)

	log.Print(string(body))

	if err := json.Unmarshal(body, &sensor); err != nil {
		log.Printf("Internal error: +%v", err)
		return
	}
}
