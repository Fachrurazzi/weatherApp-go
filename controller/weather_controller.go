package controller

import (
	"errors"
	"fmt"
	"weatherApp/helper"
	"weatherApp/model"
)

type OpenWeatherMap struct {
	APIKEY string
	APIURL string
}

func (open *OpenWeatherMap) CurrentWeatherFromCity(city string) (*model.CurrentWeatherResponse, error) {
	if open.APIKEY == "" {
		return nil, errors.New("No API keys present")
	}

	url := fmt.Sprintf("https://%s/data/2.5/weather?q=%s&appid=%s&units=metric", open.APIURL, city, open.APIKEY)

	body, err := helper.MakeApiRequest(url)
	if err != nil {
		return nil, err
	}

	var currentWeatherResponse model.CurrentWeatherResponse

	//err = json.Unmarshal(body, &currentWeatherResponse)
	//if err != nil {
	//	return nil, err
	//}

	helper.ToDecode(body, &currentWeatherResponse)

	return &currentWeatherResponse, nil
}
