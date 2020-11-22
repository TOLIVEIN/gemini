package api

import (
	"gemini/config"
	"gemini/database"
	"gemini/status"
	"gemini/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetTags ...
func GetTags(c *gin.Context) {

	conditions := make(map[string]interface{})
	data := make(map[string]interface{})

	code := status.Success

	size, _ := strconv.Atoi(config.CONFIG.PageSize)

	data["tags"] = database.GetTags(util.GetPage(c), size, conditions)
	data["total"] = database.GetTagsCount(conditions)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": status.GetMessage(code),
		"data":    data,
	})

}

//CreateTag ...
func CreateTag(c *gin.Context) {
	name := c.Query("name")
	createdBy := c.Query("createdBy")

	// fmt.Println(name, createdBy)

	code := status.Success

	// database.CreateTag(database.Tag{
	// 	Name:      name,
	// 	CreatedBy: createdBy,
	// })

	if !database.ExistTag(name) {
		// code = status.Success
		database.CreateTag(database.Tag{
			Name:      name,
			CreatedBy: createdBy,
		})
	} else {
		code = status.ErrorExistTag
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": status.GetMessage(code),
		"data":    make(map[string]string),
	})
}

//EditTag ...
func EditTag(c *gin.Context) {

}

//DeleteTag ...
func DeleteTag(c *gin.Context) {

}
