package models

type Group struct {
	Number uint   `json:"number" gorm:"primaryKey"`
	Major  string `json:"major" gorm:"size:255;index:,unique"`
	DeptID uint
	Dept   Department `json:"dept" gorm:"constraint:OnDelete:CASCADE"`
}
