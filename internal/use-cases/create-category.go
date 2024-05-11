package use_cases

import (
	"github.com/peruccii/microservicos_go/internal/entities"
	"github.com/peruccii/microservicos_go/internal/repositorie"
)

type createCategoryUseCase struct {
	repository repositorie.ICategoryRepository
}

func NewCreateCategoryUseCase(repository repositorie.ICategoryRepository) *createCategoryUseCase {
	return &createCategoryUseCase{repository}
}

func (u *createCategoryUseCase) Execute(name string) error {
	category, err := entities.NewCategory(name)

	if err != nil {
		return err
	}
	
	err = u.repository.Save(category)

	if err != nil {
		return err
	}

	return nil
}