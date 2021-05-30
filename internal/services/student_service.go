package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tuanden0/simple_api/internal/models"
	"github.com/tuanden0/simple_api/internal/repository"
)

type StudentService struct {
	student repository.StudentObject
}

func NewStudentService(s repository.StudentObject) *StudentService {
	return &StudentService{student: s}
}

// Create Student API
func (srv *StudentService) Create(c *gin.Context) {

	// Create in object to bind user value
	in := models.Student{}

	// Validate user input
	if err := c.ShouldBindJSON(&in); err != nil {
		log.Printf("unable to parse student: %v\n", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Prevent user insert "id" field
	student := models.Student{Name: in.Name, GPA: in.GPA}
	if err := srv.student.Create(&student); err != nil {
		log.Printf("unable to create student: %v\n", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Will return wrong ID if user send "id" field
	c.JSON(
		http.StatusCreated,
		gin.H{"data": student},
	)

}

// Retrieve Student API
func (srv *StudentService) Retrieve(c *gin.Context) {
	// Query student id with database
	student, err := srv.student.Retrieve(c.Param("id"))
	if err != nil {
		// Handle error if unable to connect database
		log.Printf("unable to get student: %v\n", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Return student data
	c.JSON(
		http.StatusOK,
		gin.H{
			"data": student,
		},
	)

}

// Update Student API
func (srv *StudentService) Update(c *gin.Context) {
	// Create in object to bind value of user input
	in := models.Student{}

	// Validate user input
	if err := c.ShouldBindJSON(&in); err != nil {
		log.Printf("unable to parse student: %v\n", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	student, err := srv.student.Update(c.Param("id"), in)
	if err != nil {
		log.Printf("unable to update student: %v\n", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		gin.H{"data": student},
	)

}

// Delete Student API
func (srv *StudentService) Delete(c *gin.Context) {
	if err := srv.student.Delete(c.Param("id")); err != nil {
		log.Printf("unable to delete student: %v\n", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		http.StatusNoContent,
		gin.H{},
	)
}

// List Student API
func (srv *StudentService) List(c *gin.Context) {

	// Get

	// Get paging
	page := mutatePagination(c)

	// Get sort
	order_by := mutateSort(c)

	students, err := srv.student.List(*page, *order_by)

	if err != nil {
		// Handle error if unable to connect database
		log.Printf("unable to get all students: %v\n", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Return students data
	c.JSON(http.StatusOK, gin.H{"data": students})

}

// Mutate Pagination
func mutatePagination(c *gin.Context) *repository.Pagination {
	// Value to pagination students list
	limit := 5
	page := 1

	// Parse paging data
	p := c.Query("page")
	if p != "" {
		page, _ = strconv.Atoi(p)
	}

	// Validate page
	if page <= 0 {
		page = 1
	}

	// Parse limit data
	l := c.Query("limit")
	if l != "" {
		limit, _ = strconv.Atoi(l)
	}

	if limit <= 0 {
		limit = 5
	}

	return repository.NewPagination(page, limit)

}

// Mutate Sort
func mutateSort(c *gin.Context) *repository.Sort {
	// Value to sort student list
	field := "id"
	isASC := true
	s := "asc"

	// Parse sort data
	f := c.Query("order_by")
	if f != "" {
		field = f
	}

	// Parse isASC
	i := c.Query("is_asc")
	if i != "" {
		isASC, _ = strconv.ParseBool(i)
	}

	if !isASC {
		s = "desc"
	}

	return repository.NewSort(field, s)

}
