package service

import (
	"fmt"
	"gemini/database"
	"gemini/status"
	"gemini/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//CheckAuth ...
func CheckAuth(c *gin.Context) {
	// username := c.Query("username")
	// password := c.Query("password")
	user := database.Auth{}

	if err := c.ShouldBind(&user); err != nil {
		fmt.Println(err)
	}

	validErrors := validate.Struct(user)
	data := make(map[string]interface{})
	code := status.Success
	if validErrors == nil {
		authCode := database.CheckAuth(user.Username, user.Password)
		if authCode == status.Success {
			token, err := util.GenerateToken(user.Username)
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
		"code":    code,
		"message": status.GetMessage(code),
		"data":    data,
	})
}
