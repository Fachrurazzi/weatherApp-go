package model

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_Min"`
	TempMax  float64 `json:"temp_Max"`
}

type Sys struct {
	Country string `json:"country"`
}

type CurrentWeatherResponse struct {
	Name    string    `json:"name"`
	Sys     Sys       `json:"sys"`
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}
