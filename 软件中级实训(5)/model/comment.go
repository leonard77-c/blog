package model

import (
	errmsg "Project/utils"
	"fmt"
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	User    User    `gorm:"foreignKey:Uid"`
	Article Article `gorm:"foreignKey:Aid"`
	Uid     int     `json:"uid"`
	Aid     int     `json:"aid"`
	Content string  `gorm:"type:varchar(500);not null;" json:"content"`
	Status  int8    `gorm:"type:tinyint;default:2" json:"status"`
}

type CommentList struct {
	Comment []Comment
	Total   int
}

// 发表评论
func CreateComment(data *Comment) int {
	err := Db.Create(&data).Error
	fmt.Println(data)
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

func SearchUserComment(username string) ([]Comment, int, int) {
	var comment []Comment
	var total int
	var user User
	err = Db.Where("username=?", username).First(&user).Error
	var id = user.ID
	Db.Preload("Article").Preload("User").Where("uid = ?", id).Find(&comment)
	total = len(comment)
	if total == 0 {
		return comment, errmsg.ERROR_USER_NO_Comment, total
	}
	return comment, errmsg.SUCCSE, total
}

func DeleteComment(id uint) int {
	var comment Comment
	err = Db.Where("id = ?", id).Delete(&comment).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func SearchArticleComment(id uint) ([]Comment, int, int) {
	var comment []Comment
	var total int
	Db.Preload("Article").Preload("User").Where("aid = ?", id).Find(&comment)
	total = len(comment)
	if total == 0 {
		return comment, errmsg.ERROR_USER_NO_Comment, total
	}
	return comment, errmsg.SUCCSE, total
}
