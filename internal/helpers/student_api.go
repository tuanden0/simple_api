package helpers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tuanden0/simple_api/internal/models"
)

// Create Student API
func CreateStudent(c *gin.Context) {

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
	if err := models.DB.Create(&student).Error; err != nil {
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
func RetrieveStudent(c *gin.Context) {
	// Create object student to return
	student := models.Student{}

	// Query student id with database
	if err := models.DB.First(&student, c.Param("id")).Error; err != nil {
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
func UpdateStudent(c *gin.Context) {
	// Create in object to bind value of user input
	in := models.Student{}
	student := models.Student{}

	// Validate user input
	if err := c.ShouldBindJSON(&in); err != nil {
		log.Printf("unable to parse student: %v\n", err.Error())
		c.JSON(
			http.StatusBadRequest,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Get student from database
	if err := models.DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		log.Printf("unable to get student: %v\n", err.Error())
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)
		return
	}

	// Update student and avoid student update their id
	if err := models.DB.Model(&student).Omit("id").Updates(&in).Error; err != nil {
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
func DeleteStudent(c *gin.Context) {
	if err := models.DB.Delete(&models.Student{}, c.Param("id")).Error; err != nil {
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
func ListStudent(c *gin.Context) {

	// Create slice of student to return
	students := []models.Student{}

	// Value to pagination students list
	size := 5
	page := 1

	p := c.Query("page")
	if p != "" {
		page, _ = strconv.Atoi(p)
	}

	// Validate page
	if page <= 0 {
		log.Printf("page must be larger than 0")
		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": "page must be larger than 0"},
		)
		return
	}

	// Query and pagination student
	if err := models.DB.
		Limit(size).
		Offset(size * (page - 1)).
		Find(&students).Error; err != nil {
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
