package controlers

import (
	"../e"
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateProject(c *gin.Context) {
	var project models.Project
	var token models.Token
	if err := c.ShouldBindHeader(&token); err != nil ||  token.Query() {
		c.JSON(http.StatusUnauthorized, gin.H{"status": e.ERROR_NOT_LOGIN,"message":e.GetMsg(e.ERROR_NOT_LOGIN)})
	} else if err := c.ShouldBind(&project); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS,"message":e.GetMsg(e.INVALID_PARAMS)})
	} else if err := project.Create(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func GetProjectList(c *gin.Context) {
	var projects []models.Project
	var token models.Token
	err := c.ShouldBindHeader(&token)
	if err != nil ||  token.Query() {
		c.JSON(http.StatusUnauthorized, gin.H{"status": e.ERROR_NOT_LOGIN,"message":e.GetMsg(e.ERROR_NOT_LOGIN)})
	} else if err := models.QueryProjectList(&projects); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success","projects":projects})
	}
}

func GetProject(c *gin.Context) {
	var project models.Project
	project_id := c.DefaultQuery("id","0")
	var token models.Token
	if err := c.ShouldBindHeader(&token); err != nil ||  token.Query() {
		c.JSON(http.StatusUnauthorized, gin.H{"status": e.ERROR_NOT_LOGIN,"message":e.GetMsg(e.ERROR_NOT_LOGIN)})
	} else if err := project.QueryProject(project_id); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_NOT_EXIST_PROJECT,"message":e.GetMsg(e.ERROR_NOT_EXIST_PROJECT)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success","project":project})
	}
}

func CommentProject(c *gin.Context) {
	var comment models.ProjectComment
	var token models.Token
	if err := c.ShouldBindHeader(&token); err != nil ||  token.Query() {
		c.JSON(http.StatusUnauthorized, gin.H{"status": e.ERROR_NOT_LOGIN,"message":e.GetMsg(e.ERROR_NOT_LOGIN)})
	} else if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS,"message":e.GetMsg(e.INVALID_PARAMS)})
	} else if err := comment.CreateComment(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func GetComments(c *gin.Context) {
	var comments []models.ProjectComment
	var project_id = c.Query("project_id")
	var token models.Token
	if err := c.ShouldBindHeader(&token); err != nil ||  token.Query() {
		c.JSON(http.StatusUnauthorized, gin.H{"status": e.ERROR_NOT_LOGIN,"message":e.GetMsg(e.ERROR_NOT_LOGIN)})
	} else if err := models.QueryComments(&comments,project_id);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success","comments":comments})
	}
}