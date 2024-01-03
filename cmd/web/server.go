package main

import (
	"fmt"
	"log"
	"net/http"
)

const server_port = ":8007"

func main() {
	http.HandleFunc("/main", MainPage)
	http.HandleFunc("/", CalibrationIntroPage)
	http.HandleFunc("/calibration", CalibrationPage)
	http.HandleFunc("/content", ContentPage)
	http.HandleFunc("/cinema", CinemaPage)
	http.HandleFunc("/recoring", RecordingPage)
	http.Handle("/ui/html/", http.StripPrefix("/ui/html", http.FileServer(http.Dir("../../ui/html"))))
	http.Handle("/ui/css/", http.StripPrefix("/ui/css", http.FileServer(http.Dir("../../ui/css"))))
	http.Handle("/ui/pictures/", http.StripPrefix("/ui/pictures", http.FileServer(http.Dir("../../ui/pictures"))))
	http.Handle("/ui/content/", http.StripPrefix("/ui/content", http.FileServer(http.Dir("../../ui/content"))))
	http.Handle("/ui/content_pictures/", http.StripPrefix("/ui/content_pictures", http.FileServer(http.Dir("../../ui/content_pictures"))))
	fmt.Println("Running server on http://localhost" + server_port)
	if err := http.ListenAndServe(server_port, nil); err != nil {
		log.Println(err)
		return
	}
}
