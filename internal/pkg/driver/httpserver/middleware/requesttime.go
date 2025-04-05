package middleware

import (
	"fmt"
	"starter-go/internal/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestTimer() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		duration := time.Since(startTime)

		ctx := GetContext(c)
		logger.InfoCtx(ctx,
			fmt.Sprintf("[Request finished] method:%s, path:%s, status:%d, elapsed:%d",
				c.Request.URL.Path, c.Request.Method, c.Writer.Status(), duration),
		)
	}
}
