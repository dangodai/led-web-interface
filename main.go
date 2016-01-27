package main

import (
	//"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/dangodai/led-web-interface/colour"
	"github.com/dangodai/led-web-interface/pigpio"
)

var (
	myRGB *pigpio.RGB
)

func main() {
	myRGB = &pigpio.RGB{17, 22, 24}

	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/static/", staticFileHandler)
	http.ListenAndServe(":80", nil)
}

/*
	pageHandler func handle all the POST and GET requests
	for the front page of the website
*/
func pageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		c := parseFormColors(r)

		myRGB.ExecuteColour(c)
	}

	homeTemplate, _ := template.ParseFiles("templates/index.html")
	homeTemplate.Execute(w, nil)
}

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func parseFormColors(r *http.Request) colour.Colour {
	_, err := strconv.Atoi(r.PostForm["red"][0])
	if err != nil {
		return colour.Colour{"0", "0", "0"}
	}

	_, err = strconv.Atoi(r.PostForm["green"][0])
	if err != nil {
		return colour.Colour{"0", "0", "0"}
	}

	_, err = strconv.Atoi(r.PostForm["blue"][0])
	if err != nil {
		return colour.Colour{"0", "0", "0"}
	}

	return colour.Colour{r.PostForm["red"][0], r.PostForm["green"][0], r.PostForm["blue"][0]}
}