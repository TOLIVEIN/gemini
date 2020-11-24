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
	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

//GetTags ...
func GetTags(c *gin.Context) {

	conditions := make(map[string]interface{})
	data := make(map[string]interface{})

	code := status.Success

	size, _ := strconv.Atoi(config.GetConfig().PageSize)

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

	tag := database.Tag{}

	if err := c.ShouldBind(&tag); err != nil {
		fmt.Println(err)
	}

	code := status.Success

	err := validate.Struct(tag)
	if err == nil {
		if !database.ExistTagByName(tag.Name) {
			// code = status.Success
			database.CreateTag(tag)
		} else {
			code = status.ErrorExistTag
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

//EditTag ...
func EditTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	name := c.Query("name")
	updatedBy := c.Query("updatedBy")

	fmt.Println(id, name, updatedBy)

	code := status.Success

	err := validate.Var(updatedBy, "required,max=20")
	if err == nil {

		if database.ExistTagByID(uint(id)) {
			data := make(map[string]interface{})
			data["updated_by"] = updatedBy
			if name != "" {
				data["name"] = name
			}
			database.EditTag(uint(id), data)
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

//DeleteTag ...
func DeleteTag(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := status.Success
	err := validate.Var(id, "required,number")

	if err == nil {

		if database.ExistTagByID(uint(id)) {

			database.DeleteTag(uint(id))
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
