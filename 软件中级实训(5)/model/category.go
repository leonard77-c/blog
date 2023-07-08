package model

import (
	errmsg "Project/utils"
)

type Category struct {
	ID   uint   `gorm:"primary_key;auto_increment" json:"id"`
	Name string `gorm:"type:varchar(20);not null" json:"name"`
}

type ListOfCate struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
}

func ReturnCate() ([50]ListOfCate, int, int) {

	var CateList []Category
	//fmt.Println("数据库查找")
	Db.Find(&CateList)
	//num := len(CateList)
	var List [50]ListOfCate
	for i := 0; i < len(CateList); i++ {
		//fmt.Println(i)
		List[i] = ListOfCate{
			CateList[i].ID,
			CateList[i].Name,
		}
	}
	//fmt.Println(2)
	return List, errmsg.SUCCSE, len(CateList)
}
