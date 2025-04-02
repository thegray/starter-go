package example_test

import (
	"errors"
	"testing"

	"starter-go/internal/domain/example"
	exampleService "starter-go/internal/service/example"
)

type mockRepo struct {
	findByID func(int) (*example.Example, error)
	save     func(*example.Example) error
}

func (m *mockRepo) FindByID(id int) (*example.Example, error) {
	return m.findByID(id)
}

func (m *mockRepo) Save(e *example.Example) error {
	return m.save(e)
}

func TestGetExample_Success(t *testing.T) {
	mock := &mockRepo{
		findByID: func(id int) (*example.Example, error) {
			return &example.Example{ID: id, Description: "desc"}, nil
		},
	}

	service := exampleService.NewService(mock)
	e, err := service.GetExample(1)
	if err != nil || e.Description != "desc" {
		t.Fatalf("expected desc, got %v (err=%v)", e, err)
	}
}

func TestGetExample_NotFound(t *testing.T) {
	mock := &mockRepo{
		findByID: func(id int) (*example.Example, error) {
			return nil, errors.New("not found")
		},
	}

	service := exampleService.NewService(mock)
	_, err := service.GetExample(123)
	if err == nil {
		t.Fatal("expected error for missing example, got nil")
	}
}

func TestCreateExample(t *testing.T) {
	saved := false
	mock := &mockRepo{
		save: func(e *example.Example) error {
			saved = true
			return nil
		},
	}

	service := exampleService.NewService(mock)
	_, err := service.CreateExample("creating")
	if err != nil || !saved {
		t.Fatalf("expected save to be called, got err=%v", err)
	}
}
