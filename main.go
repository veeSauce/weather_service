package main

// This is the main package for the weather service application.
// It sets up an HTTP server that listens for incoming requests,
// serves a landing page with a form for users to input latitude and longitude,
// and handles form submissions to fetch and display weather forecast data based
// on the provided coordinates.
// Service runs on localhost:8080

// Assumption:
// 1. Only data from USA will be provided using the NWS in Farenheit
// for US Land Coordinates.
// 2. It will return the forecast for the latest time of the day that the service is accessed.
// For example during the day it will not return the night and vice versa.

import (
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"text/template"

	"github.com/veeSauce/Weather_service/page_models"
)

type Data struct {
	Title        string
	ForecastData string
	TempFeeling  string
}

// Main function to set up the HTTP server and route handlers
func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", inputForm)
	mux.HandleFunc("/submit", submitForm)

	// run the server and log to console if the server crashes
	log.Fatal(http.ListenAndServe(":8080", mux))

}

// Handler function to serve the landing page with the input form
func inputForm(w http.ResponseWriter, r *http.Request) {

	// wc, err := w.Write([]byte("Hello, World!\n"))
	// if err != nil {
	// 	slog.Error("Error writing response", "error", err)
	// 	return
	// }

	// fmt.Printf("Bytes written: %d\n", wc)

	t, err := template.ParseFiles("page/landingPage.html")
	if err != nil {
		slog.Error("Error parsing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		slog.Error("Error executing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}

// Handler function to process form submissions, fetch weather data, and display the forecast
func submitForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	latitude := r.FormValue("latitude")
	longitude := r.FormValue("longitude")

	slog.Info("Received form submission", "latitude", latitude, "longitude", longitude)

	link := "https://api.weather.gov/points/%s,%s"
	link = fmt.Sprintf(link, latitude, longitude)

	resp, err := http.Get(link)
	if err != nil {
		http.Error(w, "Error making API request", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	// Extract the URL from the API response
	newURL, err := extractCoordinates(resp)
	if err != nil {
		http.Error(w, "Error extracting coordinates", http.StatusInternalServerError)
		return
	}

	// Make a new API request to the extracted URL
	resp2, err := http.Get(newURL)
	if err != nil {
		http.Error(w, "Error making second API request", http.StatusInternalServerError)
		return
	}

	defer resp2.Body.Close()

	// Need to extract the new response and map it to the forecasest model, then display the forecast data on the webpage
	forecastData, tempFeeling, _, err := extractForecastData(resp2)
	if err != nil {
		slog.Error("Error extracting forecast data", "error", err)
		http.Error(w, "Error extracting forecast data", http.StatusInternalServerError)
		return
	}

	renderResponsePage(w, forecastData, tempFeeling)

}

// Function to render the response page with forecast data
func renderResponsePage(w http.ResponseWriter, forecastData, tempFeeling string) {

	t, err := template.ParseFiles("page/responsePage.html")
	if err != nil {
		slog.Error("Error parsing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := Data{
		Title:        "Weather Forecast Service",
		ForecastData: forecastData,
		TempFeeling:  tempFeeling,
	}

	err = t.Execute(w, data)
	if err != nil {
		slog.Error("Error executing template", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// Function to extract the forecast URL from the API response
func extractCoordinates(resp *http.Response) (string, error) {

	var data page_models.CoordinateModel

	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return "", err
	}

	// slog.Info("Extracted data from API response", "data", data)

	return data.Properties.Forecast, nil

}

// Function to extract the forecast data from the API response and determine the time of day and temperature feeling
func extractForecastData(resp *http.Response) (string, string, string, error) {

	var data page_models.ForecastModel

	// Experiment
	// b, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	slog.Error("Error reading response body", "error", err)
	// 	return "", "", err
	// }
	// slog.Info("Attempting to decode forecast data from response body", "body", string(b))

	// return string(b), "night", nil

	err := json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		slog.Error("Error decoding forecast data from API response", "error", err)
		return "", "", "", err
	}

	slog.Info("Successfully decoded forecast data from API response", "data", data)

	if len(data.Properties.Periods) == 0 {
		return "", "", "", fmt.Errorf("No Forecast periods available")
	}

	var time_of_day string

	if data.Properties.Periods[0].Name == "Tonight" {

		slog.Info("It's currently nighttime based on forecast data")
		time_of_day = "night"

	} else {
		slog.Info("It's currently daytime based on forecast data")
		time_of_day = "day"
	}

	shortForecast := data.Properties.Periods[0].ShortForecast
	temp := data.Properties.Periods[0].Temperature
	tempUnit := data.Properties.Periods[0].TemperatureUnit

	tempFeeling, err := temperatureFeeling(temp, tempUnit)
	if err != nil {
		slog.Error("Error getting temperature feeling", "error", err)
		return "", "", "", err
	}

	slog.Info("Extracted shortForecast from API response", "data", shortForecast)

	return shortForecast, tempFeeling, time_of_day, nil

}

func temperatureFeeling(temp int, tempUnit string) (string, error) {

	if tempUnit == "F" {

		if temp <= 32 {
			return "cold", nil
		} else if temp > 32 && temp <= 70 {
			return "moderate", nil
		} else {
			return "hot", nil
		}

	} else {

		return "", fmt.Errorf("Unsupported temperature unit for USA: %s", tempUnit)
	}
}
