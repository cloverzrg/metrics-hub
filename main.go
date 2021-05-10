package main

import (
	"fmt"
	"github.com/cloverzrg/metrics-hub/api"
	"github.com/cloverzrg/metrics-hub/logger"
	"runtime"
)

var (
	BuildTime string
	GitMessage   string
)


func main() {
	err := api.Serve()
	if err != nil {
		logger.Panic(err)
	}
}

func init() {
	buildInfo := fmt.Sprintf("BuildTime: %s\nGoVersion: %s\nGitHead: %s\n", BuildTime, runtime.Version(), GitMessage)
	fmt.Println(buildInfo)
}