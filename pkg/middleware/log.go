package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Injeta o logger no contexto
		c.Set("logger", logger)
		c.Next()
	}
}
