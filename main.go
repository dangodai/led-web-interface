package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/dangodai/led-web-interface/colour"
	"github.com/dangodai/led-web-interface/pigpio"
)

var (
	myRGB *pigpio.RGB
	myWhite *pigpio.FixedColour
)

func main() {
	myRGB = &pigpio.RGB{17, 22, 24}
	myWhite = &pigpio.FixedColour{pigpio.Brightness{27}}	

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
		fmt.Println(r.Form,"\n")
		
		c, err := parseFormColors(r)
		if err == nil {myRGB.ExecuteColour(c)}
		b, err := parseFormSingleLight(r)
		if err == nil {myWhite.SetBrightness(b)}	
	}

	homeTemplate, _ := template.ParseFiles("templates/index.html")
	homeTemplate.Execute(w, nil)
}

func staticFileHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, r.URL.Path[1:])
}

func parseFormSingleLight(r *http.Request) (int, error) {
	
	l, err := strconv.Atoi(r.PostForm["white"][0])
		
	return l, err
}

func parseFormColors(r *http.Request) (colour.Colour, error) {

	c := colour.Colour{0, 0, 0}

	c.Red, _ = strconv.Atoi(r.PostForm["red"][0])
	_, err1 := strconv.Atoi(r.PostForm["red"][0])

	c.Green, _ = strconv.Atoi(r.PostForm["green"][0])
	_, err2 := strconv.Atoi(r.PostForm["green"][0])

	c.Blue, _ = strconv.Atoi(r.PostForm["blue"][0])
	_, err3 := strconv.Atoi(r.PostForm["blue"][0])

	if (err1 != nil && err2 != nil && err3 != nil) {
		return c, err3
	} else {
		return c, nil
	}
}
