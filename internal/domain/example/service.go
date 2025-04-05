package example

import "context"

type ExampleService interface {
	GetExample(ctx context.Context, id int) (*Example, error)
	CreateExample(ctx context.Context, description string) (*Example, error)
}
