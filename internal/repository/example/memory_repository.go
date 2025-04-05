package example

import (
	"context"
	"fmt"
	"sync"

	"starter-go/internal/domain/example"
)

// MemoryRepository is an in-memory implementation of the example.ExampleRepository interface
type MemoryRepository struct {
	examples map[int]*example.Example
	nextID   int
	mu       sync.RWMutex // For thread safety
}

// NewMemoryRepository creates a new in-memory repository
func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{
		examples: make(map[int]*example.Example),
		nextID:   1,
	}
}

// FindByID retrieves an example by its ID
func (r *MemoryRepository) FindByID(ctx context.Context, id int) (*example.Example, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if ex, ok := r.examples[id]; ok {
		// Return a copy to prevent external modification
		return &example.Example{
			ID:          ex.ID,
			Description: ex.Description,
		}, nil
	}
	return nil, fmt.Errorf("example with ID %d not found", id)
}

// Save stores an example
func (r *MemoryRepository) Save(ctx context.Context, e *example.Example) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// For new examples
	if e.ID == 0 {
		e.ID = r.nextID
		r.nextID++
	}

	// Store a copy to prevent external modification
	exampleCopy := &example.Example{
		ID:          e.ID,
		Description: e.Description,
	}
	r.examples[e.ID] = exampleCopy

	return nil
}

// Preload adds some example data to the repository
func (r *MemoryRepository) Preload(ctx context.Context, examples ...*example.Example) error {
	for _, e := range examples {
		if err := r.Save(ctx, e); err != nil {
			return err
		}
	}
	return nil
}
