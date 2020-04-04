package routers

import (
	"SUweb/controlers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	auth := router.Group("/auth")
	{
		auth.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message":"test",
			})
		})
		auth.POST("/register",controlers.Register)
		auth.POST("/login",controlers.Login)
	}
	project := router.Group("/project")
	{
		project.POST("/createProject",controlers.CreateProject)
		project.GET("/getProjectList",controlers.GetProjectList)
		project.GET("/getProject",controlers.GetProject)
		project.POST("/comment",controlers.CommentProject)
		project.GET("/getcomments",controlers.GetComments)
	}
	alumni := router.Group("/alumni")
	{
		alumni.GET("getMemList",controlers.GetMemList)
	}
	blog := router.Group("/blog")
	{
		blog.GET("getBlogList",controlers.GetBlogList)
		blog.GET("getBlog",controlers.GetBlog)
	}
	return router
}
