package example

import "context"

type ExampleRepository interface {
	FindByID(ctx context.Context, id int) (*Example, error)
	FindAll(ctx context.Context) ([]*Example, error)
	Save(ctx context.Context, example *Example) error
}
