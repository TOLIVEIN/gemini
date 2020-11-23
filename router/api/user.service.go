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

//CreateUser ...
func CreateUser(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")

	code := status.Success

	user := database.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	err := validate.Struct(user)
	if err == nil {
		if !database.ExistUserByUsername(username) {
			// code = status.Success
			database.CreateUser(user)
		} else {
			code = status.ErrorExistUser
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

//GetUsers ...
func GetUsers(c *gin.Context) {
	conditions := make(map[string]interface{})
	data := make(map[string]interface{})

	code := status.Success

	size, _ := strconv.Atoi(config.GetConfig().PageSize)

	data["users"] = database.GetUsers(util.GetPage(c), size, conditions)
	data["total"] = database.GetUsersCount(conditions)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": status.GetMessage(code),
		"data":    data,
	})
}

//GetUser ...
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := status.Success

	data := make(map[string]interface{})

	err := validate.Var(id, "number")
	if err == nil {
		if database.ExistUserByID(uint(id)) {
			data["user"] = database.GetUser(uint(id))
		} else {
			code = status.ErrorNotExistUser
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

//EditUser ...
func EditUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	username := c.Query("username")
	password := c.Query("password")
	email := c.Query("email")

	code := status.Success

	data := database.User{
		Username: username,
		Password: password,
		Email:    email,
	}

	err := validate.Struct(data)
	if err == nil {

		if database.ExistUserByID(uint(id)) {
			var user database.User

			user.Username = username
			user.Password = password
			user.Email = email
			database.EditUser(uint(id), user)
		} else {
			code = status.ErrorNotExistUser
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

//DeleteUser ...
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	code := status.Success
	err := validate.Var(id, "required,number")

	if err == nil {

		if database.ExistUserByID(uint(id)) {

			database.DeleteUser(uint(id))
		} else {
			code = status.ErrorNotExistUser
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
