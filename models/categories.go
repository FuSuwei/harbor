package models

type Category struct {
	ID   int    `gorm:"primary_key;type:varchar(10) unsigned auto_increment;not null" json:"id"`
	Name string `json:"name" gorm:"type:varchar(50);not null"`
}


type countArticleByCategory struct {
	Name  string `json:"name"`
	Count int	`json:"count"`
}

func GetCountByCategory() (categoryCount []*countArticleByCategory) {
	db.Raw("SELECT harbor_category.name, newTable.count AS count FROM harbor_category Left JOIN (Select  category_id, COUNT(*) AS count from harbor_article_categories  GROUP BY category_id) newTable ON newTable.category_id = harbor_category.id").Scan(&categoryCount)
	return
}

func GetArticleListByCategory(limit, offset int, tagName string) *[]Article {
	var articleList []Article
	var articleUuid []string
	rows, _ :=db.Raw("SELECT article_uuid FROM harbor_article_categories WHERE category_id IN( SELECT id FROM harbor_category WHERE NAME=?) limit ? offset ?", tagName, limit, offset).Rows()
	defer rows.Close()
	for rows.Next() {
		var uuid string
		rows.Scan(&uuid)
		articleUuid = append(articleUuid, uuid)
	}
	db.Order("created_on").Preload("Categories").Where(articleUuid).Find(&articleList)
	return &articleList
}