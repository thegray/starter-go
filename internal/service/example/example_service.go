package user

import (
	"errors"
	"starter-go/internal/domain/example"
)

type ExampleService struct {
	repo example.ExampleRepository
}

func NewService(repo example.ExampleRepository) *ExampleService {
	return &ExampleService{repo: repo}
}

func (s *ExampleService) GetExample(id int) (*example.Example, error) {
	if id <= 0 {
		return nil, errors.New("invalid example ID")
	}
	return s.repo.FindByID(id)
}

func (s *ExampleService) CreateExample(desc string) (*example.Example, error) {
	e := &example.Example{Description: desc}
	err := s.repo.Save(e)
	return e, err
}
