package repository

import "github.com/tuanden0/simple_api/internal/models"

type StudentObject interface {
	Create(s *models.Student) error
	Retrieve(id string) (*models.Student, error)
	Update(id string, in models.Student) (*models.Student, error)
	Delete(id string) error
	List(page Pagination) ([]*models.Student, error)
}

type Pagination struct {
	Page  int
	Limit int
}
