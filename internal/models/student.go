package models

import "time"

type Student struct {
	Id      uint      `json:"id" gorm:"primaryKey"`
	Name    string    `json:"name"`
	Dob     time.Time `json:"dob"`
	Picture string    `json:"picture"`
	GPA     float64   `json:"gpa"`
	Group   []Group   `json:"group" gorm:"constraint:OnDelete:CASCADE"`
	Courses []Course  `gorm:"many2many:student_courses;"`
}
