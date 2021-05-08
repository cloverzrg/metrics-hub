package consul

import (
	"github.com/cloverzrg/metrics-hub/config"
	"github.com/cloverzrg/metrics-hub/logger"
	"github.com/hashicorp/consul/api"
)

var Client *api.Client

func init() {
	c := &api.Config{
		Address: config.Consul.Address,
		Scheme:  "http",
		Token:   config.Consul.Token,
	}
	var err error
	Client, err = api.NewClient(c)
	if err != nil {
		logger.Error(err)
	}
}
