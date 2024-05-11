package entities

import (
	"fmt"
	"time"
)

type Category struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

// Regras de negócio na entidade ( para Microserviços )

func NewCategory(name string) (*Category, error) {
		category := &Category {
			Name: name,
			CreatedAt: time.Now(),
			UpdateAt: time.Now(),
	}
	//business rules
	err := category.IsValid()

	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *Category) IsValid() error {
	if (len(c.Name) < 5) {
		return fmt.Errorf("name must be grater than 5. got %d", len(c.Name))
	}
	return nil
}

