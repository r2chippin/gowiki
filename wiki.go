package main

import (
	"encoding/json"
	"fmt"
	"gowiki/controller"
	model "gowiki/models"
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
	configData, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(configData, &config)
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is starting...")
	fmt.Println("Target Directory:", config.TargetDirectory)
	fmt.Println("Log Level:", config.LogLevel)
	fmt.Println("Max Retries:", config.MaxRetries)
	fmt.Println("Timeout Seconds:", config.TimeoutSeconds)

	p1 := &model.Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
	err = p1.Save(config.TargetDirectory)
	if err != nil {
		panic(err)
	}
	p2, _ := model.LoadPage("TestPage", config.TargetDirectory)
	fmt.Println(string(p2.Body))

	http.HandleFunc("/view/", func(w http.ResponseWriter, r *http.Request) {
		controller.ViewHandler(w, r, config.TargetDirectory)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
