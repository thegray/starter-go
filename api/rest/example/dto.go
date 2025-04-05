package example

import "starter-go/internal/domain/example"

type CreateExampleRequest struct {
	Description string `json:"description"`
}

type ExampleResponse struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
}

func FromDomain(e *example.Example) ExampleResponse {
	return ExampleResponse{
		ID:          e.ID,
		Description: e.Description,
	}
}
