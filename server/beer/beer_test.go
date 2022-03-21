package beer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"server/config"
	"testing"
)

func TestGetProducts(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "")
	}))
	config.SensorBaseUrl = ts.URL
	defer ts.Close()

	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	GetAllProducts(rr, req)

	var brs []Beer
	body, _ := ioutil.ReadAll(rr.Body)

	err = json.Unmarshal(body, &brs)
	if err != nil {
		t.Fatal(err)
	}

	if brs[0].Id != "1" {
		t.Error("expected return a beer")
	}
	if brs[0].Name != "Pilsner" {
		t.Error("expected return a beer")
	}
}

func TestGetStatusMinTemperature(t *testing.T) {
	status := GetBeerStatus(1, 2, 2)
	if status != "too low" {
		t.Error("expected too low")
	}
}

func TestGetStatusMaxTemperature(t *testing.T) {
	status := GetBeerStatus(3, 2, 2)
	if status != "too high" {
		t.Error("expected too high")
	}
}

func TestGetStatusTemperatureGood(t *testing.T) {
	status := GetBeerStatus(2, 1, 3)
	if status != "all good" {
		t.Error("expected all good")
	}
}
