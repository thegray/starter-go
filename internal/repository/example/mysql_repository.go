package example

import (
	"context"
	"starter-go/internal/domain/example"

	"gorm.io/gorm"
)

type ExampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) *ExampleRepository {
	return &ExampleRepository{db: db}
}

func (r *ExampleRepository) FindByID(ctx context.Context, id int) (*example.Example, error) {
	var model ExampleModel
	if err := r.db.WithContext(ctx).First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

func (r *ExampleRepository) Save(ctx context.Context, e *example.Example) error {
	return r.db.WithContext(ctx).Create(FromDomain(e)).Error
}
