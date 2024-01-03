package main

import (
	"html/template"
	"net/http"
)

func MainPage(w http.ResponseWriter, r *http.Request) {

	parseFile, err := template.ParseFiles("../../ui/html/index.html")
	if err != nil {
		return
	}
	err = parseFile.Execute(w, nil)
	if err != nil {
		return
	}
}

func CalibrationIntroPage(w http.ResponseWriter, r *http.Request) {
	parseFile, err := template.ParseFiles("../../ui/html/calibration-intro.html")
	if err != nil {
		return
	}
	err = parseFile.Execute(w, nil)
	if err != nil {
		return
	}
}

func CalibrationPage(w http.ResponseWriter, r *http.Request) {
	parseFile, err := template.ParseFiles("../../ui/html/calibration.html")
	if err != nil {
		return
	}
	err = parseFile.Execute(w, nil)
	if err != nil {
		return
	}
}

func ContentPage(w http.ResponseWriter, r *http.Request) {
	parseFile, err := template.ParseFiles("../../ui/html/content.html")
	if err != nil {
		return
	}
	err = parseFile.Execute(w, nil)
	if err != nil {
		return
	}
}

func CinemaPage(w http.ResponseWriter, r *http.Request) {
	parseFile, err := template.ParseFiles("../../ui/html/cinema.html")
	if err != nil {
		return
	}
	err = parseFile.Execute(w, nil)
	if err != nil {
		return
	}
}

func RecordingPage(w http.ResponseWriter, r *http.Request) {
	parseFile, err := template.ParseFiles("../../ui/html/recording.html")
	if err != nil {
		return
	}
	err = parseFile.Execute(w, nil)
	if err != nil {
		return
	}
}
