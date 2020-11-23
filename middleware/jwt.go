package middleware

import (
	"gemini/status"
	"gemini/util"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//JWT ...
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		code := status.Success

		token := c.Query("token")
		if token == "" {
			code = status.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				code = status.ErrorAuthCheckTokenFail
			} else if time.Now().Unix() > claims.ExpiresAt {
				code = status.ErrorAuthCheckTokenTimeout
			}
		}
		if code != status.Success {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": status.GetMessage(code),
				"data":    data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
