package main

import (
	"test/router"
	"test/internal/logger"
	"test/internal/config"
)

func main() {
	r := router.Init()
	logger.Init()
	config.Init("config/config.local.yml")
	r.Run()
}
