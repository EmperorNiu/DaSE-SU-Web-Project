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
		c.JSON(http.StatusOK, gin.H{"message": "success","projects":blogs})
	}
}

func GetBlog(c *gin.Context) {
	blog_id := c.Query("blog_id")
	var blog models.Blog
	if err := blog.QueryBlog(blog_id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success","project":project})
	}
}