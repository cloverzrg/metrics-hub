package api

import (
	"github.com/cloverzrg/metrics-hub/api/controller"
	"github.com/cloverzrg/metrics-hub/config"
	"github.com/cloverzrg/metrics-hub/logger"
	"github.com/gin-gonic/gin"
)

func SetRoute(r *gin.Engine ) {
	g := r.Group("/metrics")

	g.GET("", controller.Index)
	// push
	g.POST("/:job", controller.Push)
	g.POST("/:job/*groupKV", controller.Push)

	// get metrics
	g.GET("/:job", controller.JobMetrics)
	g.GET("/:job/healthy", controller.JobMetricsHealth)
}

func Serve() error {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.LoggerWithWriter(logger.Entry.Writer()))
	SetRoute(r)
	return r.Run(config.Http.Listen)
}