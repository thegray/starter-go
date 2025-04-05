package middleware

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"starter-go/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

func CustomRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				err, ok := err.(error)
				ctx := GetContext(c)

				if !ok {
					err = fmt.Errorf("%+v", err)
				}
				logger.ErrorCtx(ctx, fmt.Sprintf("[HTTP:Recover] panic %s", err.Error()),
					"stacktrace", string(debug.Stack()),
				)

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
