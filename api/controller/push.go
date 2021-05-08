package controller

import (
	"github.com/cloverzrg/metrics-hub/hub"
	"github.com/cloverzrg/metrics-hub/logger"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strings"
)

func parseGroupingKey(str string) map[string]string {
	m := make(map[string]string)
	groupKV := strings.Trim(str, "/")
	split := strings.Split(groupKV, "/")
	if len(split) > 1 && len(split)%2 == 0 {
		for i := 0; i < len(split); i += 2 {
			m[split[i]] = split[i+1]
		}
	}
	return m
}

func Push(c *gin.Context) {
	job := c.Param("job")
	groupKV := parseGroupingKey(c.Param("groupKV"))

	bytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		logger.Error(err)
		c.String(500, err.Error())
		return
	}
	err = hub.AddJobMetrics(job, bytes, groupKV)
	if err != nil {
		logger.Error(err)
		c.String(500, err.Error())
		return
	}

	logger.Infof("job:%s, group:%v", job, groupKV)
	c.String(200, "ok")
}
