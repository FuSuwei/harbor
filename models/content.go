package models


type Content struct {
	Model
	Content     string `json:"content" gorm:"type:text;not null"`
}
