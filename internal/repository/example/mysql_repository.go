package example

import (
	"starter-go/internal/domain/example"

	"gorm.io/gorm"
)

type ExampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) *ExampleRepository {
	return &ExampleRepository{db: db}
}

func (r *ExampleRepository) FindByID(id int) (*example.Example, error) {
	var model ExampleModel
	if err := r.db.First(&model, id).Error; err != nil {
		return nil, err
	}
	return model.ToDomain(), nil
}

func (r *ExampleRepository) Save(e *example.Example) error {
	return r.db.Create(FromDomain(e)).Error
}
