package api

import (
	"github.com/cloverzrg/metrics-hub/api/controller"
	"github.com/cloverzrg/metrics-hub/config"
	"github.com/cloverzrg/metrics-hub/logger"
	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine ) {
	g := r.Group("/metrics")

	// push
	g.POST("/job/:job", controller.Push)
	g.POST("/job/:job/*groupKV", controller.Push)

	// get metrics
	g.GET("/job/:job", controller.JobMetrics)
	g.GET("/job/:job/healthy", controller.JobMetricsHealth)

	r.GET("/healthy", controller.Index)
}

func Serve() error {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithWriter(logger.Entry.Writer()))
	SetRoute(r)
	return r.Run(config.Http.Listen)
}