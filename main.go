package main

import (
	"github.com/fabiomzs/gopportunities/config"
	"github.com/fabiomzs/gopportunities/router"
)

var (
	logger config.Logger
)

func main() {
	logger = *config.GetLogger("main")

	//Init configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %v", err)
		return
	}

	// Init router
	router.Initialize()
}
