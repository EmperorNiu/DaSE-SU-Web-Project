package routers

import (
	"../controlers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors())
	auth := router.Group("/api/auth")
	{
		auth.GET("/test", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message":"test",
			})
		})
		auth.POST("/register",controlers.Register)
		auth.POST("/login",controlers.Login)
	}
	project := router.Group("/api/project")
	{
		project.POST("/createProject",controlers.CreateProject)
		project.GET("/getProjectList",controlers.GetProjectList)
		project.GET("/getProject",controlers.GetProject)
		project.POST("/comment",controlers.CommentProject)
		project.GET("/getcomments",controlers.GetComments)
	}
	alumni := router.Group("/api/alumni")
	{
		alumni.GET("getMemList",controlers.GetMemList)
	}
	blog := router.Group("/api/blog")
	{
		blog.GET("getBlogList",controlers.GetBlogList)
		blog.GET("getBlog",controlers.GetBlog)
		blog.POST("publishBlog",controlers.PublishBlog)
	}
	achievement := router.Group("/api/achievement")
	{
		achievement.GET("get_proj_list",controlers.GetProjectList)
		achievement.GET("getProjid",controlers.GetProject)
		achievement.POST("/publishProj",controlers.CreateProject)
	}
	return router
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}