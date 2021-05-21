package models

type Department struct {
	Number uint   `json:"number" gorm:"primaryKey;autoIncrement:true"`
	Name   string `json:"name" gorm:"size:255"`
}
