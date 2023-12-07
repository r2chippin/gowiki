package main

import (
	"encoding/json"
	"fmt"
	"gowiki/controller"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
)

// Config contains some cfg will be used
type Config struct {
	Name            string `json:"name"`
	Address         string `json:"address"`
	TargetDirectory string `json:"targetDirectory"`
	SafePath        string `json:"safePath"`
	LogLevel        string `json:"logLevel"`
	MaxRetries      int    `json:"maxRetries"`
	TimeoutSeconds  int    `json:"timeoutSeconds"`
}

var config Config

func main() {
	// Read config.json to get config info
	configData, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}
	// Show all info
	fmt.Println(config.Name + " is starting...")
	fmt.Println("Address:", config.Address)
	fmt.Println("Target Directory:", config.TargetDirectory)
	fmt.Println("Safe Path:", config.SafePath)
	fmt.Println("Log Level:", config.LogLevel)
	fmt.Println("Max Retries:", config.MaxRetries)
	fmt.Println("Timeout Seconds:", config.TimeoutSeconds)

	var templates = template.Must(template.ParseFiles("./tmpl/edit.html", "./tmpl/view.html"))
	var validPath = regexp.MustCompile(config.SafePath)

	http.HandleFunc("/view/", func(w http.ResponseWriter, r *http.Request) {
		controller.ViewHandler(w, r, config.TargetDirectory, templates, validPath)
	})
	http.HandleFunc("/edit/", func(w http.ResponseWriter, r *http.Request) {
		controller.EditHandler(w, r, config.TargetDirectory, templates, validPath)
	})
	http.HandleFunc("/save/", func(w http.ResponseWriter, r *http.Request) {
		controller.SaveHandler(w, r, config.TargetDirectory, validPath)
	})

	log.Fatal(http.ListenAndServe(config.Address, nil))
}
