package example

import (
	"net/http"
	"strconv"

	"starter-go/internal/pkg/errors"
	svc "starter-go/internal/service/example"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *svc.ExampleService
}

func NewHandler(service *svc.ExampleService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetExample(c *gin.Context) {
	idStr := c.Param("example_id") // from route :example_id
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		svcErr := errors.ErrInvalidFieldFormat("example_id", err)
		c.Error(svcErr)
		c.Abort()
		return
	}

	e, err := h.service.GetExample(id)
	if err != nil {
		svcErr := errors.ErrNotFound("example", err)
		c.Error(svcErr)
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

	// Validate required fields
	if req.Description == "" {
		svcErr := errors.ErrMissingMandatoryField("description", nil)
		c.Error(svcErr)
		c.Abort()
		return
	}

	e, err := h.service.CreateExample(req.Description)
	if err != nil {
		svcErr := errors.New("ERR_CREATE_FAILED", "Failed to create example", err)
		c.Error(svcErr)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, FromDomain(e))
}
