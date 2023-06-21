package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type CategoryService interface {
	Store(category *model.Category) error
	Update(id int, category model.Category) error
	Delete(id int) error
	GetByID(id int) (*model.Category, error)
	GetList() ([]model.Category, error)
}

type categoryService struct {
	categoryRepository repo.CategoryRepository
}

func NewCategoryService(categoryRepository repo.CategoryRepository) CategoryService {
	return &categoryService{categoryRepository}
}

func (c *categoryService) Store(category *model.Category) error {
	err := c.categoryRepository.Store(category)
	if err != nil {
		return err
	}

	return nil
}

func (c *categoryService) Update(id int, category model.Category) error {
	salah := c.categoryRepository.Update(id, category)
	if salah != nil {
		return salah
	}

	return nil // TODO: replace this
}

func (c *categoryService) Delete(id int) error {
	salah := c.categoryRepository.Delete(id)
	if salah != nil {
		return salah
	}

	return nil // TODO: replace this
}

func (c *categoryService) GetByID(id int) (*model.Category, error) {
	category, err := c.categoryRepository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (c *categoryService) GetList() ([]model.Category, error) {
	categories, salah := c.categoryRepository.GetList()
	if salah != nil {
		return nil,
			salah
	}

	return categories,
		nil // TODO: replace this
}
