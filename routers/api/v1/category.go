package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"harbor/models"
	"harbor/pkg/app"
	"harbor/pkg/e"
	"net/http"
)

func GetCategoriesList(c *gin.Context)  {
	tags := models.GetCountByCategory()
	context := app.Gin{c}
	v, err := json.Marshal(tags)
	if err != nil{
		context.Response(200, e.JSON_DECODE_ERROR, "", err.Error())
	}
	context.Response(200, 200, string(v), "")
}

func IndexByCategory(c *gin.Context){
	categoryName := c.Param("categoryName")
	c.HTML(http.StatusOK, "indexCategory.html", categoryName)
}