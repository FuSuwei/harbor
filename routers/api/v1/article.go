package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"harbor/models"
	"harbor/pkg/app"
	"harbor/pkg/e"
	"harbor/pkg/toml"
	"harbor/pkg/utils"
	"net/http"
)

func GetArticleList(c *gin.Context) {
	isNextPage := true
	page := utils.GetPage(c)
	limit := toml.GetAppConfig().PageSize
	articles := models.GetArticleList(limit, page)
	v, err := json.Marshal(articles)
	if toml.GetAppConfig().PageSize > len(*articles){
		isNextPage = false
	}
	if err != nil{
		c.JSON(403, gin.H{
			"code":   e.JSON_DECODE_ERROR,
			"msg":    e.GetMsg(e.JSON_DECODE_ERROR),
			"data":   "",
			"errors": "",
			"isNextPage": isNextPage,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    e.GetMsg(200),
		"data":   string(v),
		"errors": "",
		"isNextPage": isNextPage,
	})
	return
}

func GetArticleListByTag(c *gin.Context){
	isNextPage := true
	page := utils.GetPage(c)
	limit := toml.GetAppConfig().PageSize
	tagName := c.Param("tagName")
	articles := models.GetArticleListBuTag(limit, page, tagName)
	v, err := json.Marshal(articles)
	if toml.GetAppConfig().PageSize > len(*articles){
		isNextPage = false
	}
	if err != nil{
		c.JSON(403, gin.H{
			"code":   e.JSON_DECODE_ERROR,
			"msg":    e.GetMsg(e.JSON_DECODE_ERROR),
			"data":   "",
			"errors": "",
			"isNextPage": isNextPage,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    e.GetMsg(200),
		"data":   string(v),
		"errors": "",
		"isNextPage": isNextPage,
	})
	return
}

func GetArticleListByCategory(c *gin.Context){
	isNextPage := true
	page := utils.GetPage(c)
	limit := toml.GetAppConfig().PageSize
	tagName := c.Param("tagName")
	articles := models.GetArticleListByCategory(limit, page, tagName)
	v, err := json.Marshal(articles)
	if toml.GetAppConfig().PageSize > len(*articles){
		isNextPage = false
	}
	if err != nil{
		c.JSON(403, gin.H{
			"code":   e.JSON_DECODE_ERROR,
			"msg":    e.GetMsg(e.JSON_DECODE_ERROR),
			"data":   "",
			"errors": "",
			"isNextPage": isNextPage,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":   200,
		"msg":    e.GetMsg(200),
		"data":   string(v),
		"errors": "",
		"isNextPage": isNextPage,
	})
	return
}

func GetArticle(c *gin.Context){
	context := app.Gin{c}
	articleUuid := c.Param("articleUuid")
	article := models.GetArticle(articleUuid)
	v, err := json.Marshal(article)
	if err != nil{
		context.Response(200, e.JSON_DECODE_ERROR, "", err.Error())
	}
	context.Response(200, 200, string(v), "")
}