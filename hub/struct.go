package hub

import "time"

type JobMetrics struct {
	Data        []byte
	PushTime    time.Time
	JobName     string
	Application string
	GroupingKey map[string]string
}

// IsValid 检查该 metrics 是否有效（过期则无效）
func (jobMetrics *JobMetrics) IsValid() bool {
	return jobMetrics.PushTime.After(time.Now().Add(-60 * time.Second))
}

func (jobMetrics *JobMetrics) Update(data []byte) {
	jobMetrics.PushTime = time.Now()
	jobMetrics.Data = data
}
