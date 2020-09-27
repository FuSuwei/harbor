package app

import (
	"github.com/gin-gonic/gin"
	"harbor/pkg/e"
)

type Gin struct {
	C *gin.Context
}

func (g *Gin) Response(httpCode, errCode int, data interface{}, errors string) {
	g.C.JSON(httpCode, gin.H{
		"code":   errCode,
		"msg":    e.GetMsg(errCode),
		"data":   data,
		"errors": errors,
	})
	return
}
