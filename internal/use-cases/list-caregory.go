package use_cases

import (
	"github.com/peruccii/microservicos_go/internal/entities"
	"github.com/peruccii/microservicos_go/internal/repositorie"
)

type listCategoryUseCase struct {
	repository repositorie.ICategoryRepository
}

func NewListCategoryUseCase(repository repositorie.ICategoryRepository) *listCategoryUseCase {
	return &listCategoryUseCase{repository}
}

func (u *listCategoryUseCase) Execute() ([]*entities.Category,error) {
	
	categories, err := u.repository.List()

	if err != nil {
		return nil, err
	}

	return categories, nil
}