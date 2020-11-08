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