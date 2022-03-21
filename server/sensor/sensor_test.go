package sensor

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"server/config"
	"testing"
)

func mockSensorApi(response string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, response)
	}))
	config.SensorBaseUrl = ts.URL
	return ts
}

func TestSendGetSensorAsync(t *testing.T) {
	ts := mockSensorApi(`{"id":"6","temperature":5}`)
	defer ts.Close()

	resultChannel := make(chan ResultsItem)
	go SendGetSensorAsync("6", resultChannel)
	item := <-resultChannel

	if item.Res.Temperature != 5 {
		t.Error("Expect Temperature to be 5")
	}
}

func TestSendGetSensorAsyncSensorApiFail(t *testing.T) {
	ts := mockSensorApi(`{"id":"6","temperature":5}`)
	config.SensorBaseUrl = "invalid url"
	defer ts.Close()

	resultChannel := make(chan ResultsItem)
	go SendGetSensorAsync("6", resultChannel)
	item := <-resultChannel

	if item.Err == nil {
		t.Error("Expect api error")
	}
}

func TestSendGetSensorAsyncFailMarshal(t *testing.T) {
	ts := mockSensorApi("")
	defer ts.Close()

	resultChannel := make(chan ResultsItem)
	go SendGetSensorAsync("6", resultChannel)
	item := <-resultChannel

	if item.Err == nil {
		t.Error("Expect marshal error")
	}
}
