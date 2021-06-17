package consul

import (
	"fmt"

	"github.com/cloverzrg/metrics-hub/config"
	"github.com/cloverzrg/metrics-hub/logger"
	"github.com/cloverzrg/metrics-hub/util"
	"github.com/hashicorp/consul/api"
)

func init() {
	err := Register("metrics-hub")
	if err != nil {
		logger.Error("Register metrics-hub error", err)
	}
}

func Register(serviceName string) (err error) {
	address, err := util.GetLocalHostAddress()
	if err != nil {
		logger.Error(err)
		return err
	}
	serviceId := fmt.Sprintf("%s-%s", serviceName, address)
	checkId := fmt.Sprintf("%s-check", serviceId)

	check := api.AgentServiceCheck{
		CheckID:                        checkId,
		Interval:                       "15s",
		HTTP:                           fmt.Sprintf("%s/healthy", config.Http.ExternalUrl),
		DeregisterCriticalServiceAfter: "45s",
	}
	server := api.AgentServiceRegistration{
		ID:      serviceId,
		Name:    serviceName,
		Port:    config.Http.Port,
		Address: address,
		Check:   &check,
	}
	logger.Infof("register service %s to consul", serviceId)
	err = Client.Agent().ServiceRegister(&server)
	return err
}

func JobRegister(job string, application string, groupingKey map[string]string) (err error) {
	groupingKey["metrics_path"] = fmt.Sprintf("/metrics/job/%s", job)
	address, err := util.GetLocalHostAddress()
	if err != nil {
		logger.Error(err)
		return err
	}

	serviceId := job
	checkId := fmt.Sprintf("%s-check", serviceId)

	check := api.AgentServiceCheck{
		CheckID:                        checkId,
		Interval:                       "15s",
		HTTP:                           fmt.Sprintf("%s/metrics/job/%s/healthy", config.Http.ExternalUrl, job),
		DeregisterCriticalServiceAfter: "60s",
	}

	server := api.AgentServiceRegistration{
		ID:      serviceId,
		Name:    application,
		Port:    config.Http.Port,
		Address: address,
		Meta:    groupingKey,
		Check:   &check,
		Tags:    []string{"prometheus-metrics"},
	}
	logger.Infof("register service %s to consul", serviceId)
	err = Client.Agent().ServiceRegister(&server)
	return err
}
