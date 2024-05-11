package repositorie

import "github.com/peruccii/microservicos_go/internal/entities"

type ICategoryRepository interface {
	Save(category *entities.Category) error
	List() ([]*entities.Category, error) 
}