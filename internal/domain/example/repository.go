package example

type ExampleRepository interface {
	FindByID(id int) (*Example, error)
	Save(example *Example) error
}
