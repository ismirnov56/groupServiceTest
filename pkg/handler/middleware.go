package handler

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

const xRequestID = "x-request-id"

func JSONLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		end := time.Now()

		duration := end.Sub(start)

		status := c.Writer.Status()

		logger := log.WithFields(log.Fields{
			"client_ip":  c.ClientIP(),
			"duration":   duration,
			"method":     c.Request.Method,
			"path":       c.Request.RequestURI,
			"status":     status,
			"request_id": c.GetHeader(xRequestID),
		})

		if c.Writer.Status() >= 500 {
			logger.Error(c.Errors.String())
		} else {
			logger.Info("")
		}
	}
}
