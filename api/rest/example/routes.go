package example

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, h *Handler) {
	ex := r.Group("/api/v1/examples")
	exampleRoutes(ex, h)
}

func exampleRoutes(r *gin.RouterGroup, h *Handler) {
	r.GET("/", h.GetAllExamples)
	r.GET("/:example_id", h.GetExample)
	r.POST("/", h.CreateExample)
}
