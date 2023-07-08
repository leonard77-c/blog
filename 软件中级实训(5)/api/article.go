package api

import (
	"Project/model"
	"Project/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddArticle 添加文章
func AddArticle(c *gin.Context) {
	var data model.Article
	_ = c.ShouldBindJSON(&data)

	code := model.CreateArt(&data)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetArtInfo 通过ID查询单个文章信息
func GetArtInfo(c *gin.Context) {
	type temp struct {
		Id uint `json:"id"`
	}
	var tem temp
	c.ShouldBindJSON(&tem)
	//fmt.Print(tem)
	data, code := model.GetArtInfo(tem.Id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// SearchArt 通过一个标题进行模糊查询
func SearchArt(c *gin.Context) {
	var blog model.BlogList
	type temp struct {
		Title string `json:"title"`
	}
	var tem temp
	c.ShouldBindJSON(&tem)
	var code int
	var art []model.Article
	var num int
	art, code, num = model.SearchArt(tem.Title)
	blog = model.BlogList{
		ArticleList: art,
		Total:       num,
	}
	fmt.Println(blog)
	if code == errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":   200,
			"message":  errmsg.GetErrMsg(200),
			"blogList": blog,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}

}

// SearchUserArt 查找用户的文章
func SearchUserArt(c *gin.Context) {
	type temp struct {
		Username string `json:"username"`
	}
	var username temp
	c.ShouldBindJSON(&username)
	fmt.Print("进入用户文章查询")
	//userName := c.Query("username")
	var code int
	var art []model.Article
	var num int
	fmt.Println(username.Username)
	var blogList model.BlogList
	art, code, num = model.SearchUserArt(username.Username)
	blogList = model.BlogList{
		art,
		num,
	}
	if code == errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":   200,
			"message":  errmsg.GetErrMsg(200),
			"blogList": blogList,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(c *gin.Context) {

	type temp struct {
		Name string `json:"name"`
	}
	var tem temp
	c.ShouldBindJSON(&tem)

	fmt.Println(tem.Name)
	data, code, total := model.GetCateArt(tem.Name)

	var blogList model.BlogList
	blogList = model.BlogList{
		ArticleList: data,
		Total:       int(total),
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   code,
		"blogList": blogList,
		"total":    total,
		"message":  errmsg.GetErrMsg(code),
	})
}

func UpdatePraise(c *gin.Context) {
	type temp struct {
		ID uint `json:"uid"`
	}
	var tem temp
	fmt.Println(tem)
	c.ShouldBindJSON(&tem)
	var praise_num int
	var code int
	praise_num, code = model.UpdatePraise(tem.ID)
	if code == errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":       200,
			"message":      errmsg.GetErrMsg(200),
			"praise_count": praise_num,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

func DeleteArticle(c *gin.Context) {
	type temp struct {
		AID uint `json:"aid"`
	}
	var tem temp
	c.ShouldBindJSON(&tem)
	code := model.DeleteArt(tem.AID)
	c.JSON(200, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// 获取文章图片
func GetArtImg(c *gin.Context) {
	//type temp struct {
	//	Id string `json:"id"`
	//}
	//var tem temp
	//c.ShouldBindJSON(&tem)
	//id, _ := strconv.Atoi(tem.Id)
	//filename, _ := model.GetPro(id)
	//c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	//c.File("./" + filename)

	//c.JSON(200, gin.H{
	//	"data": pro,
	//	"code": 200,
	//})
}
