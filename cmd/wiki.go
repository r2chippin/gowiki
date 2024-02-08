package main

import (
	"gowiki/config"
	"gowiki/routes"
)

var cfg config.Config

func main() {
	// Read Config
	cfg.InitConfig()

	// Start Gin server
	e := routes.NewRouter(cfg)
	_ = e.Run(cfg.Port)
}
