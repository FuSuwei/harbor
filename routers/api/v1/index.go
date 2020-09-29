package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context){
	c.HTML(http.StatusOK, "index.html", nil)
}
func Tags(c *gin.Context)  {
	c.HTML(http.StatusOK, "tags.html", nil)
}

func Categories(c *gin.Context)  {
	c.HTML(http.StatusOK, "categories.html", nil)
}

func ArticleIndex(c *gin.Context){
	articleUuid := c.Param("articleUuid")
	c.HTML(http.StatusOK, "article.html", articleUuid)
}