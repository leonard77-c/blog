package api

import (
	"Project/model"
	errmsg "Project/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context) {
	var formData model.User
	err := c.ShouldBindJSON(&formData)
	fmt.Println(err)
	fmt.Println(formData.Password)
	var code int
	code = model.DoRegister(formData.Password, formData.Username)
	if code == errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":  200,
			"message": errmsg.GetErrMsg(200),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// 邮件控制
func Email_Control(c *gin.Context) {
	type temp struct {
		Email string
	}
	var tem temp
	c.ShouldBindJSON(&tem)
	code := model.EmailSend(tem.Email)
	c.JSON(200, gin.H{
		"code":   code,
		"status": 200,
	})
}
