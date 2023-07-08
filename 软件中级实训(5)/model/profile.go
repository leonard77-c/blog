package model

import (
	errmsg "Project/utils"
	"fmt"
)

type Profile struct {
	ID        int    `gorm:"primaryKey" json:"id"`
	Name      string `gorm:"type:varchar(20)" json:"name"`
	Desc      string `gorm:"type:varchar(200)" json:"desc"`
	Qqchat    string `gorm:"type:varchar(200)" json:"qq_chat"`
	Wechat    string `gorm:"type:varchar(100)" json:"wechat"`
	Weibo     string `gorm:"type:varchar(200)" json:"weibo"`
	Bili      string `gorm:"type:varchar(200)" json:"bili"`
	Email     string `gorm:"type:varchar(200)" json:"email"`
	Img       string `gorm:"type:varchar(200)" json:"img"`
	Avatar    string `gorm:"type:varchar(200)" json:"avatar"`
	IcpRecord string `gorm:"type:varchar(200)" json:"icp_record"`
}

func GetPro(username string) (filename string, pro Profile) {
	//Db.Where("id=?", id).First(&user)
	Db.Where("name=?", username).First(&pro)
	//filename = pro.Img
	return filename, pro
}

func UpdateProfile(data Profile) (code int, pro Profile) {
	err = Db.Model(&pro).Where("name = ?", data.Name).UpdateColumns(data).Error
	fmt.Println(data)
	if err != nil {
		return errmsg.ERROR, pro
	}
	return 200, pro
}
