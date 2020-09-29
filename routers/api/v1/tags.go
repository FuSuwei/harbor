package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"harbor/models"
	"harbor/pkg/app"
	"harbor/pkg/e"
	"net/http"
)

func GetTagsList(c *gin.Context)  {
	tags := models.GetCountByTag()
	context := app.Gin{c}
	v, err := json.Marshal(tags)
	if err != nil{
		context.Response(200, e.JSON_DECODE_ERROR, "", err.Error())
	}
	context.Response(200, 200, string(v), "")
}


func IndexByTag(c *gin.Context){
	tagName := c.Param("tagName")
	c.HTML(http.StatusOK, "indexTag.html", tagName)
}