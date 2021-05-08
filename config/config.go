package config

import (
	"fmt"
	"github.com/cloverzrg/metrics-hub/logger"
	"github.com/cloverzrg/metrics-hub/util"
	"os"
	"strconv"
)

var Http = struct {
	Listen      string
	Port        int
	ExternalUrl string
}{}

var Consul = struct {
	Address string
	Token   string
}{}

func init() {
	var err error
	Http.Port = 9101
	if os.Getenv("http-port") != "" {
		Http.Port, err = strconv.Atoi(os.Getenv("http-port"))
		if err != nil {
			logger.Panic(err)
		}
	}
	Http.Listen = fmt.Sprintf(":%d", Http.Port)
	if os.Getenv("http-listen") != "" {
		Http.Listen = os.Getenv("http-listen")
	}
	hostIP, err := util.GetLocalHostAddress()
	if err != nil {
		logger.Panic(err)
	}
	Http.ExternalUrl = fmt.Sprintf("http://%s:%d", hostIP, Http.Port)
	if os.Getenv("http-external-url") != "" {
		Http.ExternalUrl = os.Getenv("http-external-url")
	}
	Consul.Address = "127.0.0.1:8085"
	if os.Getenv("consul-address") != "" {
		Consul.Address = os.Getenv("consul-address")
	}
	if os.Getenv("consul-token") != "" {
		Consul.Token = os.Getenv("consul-token")
	}
}
