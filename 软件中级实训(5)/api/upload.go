package api

import (
	"Project/model"
	errmsg "Project/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// UpLoad 上传图片接口
func UpLoad(c *gin.Context) {
	//读取文件
	f, err := c.FormFile("file")
	id := c.PostForm("id")
	iid, _ := strconv.Atoi(id)
	if err != nil {
		c.JSON(200, gin.H{
			"error": err.Error(),
		})
	} else {
		dst := fmt.Sprintf("./%s", f.Filename)
		//dst = path.Join("./", f.Filename) //或者
		c.SaveUploadedFile(f, dst)
	}
	code := model.SaveImg(f.Filename, iid)

	//保存在服务端本地
	filename := f.Filename
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"message":  errmsg.GetErrMsg(code),
		"filename": filename,
	})
}
