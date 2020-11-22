package util

import (
	"gemini/config"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetPage ...
func GetPage(c *gin.Context) (result int) {
	result = 0
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(config.CONFIG.PageSize)

	if page > 0 {
		result = (page - 1) * pageSize
	}

	return

}
