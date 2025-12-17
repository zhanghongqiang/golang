package middleware

import (
	"net/http"

	"task4/utils"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				utils.Logger.WithFields(logrus.Fields{
					"error":   err,
					"path":    c.Request.URL.Path,
					"method":  c.Request.Method,
					"headers": c.Request.Header,
				}).Error("Panic recovered")

				c.JSON(http.StatusInternalServerError, gin.H{
					"success": false,
					"error":   "Internal server error",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
