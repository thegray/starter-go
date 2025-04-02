package example

import (
	"encoding/json"
	"net/http"
	"strconv"

	svc "starter-go/internal/service/example"
)

type Handler struct {
	service *svc.ExampleService
}

func NewHandler(service *svc.ExampleService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetExample(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	e, err := h.service.GetExample(id)
	if err != nil {
		http.Error(w, "Example not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(FromDomain(e))
}

func (h *Handler) CreateExample(w http.ResponseWriter, r *http.Request) {
	var req CreateExampleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid body", http.StatusBadRequest)
		return
	}

	e, err := h.service.CreateExample(req.Description)
	if err != nil {
		http.Error(w, "Create failed", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(FromDomain(e))
}
