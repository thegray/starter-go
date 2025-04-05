package middleware

import (
	"context"
	"net/http"

	"starter-go/internal/pkg/logger/contextid"

	"github.com/gin-gonic/gin"
)

const (
	contextKey              = "int-context"
	responseHeaderContextID = "Context-ID"
	requestHeaderRequestID  = "X-Request-Id"
)

func ContextMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := newContext(c.Request)
		c.Set(contextKey, ctx)
		c.Header(responseHeaderContextID, contextid.Value(ctx))
		c.Next()
	}
}

func GetContext(c *gin.Context) context.Context {
	i, exists := c.Get(contextKey)
	if !exists {
		ctx := newContext(c.Request)
		c.Set(contextKey, ctx)
		return ctx
	}

	switch v := i.(type) {
	case context.Context:
		return v
	default:
		ctx := newContext(c.Request)
		c.Set(contextKey, ctx)
		return ctx
	}
}

func newContext(req *http.Request) context.Context {
	if requestID := req.Header.Get(requestHeaderRequestID); requestID != "" {
		return contextid.NewWithValue(context.Background(), requestID)
	} else {
		return contextid.New(context.Background())
	}
}
