package main

import (
	"fmt"
	"harbor/models"
)

func main() {
	c := models.Article{}
	a, b := c.GetArticle()
	fmt.Println(a, b)
}
