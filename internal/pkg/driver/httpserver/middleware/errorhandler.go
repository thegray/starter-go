package middleware

import (
	"fmt"
	"net/http"
	"starter-go/internal/pkg/env"
	"starter-go/internal/pkg/errors"
	"starter-go/internal/pkg/logger"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) == 0 {
			return
		}

		err := c.Errors[0].Err
		var (
			statusCode = http.StatusInternalServerError
			resp       = gin.H{"message": "Internal Server Error"}
		)

		switch e := err.(type) {
		case errors.ServiceError:
			statusCode = errors.HTTPStatusFromCode(e.Code())
			resp = gin.H{
				"code":    e.Code(),
				"message": e.Message(),
			}

			if env.IsDevelopment() {
				resp["stack"] = e.Stacktrace()
			}

			logger.ErrorCtx(GetContext(c),
				fmt.Sprintf("[ErrHandler] %s", e.Message()),
				"code", e.Code(),
				"error", e.Error(),
			)

			if env.IsDevelopment() {
				logger.ErrorCtx(GetContext(c),
					"[ErrHandler] stacktrace",
					"stack", e.Stacktrace(),
				)
			}

		default:
			logger.ErrorCtx(GetContext(c),
				"[ErrHandler] Unknown error",
				"error", err.Error(),
			)
		}

		if !c.Writer.Written() {
			c.JSON(statusCode, resp)
		}
	}
}
