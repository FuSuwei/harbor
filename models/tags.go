package models

type Tag struct {
	ID   int    `gorm:"primary_key;type:varchar(10) unsigned auto_increment;not null" json:"id"`
	Name string `json:"name" gorm:"type:varchar(50);not null"`
}

type countArticleByTag struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func GetCountByTag() (tagCount []*countArticleByTag) {
	db.Raw("SELECT harbor_tag.name, newTable.count AS count FROM harbor_tag  Left JOIN (Select  tag_id, COUNT(*) AS count from harbor_article_tag  GROUP BY tag_id) newTable ON newTable.tag_id = harbor_tag.id").Scan(&tagCount)
	return
}
