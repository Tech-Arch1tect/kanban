package middleware

import (
	"server/config"

	"github.com/gin-gonic/gin"
)

func Cors(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		AllowedOrigin := cfg.AllowOrigin
		c.Writer.Header().Set("Access-Control-Allow-Origin", AllowedOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}
