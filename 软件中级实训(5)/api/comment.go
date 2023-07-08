package api

import (
	"Project/model"
	"Project/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AddComment(c *gin.Context) {
	var data model.Comment
	//var user model.User
	type temp struct {
		Value string `json:"value"`
		Uid   int    `json:"username1"`
		Id    int    `json:"id1"`
	}
	var tem temp
	_ = c.ShouldBindJSON(&tem)
	//fmt.Println(tem)
	//model.Db.Where("username=?", tem.Username).First(&user)
	data.Uid = tem.Uid
	data.Aid = tem.Id
	data.Content = tem.Value
	code := model.CreateComment(&data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

func SearchUserComment(c *gin.Context) {
	type temp struct {
		Username string `json:"username"`
	}
	var username temp
	c.ShouldBindJSON(&username)
	var code int
	var comment []model.Comment
	var num int
	comment, code, num = model.SearchUserComment(username.Username)
	var commentList model.CommentList
	commentList = model.CommentList{
		comment,
		num,
	}
	if code == errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":      200,
			"message":     errmsg.GetErrMsg(200),
			"commentList": commentList,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}

func DeleteComment(c *gin.Context) {
	fmt.Println("删除评论")
	type temp struct {
		Id uint `json:"ID"`
	}
	var id temp
	_ = c.ShouldBindJSON(&id)
	var code int
	code = model.DeleteComment(id.Id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

func SearchArticleComment(c *gin.Context) {
	type temp struct {
		Id uint `json:"ID"`
	}
	var id temp
	c.ShouldBindJSON(&id)
	fmt.Println("comment")
	fmt.Println(id.Id)
	var code int
	var comment []model.Comment
	var num int
	comment, code, num = model.SearchArticleComment(id.Id)
	var commentList model.CommentList
	commentList = model.CommentList{
		comment,
		num,
	}
	if code == errmsg.SUCCSE {
		c.JSON(http.StatusOK, gin.H{
			"status":      200,
			"message":     errmsg.GetErrMsg(200),
			"commentList": commentList,
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		})
	}
}
