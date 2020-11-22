package api

import (
	"fmt"
	"gemini/config"
	"gemini/database"
	"gemini/status"
	"gemini/util"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//CreateArticle ...
func CreateArticle(c *gin.Context) {
	tagID, _ := strconv.Atoi(c.Query("tagID"))
	title := c.Query("title")
	description := c.Query("description")
	content := c.Query("content")
	createdBy := c.Query("createdBy")

	// fmt.Println(name, createdBy)

	code := status.Success

	// database.CreateTag(database.Tag{
	// 	Name:      name,
	// 	CreatedBy: createdBy,
	// })

	article := database.Article{
		TagID:       uint(tagID),
		Title:       title,
		Description: description,
		Content:     content,
		CreatedBy:   createdBy,
	}

	err := validate.Struct(article)
	if err == nil {
		if database.ExistTagByID(uint(tagID)) {
			// code = status.Success
			// database.CreateTag(tag)
			database.CreateArticle(article)
		} else {
			code = status.ErrorNotExistTag
		}

	} else {
		code = status.InvalidParams
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": status.GetMessage(code),
		"data":    make(map[string]string),
	})
}

//GetArticle ...
func GetArticle(c *gin.Context) {
}

//GetArticles ...
func GetArticles(c *gin.Context) {
	conditions := make(map[string]interface{})
	data := make(map[string]interface{})

	code := status.Success

	size, _ := strconv.Atoi(config.CONFIG.PageSize)

	data["articles"] = database.GetTags(util.GetPage(c), size, conditions)
	data["total"] = database.GetTagsCount(conditions)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": status.GetMessage(code),
		"data":    data,
	})

}

//EditArticle ...
func EditArticle(c *gin.Context) {

}

//DeleteArticle ...
func DeleteArticle(c *gin.Context) {

}
