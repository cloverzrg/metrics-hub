package controller

import (
	"github.com/gin-gonic/gin"
	"time"
)

func Index(c *gin.Context) {
	c.String(200, "ok, response at %v", time.Now())
}
