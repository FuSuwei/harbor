package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"harbor/pkg/utils"
	"time"
)

type Article struct {
	Model

	Title     string `json:"title" gorm:"type:varchar(100);not null"`
	Summary   string `json:"summary" gorm:"type:varchar(250);not null"`
	ContentId string `json:"content_id" gorm:"type:varchar(10);not null"`
	CreatedOn string `json:"created_on" gorm:"type:varchar(10);not null"`

	Tags       []Tag      `json:"tags" gorm:"many2many:article_tag;"`
	Categories []Category `json:"categories" gorm:"many2many:article_categories;"`
	Content    Content    `json:"content" gorm:"foreignkey:content_uuid"`
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	err := scope.SetColumn("CreatedOn", time.Now().Unix())
	if err != nil {
		return err
	}
	err = scope.SetColumn("Uuid", utils.GetUuid())
	return err
}

func GetArticleList(limit, offset int) *[]Article {
	var articleList []Article
	db.Preload("Categories").Limit(limit).Offset(offset).Find(&articleList)
	return &articleList
}

func (article *Article) GetArticle() {
	db.First(article, "uuid")
	db.Model(article).Related(&article.Tags, "Tags").Related(&article.Content, ).Related(&article.Categories, "Categories")
	return
}