package router

import (
	"gemini/config"
	"gemini/middleware"
	"gemini/router/api"
	"gemini/status"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Init ...
func Init() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	gin.SetMode(config.GetConfig().RunMode)

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    status.Success,
			"message": status.GetMessage(status.Success),
		})
	})

	r.GET("/auth", api.CheckAuth)

	g := r.Group("/api")
	// g.Use(middleware.JWT())
	{
		tag := g.Group("/tags")
		{
			tag.GET("", api.GetTags)
			tag.Use(middleware.JWT()).POST("", api.CreateTag)
			tag.Use(middleware.JWT()).PUT(":id", api.EditTag)
			tag.Use(middleware.JWT()).DELETE(":id", api.DeleteTag)

		}

		article := g.Group("/articles")
		{
			article.GET("", api.GetArticles)
			article.GET(":id", api.GetArticle)
			article.Use(middleware.JWT()).POST("", api.CreateArticle)
			article.Use(middleware.JWT()).PUT(":id", api.EditArticle)
			article.Use(middleware.JWT()).DELETE(":id", api.DeleteArticle)
		}

		user := g.Group("/users")
		// user.Use(middleware.JWT())
		{
			user.POST("", api.CreateUser)
			user.Use(middleware.JWT()).GET("", api.GetUsers)
			user.Use(middleware.JWT()).GET(":id", api.GetUser)
			user.Use(middleware.JWT()).PUT(":id", api.EditUser)
			user.Use(middleware.JWT()).DELETE(":id", api.DeleteUser)
		}
	}

	return r
}
