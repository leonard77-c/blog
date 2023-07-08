package api

import (
	"Project/model"
	"Project/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var formData model.User
	err := c.ShouldBindJSON(&formData)
	fmt.Println(err)
	fmt.Println(formData.Password)
	var code int

	formData, code = model.CheckLogin(formData.Username, formData.Password)

	if code == errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errmsg.GetErrMsg(200),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"data":    formData.Username,
			"id":      formData.ID,
			"message": errmsg.GetErrMsg(code),
		})
	}

}
