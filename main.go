package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"weatherApp/controller"
	"weatherApp/model"
)

func main() {
	router := httprouter.New()

	var currentWeather *model.CurrentWeatherResponse
	var err error
	apikey := "6868edc853c17d3430fbfa9117d661a5"

	owm := controller.OpenWeatherMap{
		APIKEY: apikey,
	}

	router.GET("/api/weather/:city", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		city := params.ByName("city")
		currentWeather, err = owm.CurrentWeatherFromCity(city)
		if err != nil {
			log.Fatal(err)
		}
		bytes, _ := json.Marshal(currentWeather)
		fmt.Fprint(writer, string(bytes))
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
