package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type Train struct {
	Axe             string `json:"axe"`
	Date            string `json:"date"`
	Destination     string `json:"destination"`
	DestinationIATA string `json:"destination_iata"`
	Entity          string `json:"entity"`
	ArrivalTime     string `json:"heure_arrivee"`
	DepartureTime   string `json:"heure_depart"`
	OdHappyCard     string `json:"od_happy_card"`
	Origine         string `json:"origine"`
	OrigineIATA     string `json:"origine_iata"`
	TrainNumber     string `json:"train_no"`
}

type Response struct {
	Results []Train `json:"results"`
}

// MakeRequest performs an HTTP GET request
func MakeRequest(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failedd to fetch data from API")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
