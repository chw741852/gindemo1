package main

import (
	"test/router"
	"test/internal/logger"
)

func main() {
	r := router.Init()
	logger.Init()
	r.Run()
}
