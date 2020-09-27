package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"harbor/models"
	"harbor/pkg/app"
	"harbor/pkg/toml"
	"harbor/pkg/utils"
)

func GetArticleList(c *gin.Context) {
	context := app.Gin{c}
	page := utils.GetPage(c)
	limit := toml.GetAppConfig().PageSize
	articles := models.GetArticleList(limit, page)
	v, err := json.Marshal(articles)
	if err != nil{
		context.Response(403, 30000, "", "")
	}
	context.Response(200, 200, string(v), "")
}
