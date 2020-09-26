package main

import (
	"fmt"
	"harbor/models"
)

func main() {
	c := models.Article{}
	c.ID = 1
	c.GetArticle()
	fmt.Println(c)
}
