package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/cloverzrg/metrics-hub/logger"
	"github.com/cloverzrg/metrics-hub/util"
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
	Http.Port = 9091
	if os.Getenv("HTTP_PORT") != "" {
		Http.Port, err = strconv.Atoi(os.Getenv("HTTP_PORT"))
		if err != nil {
			logger.Panic(err)
		}
	}
	Http.Listen = fmt.Sprintf(":%d", Http.Port)
	if os.Getenv("HTTP_LISTEN") != "" {
		Http.Listen = os.Getenv("HTTP_LISTEN")
	}
	hostIP, err := util.GetLocalHostAddress()
	if err != nil {
		logger.Panic(err)
	}
	Http.ExternalUrl = fmt.Sprintf("http://%s:%d", hostIP, Http.Port)
	if os.Getenv("HTTP_EXTERNAL_URL") != "" {
		Http.ExternalUrl = os.Getenv("HTTP_EXTERNAL_URL")
	}
	Consul.Address = "127.0.0.1:8085"
	if os.Getenv("CONSUL_ADDRESS") != "" {
		Consul.Address = os.Getenv("CONSUL_ADDRESS")
	}
	if os.Getenv("CONSUL_TOKEN") != "" {
		Consul.Token = os.Getenv("CONSUL_TOKEN")
	}
}
