package controlers

import (
	"SUweb/e"
	"SUweb/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetBlogList(c *gin.Context) {
	var blogs []models.Blog
	if err := models.QueryBlogList(&blogs);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success","blogs":blogs})
	}
}

func GetBlog(c *gin.Context) {
	blog_id := c.Query("blog_id")
	var blog models.Blog
	if err := blog.QueryBlog(blog_id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success","project":blog})
	}
}

func PublishBlog(c *gin.Context) {
	var blog models.Blog
	var token models.Token
	if err := c.ShouldBindHeader(&token); err != nil ||  token.Query() {
		c.JSON(http.StatusUnauthorized, gin.H{"status": e.ERROR_NOT_LOGIN,"message":e.GetMsg(e.ERROR_NOT_LOGIN)})
	} else if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS,"message":e.GetMsg(e.INVALID_PARAMS)})
	} else if err := blog.CreateBlog(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}