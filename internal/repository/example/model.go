package example

import "starter-go/internal/domain/example"

type ExampleModel struct {
	ID          uint `gorm:"primaryKey"`
	Description string
}

func (ExampleModel) TableName() string {
	return "examples"
}

func (e *ExampleModel) ToDomain() *example.Example {
	return &example.Example{ID: int(e.ID), Description: e.Description}
}

func FromDomain(e *example.Example) *ExampleModel {
	return &ExampleModel{ID: uint(e.ID), Description: e.Description}
}
