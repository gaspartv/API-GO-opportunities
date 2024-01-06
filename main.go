package main

import (
	"github.com/gaspartv/API-GO-opportunities/config"
	"github.com/gaspartv/API-GO-opportunities/router"
)

var (
	logger config.Logger
)

func main() {
    logger = *config.GetLogger("main")

	// Initialize Configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
        return
	}

	// Initialize Router
	router.Initialize()
}
