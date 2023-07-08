package api

import (
	"Project/model"
	errmsg "Project/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Return_Cate(c *gin.Context) {
	var cate [50]model.ListOfCate
	//fmt.Println(1)
	var code int
	var num int
	cate, code, num = model.ReturnCate()
	//fmt.Println(cate)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
		"cate":    cate[0:num],
	})

}
