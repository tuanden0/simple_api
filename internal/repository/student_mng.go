package repository

import (
	"fmt"

	"github.com/tuanden0/simple_api/internal/models"
	"gorm.io/gorm"
)

// Create StudentMng struct to handle all db stuff
type StudentMng struct {
	objects *gorm.DB
}

func NewStudentRepo(db *gorm.DB) StudentObject {
	return &StudentMng{objects: db}
}

func NewPagination(p int, l int) *Pagination {
	return &Pagination{
		Page:  p,
		Limit: l,
	}
}

func NewSort(f string, asc string) *Sort {
	return &Sort{
		Field: f,
		ASC:   asc,
	}
}

// Create Student
func (m *StudentMng) Create(s *models.Student) error {
	if err := m.objects.Create(s).Error; err != nil {
		return err
	}
	return nil
}

// Retrieve Student
func (m *StudentMng) Retrieve(id string) (*models.Student, error) {
	// Create object student to return
	student := models.Student{}

	// Query student id with database
	if err := m.objects.First(&student, id).Error; err != nil {
		// Handle error if unable to connect database
		return nil, err
	}

	return &student, nil
}

// Update Student
func (m *StudentMng) Update(id string, in models.Student) (*models.Student, error) {
	// Create in object to bind value of user input
	student := models.Student{}

	// Get student from database
	if err := m.objects.Where("id = ?", id).First(&student).Error; err != nil {
		return nil, err
	}

	// Update student and avoid student update their id
	if err := m.objects.Model(&student).Omit("id").Updates(&in).Error; err != nil {
		return nil, err
	}

	return &student, nil
}

// Delete Student
func (m *StudentMng) Delete(id string) error {
	if err := m.objects.Delete(&models.Student{}, id).Error; err != nil {
		return err
	}

	return nil
}

// List Student
func (m *StudentMng) List(page Pagination, sort Sort) ([]*models.Student, error) {
	// Create slice of student to return
	students := make([]*models.Student, 0)

	// Query and pagination student
	if err := m.objects.
		Order(fmt.Sprintf("%v %v", sort.Field, sort.ASC)).
		Limit(page.Limit).
		Offset(page.Limit * (page.Page - 1)).
		Find(&students).
		Error; err != nil {
		// Handle error if unable to connect database
		return students, err
	}

	return students, nil
}
