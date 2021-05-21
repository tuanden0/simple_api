package models

type Course struct {
	Name      string     `json:"name" gorm:"size:255"`
	Semester  uint       `json:"semester"`
	Lect_hour uint       `json:"lect_hour"`
	Lab_hour  uint       `json:"lab_hour"`
	Credits   uint       `json:"credits"`
	DeptID    uint       `json:"dept_id"`
	Dept      Department `gorm:"constraint:OnDelete:CASCADE"`
}
