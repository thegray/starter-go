package example

import (
	"context"
	entity "starter-go/internal/domain/example"
	"starter-go/internal/pkg/errors"
)

// compile-time check to ensure ExampleService implements the example.ExampleService interface
// no effect on the runtime
var _ entity.ExampleService = (*ExampleService)(nil)

type ExampleService struct {
	repo entity.ExampleRepository
}

func NewService(repo entity.ExampleRepository) *ExampleService {
	return &ExampleService{repo: repo}
}

func (s *ExampleService) GetExample(ctx context.Context, id int) (*entity.Example, error) {
	example, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, errors.ErrNotFound("example", err)
	}
	return example, nil
}

func (s *ExampleService) CreateExample(ctx context.Context, desc string) (*entity.Example, error) {
	e := &entity.Example{Description: desc}
	err := s.repo.Save(ctx, e)
	if err != nil {
		return nil, errors.New("EXAMPLE_CREATE_FAILED", "Failed to create example", err)
	}

	return e, nil
}
