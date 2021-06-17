package hub

import (
	"github.com/cloverzrg/metrics-hub/consul"
	"github.com/cloverzrg/metrics-hub/logger"
	"time"
)

// 有并发问题
var metricsHub = make(map[string]*JobMetrics)

func AddJobMetrics(jobName string, data []byte, groupingKey map[string]string) (err error) {
	var application = groupingKey["application"]
	if application == "" {
		application = "novadax"
	}
	// 注册服务
	if metricsHub[jobName] == nil {
		err := consul.JobRegister(jobName, application, groupingKey)
		if err != nil {
			logger.Errorf("register job metrics error: %s", err)
		}
		metricsHub[jobName] = &JobMetrics{
			Data:        data,
			PushTime:    time.Now(),
			JobName:     jobName,
			Application: application,
			GroupingKey: groupingKey,
		}
	} else {
		metricsHub[jobName].Update(data)
	}

	return err
}

func IsHealthy(jobName string) (isHealthy bool, data *JobMetrics) {
	metrics, exist, err := GetJobMetrics(jobName)
	if err != nil {
		return false, nil
	}
	if !exist {
		return false, nil
	}
	return true, metrics
}

func GetJobMetrics(jobName string) (data *JobMetrics, exist bool, err error) {
	var has bool
	var jobMetrics *JobMetrics
	if jobMetrics, has = metricsHub[jobName]; has {
		if jobMetrics.IsValid() {
			return jobMetrics, true, err
		} else {
			delete(metricsHub, jobName)
		}
	}
	return nil, false, err
}
