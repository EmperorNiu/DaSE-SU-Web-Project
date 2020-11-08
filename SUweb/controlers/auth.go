package controlers

import (
	"../e"
	"../models"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"time"
)

func Register(c *gin.Context) {
	var user models.Auth
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS,"message":e.GetMsg(e.INVALID_PARAMS)})
	} else if err := user.Create(); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": e.ERROR_EXIST_NAME,"message":e.GetMsg(e.ERROR_EXIST_NAME)})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "success"})
	}
}

func Login(c *gin.Context) {
	var user,_user models.Auth
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
	} else if err := _user.Query(user.Username); err!=nil || _user.Password != user.Password {
		c.JSON(http.StatusBadGateway,gin.H{"message":e.GetMsg(e.ERROR_PASSWORD)})
	} else {
		h := md5.New()
		io.WriteString(h,strconv.FormatInt(time.Now().Unix(),10))
		token := fmt.Sprintf("%x",h.Sum(nil))
		user.Update("token",token)
		c.JSON(http.StatusOK, gin.H{"message": "success", "token": token,"userId":_user.AuthId,"userName":_user.Username})
	}
}
