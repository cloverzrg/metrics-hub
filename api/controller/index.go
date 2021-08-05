package controller

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.String(200, "ok, response at %v", time.Now())
}
