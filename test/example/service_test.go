package example_test

import (
	"context"
	"errors"
	"testing"

	"starter-go/internal/domain/example"
	exampleService "starter-go/internal/service/example"
)

type mockRepo struct {
	findByID func(context.Context, int) (*example.Example, error)
	save     func(context.Context, *example.Example) error
}

func (m *mockRepo) FindByID(ctx context.Context, id int) (*example.Example, error) {
	return m.findByID(ctx, id)
}

func (m *mockRepo) Save(ctx context.Context, e *example.Example) error {
	return m.save(ctx, e)
}

func TestGetExample_Success(t *testing.T) {
	mock := &mockRepo{
		findByID: func(ctx context.Context, id int) (*example.Example, error) {
			return &example.Example{ID: id, Description: "desc"}, nil
		},
	}

	service := exampleService.NewService(mock)
	ctx := context.Background()
	e, err := service.GetExample(ctx, 1)
	if err != nil || e.Description != "desc" {
		t.Fatalf("expected desc, got %v (err=%v)", e, err)
	}
}

func TestGetExample_NotFound(t *testing.T) {
	mock := &mockRepo{
		findByID: func(ctx context.Context, id int) (*example.Example, error) {
			return nil, errors.New("not found")
		},
	}

	service := exampleService.NewService(mock)
	ctx := context.Background()
	_, err := service.GetExample(ctx, 123)
	if err == nil {
		t.Fatal("expected error for missing example, got nil")
	}
}

func TestCreateExample(t *testing.T) {
	saved := false
	mock := &mockRepo{
		save: func(ctx context.Context, e *example.Example) error {
			saved = true
			e.ID = 1 // Simulate ID assignment by database
			return nil
		},
	}

	service := exampleService.NewService(mock)
	ctx := context.Background()
	example, err := service.CreateExample(ctx, "creating")
	if err != nil || !saved {
		t.Fatalf("expected save to be called, got err=%v", err)
	}
	if example.ID != 1 {
		t.Fatalf("expected ID to be set, got %d", example.ID)
	}
}
