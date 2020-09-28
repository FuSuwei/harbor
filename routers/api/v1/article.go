package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"harbor/models"
	"harbor/pkg/e"
	"harbor/pkg/toml"
	"harbor/pkg/utils"
	"net/http"
)

func GetArticleList(c *gin.Context) {
	isNextPage := false

	page := utils.GetPage(c)
	limit := toml.GetAppConfig().PageSize
	articles := models.GetArticleList(limit, page)
	v, err := json.Marshal(articles)
	articleCount := models.GetArticleCount()
	if page <= articleCount-toml.GetAppConfig().PageSize{
		isNextPage = true
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
