package models

type Student struct {
	Id   uint    `json:"id,omitempty" gorm:"primaryKey"`
	Name string  `json:"name,omitempty"`
	GPA  float64 `json:"gpa,omitempty"`
}

func (Student) TableName() string {
	return "student"
}
