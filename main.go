package main

import(
	//"fmt"
	"html/template"
	"net/http"
	"strconv"
	"os/exec"
	"github.com/dangodai/led-web-interface/colour"
)

const (
	RED_PIN = "17"
	GREEN_PIN = "22"
	BLUE_PIN = "24"
)

func main() {
	http.HandleFunc("/", pageHandler)
	http.HandleFunc("/static/", staticFileHandler)
	http.ListenAndServe(":80", nil)
}

/*
	pageHandler func handle all the POST and GET requests
	for the front page of the website
*/
func pageHandler(w http.ResponseWriter, r *http.Request) {
	if(r.Method == "POST") {
		r.ParseForm()
		c := parseFormColors(r)

		execColour(c)
	}

	homeTemplate, _ := template.ParseFiles("index.html")
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

func execColour(c colour.Colour){
	cmd := "pigs"

	args := []string{"p", RED_PIN, c.Red}
	_ = exec.Command(cmd, args...).Run()
	args = []string{"p", GREEN_PIN, c.Green}
	_ = exec.Command(cmd, args...).Run()
	args = []string{"p", BLUE_PIN, c.Blue}
	_ = exec.Command(cmd, args...).Run()
}