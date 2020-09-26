package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

type Article struct {
	Model

	Title     string `json:"title" gorm:"type:varchar(100);not null"`
	Summary   string `json:"summary" gorm:"type:varchar(250);not null"`
	ContentId string `json:"content_id" gorm:"type:varchar(10);not null"`
	CreatedOn string `json:"created_on" gorm:"type:varchar(10);not null"`

	Tags []Tag `json:"tags" gorm:"many2many:article_tag;"`
	Content Content `json:"content" gorm:"many2many:article_tag;"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	return scope.SetColumn("CreatedOn", time.Now().Unix())
}

func (article *Article) GetArticle() {
	db.First(article)
	db.Model(article).Related(&article.Tags, "Tags").Related(&article.Content, "Content")
	return
}
