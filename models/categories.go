package models

type Category struct {
	ID   int    `gorm:"primary_key;type:varchar(10) unsigned auto_increment;not null" json:"id"`
	Name string `json:"name" gorm:"type:varchar(50);not null"`
}


type countArticleByCategory struct {
	Name  string
	Count int
}

func GetCountByCategory() (categoryCount []*countArticleByCategory) {
	db.Raw("SELECT harbor_category.name, newTable.count AS count FROM harbor_category Left JOIN (Select  category_id, COUNT(*) AS count from harbor_article_categories  GROUP BY category_id) newTable ON newTable.category_id = harbor_category.id").Scan(&categoryCount)
	return
}
