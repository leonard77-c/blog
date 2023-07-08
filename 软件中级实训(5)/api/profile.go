package api

import (
	model "Project/model"
	errmsg "Project/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

// 获取用户头像
func GetImg(c *gin.Context) {
	type temp struct {
		Id string `json:"id"`
	}
	var tem temp
	c.ShouldBindJSON(&tem)
	//id, _ := strconv.Atoi(tem.Id)
	//filename, _ := model.GetPro(id)
	//c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	//c.File("./" + filename)

	//c.JSON(200, gin.H{
	//	"data": pro,
	//	"code": 200,
	//})
}

// 获取用户个人信息
func GetPro(c *gin.Context) {
	type temp struct {
		Username string `json:"username"`
	}
	var tem temp
	c.ShouldBindJSON(&tem)
	fmt.Println(tem)
	//id, _ := strconv.Atoi(tem.Id)
	_, pro := model.GetPro(tem.Username)

	c.JSON(200, gin.H{
		"data":     pro,
		"code":     200,
		"password": "",
	})
}

// 更新个人信息
func UpdatePro(c *gin.Context) {
	var tem model.Profile
	c.ShouldBindJSON(&tem)
	fmt.Println(tem)
	code, pro := model.UpdateProfile(tem)

	c.JSON(200, gin.H{
		"data":    pro,
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 修改密码
func CodeCge(c *gin.Context) {
	type temp struct {
		ID       uint   `json:"id"`
		Password string `json:"password"`
	}
	var tem temp
	c.ShouldBindJSON(&tem)
	code := model.CodeChange(tem.ID, tem.Password)

	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
