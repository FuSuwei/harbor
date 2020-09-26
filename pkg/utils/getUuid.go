package utils

import (
	uuid "github.com/satori/go.uuid"
	"strings"
)

func GetUuid() string {
	ul := uuid.NewV4().String()
	return strings.ReplaceAll(ul, "-", "")
}
