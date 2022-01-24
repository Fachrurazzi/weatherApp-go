package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"weatherApp/controller"
	"weatherApp/helper"
	"weatherApp/model"
)

func main() {
	router := chi.NewRouter()

	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading env file")
	}

	owm := controller.OpenWeatherMap{
		APIKEY: os.Getenv("API_KEY"),
		APIURL: os.Getenv("API_URL"),
	}
	router.Get("/api/weather/{city}", func(writer http.ResponseWriter, request *http.Request) {
		city := chi.URLParam(request, "city")
		fmt.Fprint(writer, city)
		var currentWeather *model.CurrentWeatherResponse
		var err error

		fmt.Fprint(writer, city)
		currentWeather, err = owm.CurrentWeatherFromCity(city)
		if err != nil {
			log.Fatal(err)
		}
		webResponse := model.WebResponse{
			Code:   200,
			Status: "OK",
			Data:   currentWeather,
		}

		helper.WriteToResponseBody(writer, webResponse)
		//writer.Write(weather)
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
