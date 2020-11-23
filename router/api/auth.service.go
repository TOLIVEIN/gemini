package api

import (
	"gemini/database"
	"gemini/status"
	"gemini/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CheckAuth ...
func CheckAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	user := database.Auth{Username: username, Password: password}

	validErrors := validate.Struct(user)
	data := make(map[string]interface{})
	code := status.Success
	if validErrors == nil {
		authCode := database.CheckAuth(username, password)
		if authCode == status.Success {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = status.ErrorAuthToken
			} else {
				data["token"] = token
				code = status.Success
			}
		} else {
			code = authCode
		}
	} else {

		log.Println(validErrors)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  status.GetMessage(code),
		"data": data,
	})
}
