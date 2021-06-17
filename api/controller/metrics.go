package controller

import (
	"bytes"

	"github.com/cloverzrg/metrics-hub/hub"
	"github.com/cloverzrg/metrics-hub/logger"
	"github.com/gin-gonic/gin"
)

func JobMetrics(c *gin.Context) {
	job := c.Param("job")
	metrics, exist, err := hub.GetJobMetrics(job)
	if err != nil {
		logger.Error(err)
		c.String(500, err.Error())
		return
	}
	if !exist {
		c.String(404, "# Not Found")
		return
	}
	buffer := bytes.NewBuffer(metrics.Data)
	n, err := buffer.WriteTo(c.Writer)
	if err != nil {
		logger.Error(err)
		c.String(500, err.Error())
		return
	}
	logger.Infof("send %s metrics, write %d bytes", job, n)
}

func JobMetricsHealth(c *gin.Context) {
	job := c.Param("job")
	isHealthy, data := hub.IsHealthy(job)
	if !isHealthy {
		c.String(404, "404")
		return
	}
	c.String(200, "last push at %v", data.PushTime)
}
