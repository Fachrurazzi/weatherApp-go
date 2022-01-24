package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
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

	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Get("/api/weather/{city}", func(writer http.ResponseWriter, request *http.Request) {
		city := chi.URLParam(request, "city")
		var currentWeather *model.CurrentWeatherResponse
		var err error

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
	})

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}

}
