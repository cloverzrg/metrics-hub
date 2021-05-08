package main

import (
	"github.com/cloverzrg/metrics-hub/api"
	"github.com/cloverzrg/metrics-hub/logger"
)

func main() {
	err := api.Serve()
	if err != nil {
		logger.Panic(err)
	}
}
