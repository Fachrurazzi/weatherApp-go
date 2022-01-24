package model

type Weather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Pressure int     `json:"pressure"`
	Humidity int     `json:"humidity"`
	TempMin  float64 `json:"temp_Min"`
	TempMax  float64 `json:"temp_Max"`
}

type CurrentWeatherResponse struct {
	Name    string    `json:"name"`
	Weather []Weather `json:"weather"`
	Main    Main      `json:"main"`
}
