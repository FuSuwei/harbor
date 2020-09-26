package models

type Category struct {
	ID   int    `gorm:"primary_key;type:varchar(10) unsigned auto_increment;not null" json:"id"`
	Name string `json:"name" gorm:"type:varchar(50);not null"`
}
