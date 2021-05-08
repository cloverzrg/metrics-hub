package hub

import "time"

var metricsHub = make(map[string]*JobMetrics)

func AddJobMetrics(jobName string, data []byte, groupingKey map[string]string) (err error) {
	metricsHub[jobName] = &JobMetrics{
		Data:        data,
		PushTime:    time.Now(),
		GroupingKey: groupingKey,
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
