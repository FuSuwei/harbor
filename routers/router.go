package routers

import (
	"github.com/gin-gonic/gin"
	v1 "harbor/routers/api/v1"

	"harbor/pkg/toml"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(toml.GetRunMode())
	r.Static("statics", "./statics")
	r.LoadHTMLGlob("statics/html/*.html")


	getHtml := r.Group("/")
	{
		getHtml.GET("/index.html", v1.Index)
		getHtml.GET("/tags", v1.Tags)
		getHtml.GET("/categories", v1.Categories)
	}
	apiV1 := r.Group("api/v1")
	{
		apiV1.GET("/getArticleList", v1.GetArticleList)
		apiV1.GET("/getTagsList", v1.GetTagsList)
		apiV1.GET("/getCategoriesList", v1.GetCategoriesList)
	}

	return r
}