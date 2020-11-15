package controlers

import (
	"../e"
	"../models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetMemList(c *gin.Context) {
	var alumnis []models.Alumni
	if err := models.QueryAlumni(&alumnis);err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success","projects":alumnis})
	}
}

func CreateAlumni(c *gin.Context) {
	var alumni models.Alumni
	var token models.Token
	if err := c.ShouldBindHeader(&token); err != nil ||  token.Query() {
		c.JSON(http.StatusUnauthorized, gin.H{"status": e.ERROR_NOT_LOGIN,"message":e.GetMsg(e.ERROR_NOT_LOGIN)})
	} else if err := c.ShouldBindJSON(&alumni); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS,"message":e.GetMsg(e.INVALID_PARAMS)})
	} else if err := alumni.CreateAlumni(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func GetAlumniInfo(c *gin.Context) {
	username := c.Query("username")
	var alumni *models.Alumni
	if err := alumni.QueryAlumni(username); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, alumni)
	}
}