package router

import (
	"gemini/config"
	"gemini/router/api"
	"gemini/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Init ...
func Init() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	gin.SetMode(config.CONFIG.RunMode)

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    status.Success,
			"message": status.GetMessage(status.Success),
		})
	})

	g := r.Group("/api")
	{
		g.GET("/tags", api.GetTags)
		g.POST("/tags", api.CreateTag)
		g.PUT("tags/:id", api.EditTag)
		g.DELETE("tags/:id", api.DeleteTag)
	}

	return r
}
