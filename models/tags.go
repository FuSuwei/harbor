package models


type Tag struct {
	Model
	Name     string `json:"name" gorm:"type:varchar(50);not null"`
}