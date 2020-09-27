package utils

import (
	"github.com/gin-gonic/gin"
	"harbor/pkg/toml"
	"strconv"
)

func GetPage(c *gin.Context) int {
	page := 0
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return page
	} else {
		return (page - 1) * toml.GetAppConfig().PageSize
	}

}
