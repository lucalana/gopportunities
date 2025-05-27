package main

import (
	"github.com/lucalana/gopportunities/config"
	"github.com/lucalana/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}
	// Initialize Routes
	router.Initialize()
}
