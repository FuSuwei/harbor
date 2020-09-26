package models


type Content struct {
	Model
	Content     string `json:"name" gorm:"type:text;not null"`
}
