package middleware

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

	for _, beer := range beers {
		go SendGetSensorAsync(beer.Id, resultChannel)
	}

	result := make(map[string]int8)

	for range beers {
		item := <-resultChannel
		result[item.Id] = item.Res.Temperature
	}

	for j, beer := range beers {
		beers[j].Temperature = result[beer.Id]

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

type spaHandler struct {
	staticPath string
	indexPath  string
}

// ServeHTTP inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
