package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"weatherApp/helper"
	"weatherApp/model"
)

const (
	APIURL string = "api.openweathermap.org"
)

type OpenWeatherMap struct {
	APIKEY string
}

func (open *OpenWeatherMap) CurrentWeatherFromCity(city string) (*model.CurrentWeatherResponse, error) {
	if open.APIKEY == "" {
		return nil, errors.New("No API keys present")
	}

	url := fmt.Sprintf("https://%s/data/2.5/weather?q=%s&appid=%s", APIURL, city, open.APIKEY)

	body, err := helper.MakeApiRequest(url)
	if err != nil {
		return nil, err
	}

	var currentWeatherResponse model.CurrentWeatherResponse

	err = json.Unmarshal(body, &currentWeatherResponse)
	if err != nil {
		return nil, err
	}

	return &currentWeatherResponse, nil
}
