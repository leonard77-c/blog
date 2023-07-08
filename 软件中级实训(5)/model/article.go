package model

import (
	"Project/utils"
	"fmt"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Category     Category `gorm:"foreignKey:Cid"`
	User         User     `gorm:"foreignKey:Uid"`
	Title        string   `gorm:"type:varchar(100);not null" json:"title"`
	Cid          int      `gorm:"type:int;not null" json:"cid"`
	Uid          int      `gorm:"type:int;not null" json:"uid"`
	Desc         string   `gorm:"type:varchar(200)" json:"desc"`
	Content      string   `gorm:"type:longtext" json:"content"`
	Img          string   `gorm:"type:varchar(100)" json:"img"`
	CommentCount int      `gorm:"type:int;not null;default:0" json:"comment_count"`
	ReadCount    int      `gorm:"type:int;not null;default:0" json:"read_count"`
	PraiseNum    int      `gorm:"default:0" json:"praise_count"`
}

type BlogList struct {
	ArticleList []Article
	Total       int
}

// CreateArt 新增文章
func CreateArt(data *Article) int {
	//var tem1 Article
	//var tem Category
	//fmt.Println(data)
	//db.Last(&tem1)
	//data.ID = tem1.ID + 1
	//data.Category = tem.Category
	err := Db.Create(&data).Error
	fmt.Println(data)
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCSE
}

// GetArtInfo 查询单个文章
func GetArtInfo(id uint) (Article, int) {
	var art Article
	err = Db.Where("id = ?", id).Preload("Category").Preload("User").First(&art).Error
	Db.Model(&art).Where("id = ?", id).UpdateColumn("read_count", gorm.Expr("read_count + ?", 1))
	if err != nil {
		return art, errmsg.ERROR_ART_NOT_EXIST
	}
	return art, errmsg.SUCCSE
}

// SearchArt 查询相关标题文章
func SearchArt(title string) ([]Article, int, int) {
	var art []Article
	var total int
	Db.Preload("Category").Preload("User").Where("title LIKE ?", "%"+title+"%").Find(&art)
	total = len(art)
	fmt.Println(total)
	fmt.Println(art)
	if total == 0 {
		return art, errmsg.ERROR_ART_NOT_EXIST, total
	}
	return art, errmsg.SUCCSE, total

}

func SearchUserArt(username string) ([]Article, int, int) {
	var art []Article
	var total int
	var user User
	err = Db.Where("username=?", username).First(&user).Error
	//fmt.Println(err)
	//fmt.Println(user)
	var id = user.ID
	//fmt.Println(id)
	Db.Preload("Category").Preload("User").Where("uid LIKE ?", id).Find(&art)
	total = len(art)
	if total == 0 {
		return art, errmsg.ERROR_ART_NOT_EXIST, total
	}
	return art, errmsg.SUCCSE, total

}

// GetCateArt 查询分类下的所有文章
func GetCateArt(name string) ([]Article, int, int64) {
	var cateArtList []Article
	var total int64
	var cid int
	var cate Category

	Db.Where("name = ?", name).First(&cate)
	cid = int(cate.ID)
	fmt.Println(cid)
	err = Db.Preload("Category").Preload("User").Where("cid=?", cid).Find(&cateArtList).Error
	Db.Model(&cateArtList).Where("cid =?", cid).Count(&total)
	if err != nil {
		return nil, errmsg.ERROR_CATE_NOT_EXIST, 0
	}
	return cateArtList, errmsg.SUCCSE, total
}

func UpdatePraise(id uint) (int, int) {
	var art Article
	var Praise int
	Db.Where("id = ?", id).First(&art)
	art.PraiseNum = art.PraiseNum + 1
	Praise = art.PraiseNum
	result := Db.Where("id = ?", id).Updates(&art)
	if result.Error != nil {
		return Praise, errmsg.ERROR_ART_NOT_EXIST
	}
	return Praise, errmsg.SUCCSE
}

func DeleteArt(id uint) int {
	var ART Article
	fmt.Println(id)
	err = Db.Where("id = ?", id).Delete(&ART).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCSE
}

func GetArticleImg(id int) (filename string, pro Profile) {
	var user User
	Db.Where("id=?", id).First(&user)
	Db.Where("name=?", user.Username).First(&pro)
	filename = pro.Img
	return filename, pro
}
