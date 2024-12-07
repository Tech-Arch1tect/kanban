package middleware

import (
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

var rateLimitCache = sync.Map{}

type rateLimiter struct {
	count     int
	expiresAt time.Time
}

func RateLimit(limit int, window time.Duration) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		now := time.Now()

		value, _ := rateLimitCache.LoadOrStore(ip, &rateLimiter{0, now.Add(window)})
		rl := value.(*rateLimiter)

		if now.After(rl.expiresAt) {
			rl.count = 0
			rl.expiresAt = now.Add(window)
		}

		if rl.count >= limit {
			c.AbortWithStatusJSON(429, gin.H{"error": "Too Many Requests"})
			return
		}

		rl.count++
		rateLimitCache.Store(ip, rl)
		c.Next()
	}
}
