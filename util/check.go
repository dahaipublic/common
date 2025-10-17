package util

import (
	"github.com/gin-gonic/gin"
	"time"
)

type HealthResponse struct {
	Status string `json:"status"`
	Uptime string `json:"uptime"`
}

func CheckHealth(c *gin.Context) {
	uptime := time.Since(time.Now()).String()
	rsp := HealthResponse{
		Status: "ok",
		Uptime: uptime,
	}
	RespDataResult(c, rsp)
	return
}
