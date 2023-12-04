package main

import (
	"encoding/json"
	"fmt"
	"gowiki/controller"
	"log"
	"net/http"
	"os"
)

// Config contains some cfg will be used
type Config struct {
	TargetDirectory string `json:"targetDirectory"`
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
	fmt.Println("Server is starting...")
	fmt.Println("Target Directory:", config.TargetDirectory)
	fmt.Println("Log Level:", config.LogLevel)
	fmt.Println("Max Retries:", config.MaxRetries)
	fmt.Println("Timeout Seconds:", config.TimeoutSeconds)

	http.HandleFunc("/view/", func(w http.ResponseWriter, r *http.Request) {
		controller.ViewHandler(w, r, config.TargetDirectory)
	})
	http.HandleFunc("/edit/", func(w http.ResponseWriter, r *http.Request) {
		controller.EditHandler(w, r, config.TargetDirectory)
	})
	/*
		http.HandleFunc("/save/", func(w http.ResponseWriter, r *http.Request) {
			controller.SaveHandler(w, r, config.TargetDirectory)
		})
	*/
	log.Fatal(http.ListenAndServe(":8080", nil))
}
