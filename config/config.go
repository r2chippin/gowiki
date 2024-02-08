package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config contains some cfg will be used
type Config struct {
	Name   string `json:"name"`
	Port   string `json:"port"`
	DBInfo string `json:"db"`
}

func (conf *Config) InitConfig() {
	// Read config.json to get config info
	configData, err := os.ReadFile("config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configData, conf)
	if err != nil {
		panic(err)
	}
	// Show all info
	fmt.Println(conf.Name + " is starting...")
	fmt.Println("Port:", conf.Port)
}
