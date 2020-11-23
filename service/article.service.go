package service

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

	code := status.Success

	article := database.Article{
		// TagID: uint(tagID),
		Tag:         database.FindTagByID(uint(tagID)),
		Title:       title,
		Description: description,
		Content:     content,
		CreatedBy:   createdBy,
	}

	err := validate.Struct(article)
	if err == nil {
		if database.ExistTagByID(uint(tagID)) {
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
	id, _ := strconv.Atoi(c.Param("id"))
	code := status.Success

	data := make(map[string]interface{})

	err := validate.Var(id, "number")
	if err == nil {
		if database.ExistArticleByID(uint(id)) {
			data["article"] = database.GetArticle(uint(id))
		} else {
			code = status.ErrorNotExistArticle
		}

	} else {
		code = status.InvalidParams
		fmt.Println(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": status.GetMessage(code),
		"data":    data,
	})
}

//GetArticles ...
func GetArticles(c *gin.Context) {
	conditions := make(map[string]interface{})
	data := make(map[string]interface{})

	code := status.Success

	size, _ := strconv.Atoi(config.GetConfig().PageSize)

	data["articles"] = database.GetArticles(util.GetPage(c), size, conditions)
	data["total"] = database.GetArticlesCount(conditions)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": status.GetMessage(code),
		"data":    data,
	})

}

//EditArticle ...
func EditArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tagID, _ := strconv.Atoi(c.Query("tagID"))
	title := c.Query("title")
	description := c.Query("description")
	content := c.Query("content")
	updatedBy := c.Query("updatedBy")

	code := status.Success

	err := validate.Var(updatedBy, "required,max=20")
	if err == nil {

		if database.ExistArticleByID(uint(id)) {
			var article database.Article
			article.UpdatedBy = updatedBy
			if title != "" {
				article.Title = title
			}
			if database.ExistTagByID(uint(tagID)) {
				article.Tag = database.FindTagByID(uint(tagID))
			}
			article.Description = description
			article.Content = content
			database.EditArticle(uint(id), article)
		} else {
			code = status.ErrorNotExistArticle
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

//DeleteArticle ...
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := status.Success
	err := validate.Var(id, "required,number")

	if err == nil {

		if database.ExistArticleByID(uint(id)) {

			database.DeleteArticle(uint(id))
		} else {
			code = status.ErrorNotExistArticle
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
