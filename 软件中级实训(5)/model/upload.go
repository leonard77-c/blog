package model

import (
	"mime/multipart"
)

// UpLoadFile 上传文件函数
func UpLoadFile(file multipart.File, fileSize int64) {

}
func SaveImg(name string, id int) (code int) {
	var user User
	var pro Profile
	Db.Where("id =?", id).First(&user)
	Db.Where("name=?", user.Username).First(&pro)
	Db.Model(&pro).Where("name=?", user.Username).UpdateColumn("img", name)
	return 200
}
