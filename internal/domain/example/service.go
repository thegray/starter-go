package example

import "context"

type ExampleService interface {
	GetExample(ctx context.Context, id int) (*Example, error)
	GetAllExamples(ctx context.Context) ([]*Example, error)
	CreateExample(ctx context.Context, description string) (*Example, error)
}
