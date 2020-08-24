package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const apiKey = "a4ca47eafe1fc0cf1a3b66a113330131"

// Weather struct
// Generated with mholt.github.io/json-to-go/
type Weather struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func get(loc string) Weather {
	var apiCall = "http://api.openweathermap.org/data/2.5/weather?q=" + loc + "&appid=" + apiKey
	resp, err := http.Get(apiCall)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to Todo struct
	var weatherStruct Weather
	json.Unmarshal(bodyBytes, &weatherStruct)

	return weatherStruct
}

func main() {
	// We take the input of the location
	location := os.Args[1:]

	// We help the user by showing the valid inputs on (go run *.go -h)
	flagLocation := flag.String("location", "Barcelona", "a valid location")
	flag.Parse()
	_ = flagLocation

	_ = location
	var weatherStruct Weather
	weatherStruct = get(strings.Join(location, " "))

	fmt.Printf("%s: %s\nCurrent temperature: %f\tMax: %f\tMin: %f\n",
		weatherStruct.Weather[0].Main,
		weatherStruct.Weather[0].Description,
		weatherStruct.Main.Temp,
		weatherStruct.Main.TempMax,
		weatherStruct.Main.TempMax)
}
