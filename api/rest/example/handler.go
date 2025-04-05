package example

import (
	"net/http"
	"strconv"

	"starter-go/internal/domain/example"
	"starter-go/internal/pkg/errors"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service example.ExampleService
}

func NewHandler(service example.ExampleService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetExample(c *gin.Context) {
	idStr := c.Param("example_id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		svcErr := errors.ErrInvalidFieldFormat("example_id", err)
		c.Error(svcErr)
		c.Abort()
		return
	}

	e, err := h.service.GetExample(c.Request.Context(), id)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, FromDomain(e))
}

func (h *Handler) CreateExample(c *gin.Context) {
	var req CreateExampleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		svcErr := errors.ErrInvalidRequest(err)
		c.Error(svcErr)
		c.Abort()
		return
	}

	e, err := h.service.CreateExample(c.Request.Context(), req.Description)
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, FromDomain(e))
}
