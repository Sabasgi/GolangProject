package middleware

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type RateLimiter struct {
	Tokens         float64
	LastRefillTime time.Time
	RatePerSecond  float64
	BurstSize      float64
	mu             sync.Mutex
}

func (rl *RateLimiter) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(rl.LastRefillTime).Seconds()

	// Refill tokens
	rl.Tokens += elapsed * rl.RatePerSecond
	if rl.Tokens > rl.BurstSize {
		rl.Tokens = rl.BurstSize
	}
	rl.LastRefillTime = now

	// Allow request if token available
	if rl.Tokens >= 1 {
		rl.Tokens -= 1
		return true
	}
	return false
}

type Client struct {
	Limiter  *RateLimiter
	LastSeen time.Time
}

// Global storage for all IPs
var (
	clients   = make(map[string]*Client)
	clientsMu sync.Mutex
)

// RateLimitMiddleware applies rate limiting based on IP
func RateLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		clientsMu.Lock()
		client, exists := clients[ip]
		if !exists {
			client = &Client{
				Limiter: &RateLimiter{
					Tokens:         5,
					LastRefillTime: time.Now(),
					RatePerSecond:  1,
					BurstSize:      5,
				},
				LastSeen: time.Now(),
			}
			clients[ip] = client
		}
		client.LastSeen = time.Now()
		clientsMu.Unlock()

		if !client.Limiter.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}
		c.Next()
	}
}
