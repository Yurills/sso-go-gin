package policy

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func AddSecurityPolicyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Security-Policy",
			"default-src 'self'; "+
				"script-src 'self'; "+
				"style-src 'self'; "+
				"img-src 'self' data:; "+
				"font-src 'self'; "+
				"object-src 'none'; "+
				"base-uri 'self'; "+
				"form-action 'self'; "+
				"frame-ancestors 'none';")
		c.Next()
	}
}

func AddVaryOriginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			c.Header("Vary", "Origin")
		}
		c.Next()
	}
}

func AddSecurityHeadersMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "no-referrer-when-downgrade")
		c.Header("Strict-Transport-Security", "max-age=63072000; includeSubDomains; preload")

		c.Next()
	}
}

func RejectSuspiciousEndpointsMiddleware() gin.HandlerFunc { //reject BitKeeper endpoint, suggested from ZAP scanner
	return func(c *gin.Context) {
		if strings.HasPrefix(c.Request.URL.Path, "/BitKeeper") {
			c.AbortWithStatus(403)
			return
		}

		c.Next()
	}
}
