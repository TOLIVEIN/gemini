package router

import (
	"gemini/config"
	"gemini/middleware"
	"gemini/service"
	"gemini/status"
	"gemini/util"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

//Init ...
func Init() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery(), middleware.CORS())

	gin.SetMode(config.GetConfig().RunMode)

	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code":    status.Success,
			"message": status.GetMessage(status.Success),
		})
	})

	r.POST("/auth", service.CheckAuth)

	r.POST("/upload", service.UploadImage)

	// fmt.Println(util.GetImageFullPath())
	// fmt.Println(strings.Replace(util.GetImageFullPath(), "\\", "/", -1))
	dir, _ := os.Getwd()
	r.StaticFS("/upload/images", http.Dir(dir+util.GetImageFullPath()))

	api := r.Group("/api")
	// api.Use(middleware.JWT())
	{
		tag := api.Group("/tags")
		{
			tag.GET("", service.GetTags)
			tag.Use(middleware.JWT()).POST("", service.CreateTag)
			tag.Use(middleware.JWT()).PUT(":id", service.EditTag)
			tag.Use(middleware.JWT()).DELETE(":id", service.DeleteTag)

		}

		article := api.Group("/articles")
		{
			article.GET("", service.GetArticles)
			article.GET(":id", service.GetArticle)
			article.Use(middleware.JWT()).POST("", service.CreateArticle)
			article.Use(middleware.JWT()).PUT(":id", service.EditArticle)
			article.Use(middleware.JWT()).DELETE(":id", service.DeleteArticle)
		}

		user := api.Group("/users")
		// user.Use(middleware.JWT())
		{
			user.POST("", service.CreateUser)
			user.Use(middleware.JWT()).GET("", service.GetUsers)
			user.Use(middleware.JWT()).GET(":id", service.GetUser)
			user.Use(middleware.JWT()).PUT(":id", service.EditUser)
			user.Use(middleware.JWT()).DELETE(":id", service.DeleteUser)
		}
	}

	return r
}
