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
		tag := g.Group("/tags")
		{
			tag.GET("", api.GetTags)
			tag.POST("", api.CreateTag)
			tag.PUT(":id", api.EditTag)
			tag.DELETE(":id", api.DeleteTag)

		}

		article := g.Group("/articles")
		{
			article.GET("", api.GetArticles)
			article.GET(":id", api.GetArticle)
			article.POST("", api.CreateArticle)
			article.PUT(":id", api.EditArticle)
			article.DELETE(":id", api.DeleteArticle)
		}

		// user := g.Group("/user")
		// {
		// }
	}

	return r
}
