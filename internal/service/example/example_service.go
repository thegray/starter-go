package user

import (
	"errors"
	"fmt"
	"starter-go/internal/domain/example"
)

type ExampleService struct {
	repo example.ExampleRepository
}

func NewService(repo example.ExampleRepository) *ExampleService {
	return &ExampleService{repo: repo}
}

func (s *ExampleService) GetExample(id int) (*example.Example, error) {
	fmt.Println("PBPBPB-get example service")
	if id <= 0 {
		return nil, errors.New("invalid example ID")
	}
	// return &example.Example{ID: id, Description: "test"}, nil
	return s.repo.FindByID(id)
}

func (s *ExampleService) CreateExample(desc string) (*example.Example, error) {
	e := &example.Example{Description: desc}
	err := s.repo.Save(e)
	return e, err
}
