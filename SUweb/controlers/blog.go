package controlers

import (
	"../e"
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetBlogList(c *gin.Context) {
	var blogs []models.Blog
	if err := models.QueryBlogList(&blogs); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME, "message": e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "blogs": blogs})
	}
}

func GetBlogNums(c *gin.Context) {
	result := models.QueryBlogNums()
	c.JSON(http.StatusOK, result)
}

func GetBlog(c *gin.Context) {
	blog_id := c.Query("blog_id")
	var blog models.Blog
	if err := blog.QueryBlog(blog_id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME, "message": e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "project": blog})
	}
}

func UpdateBlog(c *gin.Context) {
	var blog models.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
	} else if err := blog.UpdateBlog(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
	}else {
		c.JSON(http.StatusOK,gin.H{"message":"ok"})
	}
}

func PublishBlog(c *gin.Context) {
	var blog models.Blog
	var token models.Token
	if err := c.ShouldBindHeader(&token); err != nil || token.Query() {
		c.JSON(http.StatusUnauthorized, gin.H{"status": e.ERROR_NOT_LOGIN, "message": e.GetMsg(e.ERROR_NOT_LOGIN)})
	} else if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
	} else if err := blog.CreateBlog(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME, "message": e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func GetBlogComments(c *gin.Context) {
	var comments []models.BlogComment
	var project_id = c.Query("project_id")
	if err := models.QueryBlogComments(&comments, project_id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME, "message": e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "comments": comments})
	}
}

func GetMyBlogList(c *gin.Context) {
	userid, err := strconv.Atoi(c.Query("user_id"))
	var blogs []models.Blog
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME, "message": e.GetMsg(e.ERROR_EXIST_NAME)})
	} else if err := models.QueryBlogListByUserId(&blogs, uint(userid)); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME, "message": e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success", "blogs": blogs})
	}
}

func WatchLater(c *gin.Context) {
	var blogId uint
	var userId uint
	var token models.Token
	if err := c.ShouldBindHeader(&blogId); err != nil || token.Query() {
		c.JSON(http.StatusUnauthorized, gin.H{"status": e.ERROR_NOT_LOGIN, "message": e.GetMsg(e.ERROR_NOT_LOGIN)})
	} else if err := c.ShouldBindJSON(&userId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
	} else {
		wl := models.WatchLater{
			UserId: userId,
			BlogId: blogId,
		}
		if err := wl.Create(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "success"})
		}
	}
}

func GetWathLaterList(c *gin.Context) {
	userid := c.Query("user_id")
	var wls []models.WatchLater
	var blogs []models.Blog
	if err := models.QueryWatchLater(&wls, userid); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME, "message": e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		for _, wl := range wls {
			var blog models.Blog
			blog.QueryBlog(strconv.Itoa(int(wl.BlogId)))
			blogs = append(blogs, blog)
		}
		c.JSON(http.StatusOK, gin.H{"blogs": blogs})
	}
}

func AddBrowse(c *gin.Context) {
	var blogId string
	var blog models.Blog
	if err := c.ShouldBindHeader(&blogId); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": e.ERROR_NOT_LOGIN, "message": e.GetMsg(e.ERROR_NOT_LOGIN)})
	} else if err := blog.QueryBlog(blogId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
	} else {
		blog.ReadTimes = blog.ReadTimes + 1
		if err := blog.UpdateBlog(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "message": e.GetMsg(e.INVALID_PARAMS)})
		} else {
			c.JSON(http.StatusOK, gin.H{"message": "ok"})
		}
	}
}
