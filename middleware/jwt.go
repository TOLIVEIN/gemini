package middleware

import (
	"gemini/status"
	"gemini/util"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//JWT ...
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		header := c.Request.Header
		token := strings.Join(header["token"], "")

		var data interface{}
		code := status.Success

		if token == "" {
			code = status.InvalidParams
		} else {
			claims, err := util.ParseToken(token)
			// fmt.Println(claims)
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
