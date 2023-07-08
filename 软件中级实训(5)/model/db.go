package model

import (
	_ "Project/utils"
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
	_ "gorm.io/gorm"
	_ "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	_ "gorm.io/gorm/schema"
	"os"
)

var Db *gorm.DB
var err error

func InitDb() {
	dns := "root:123456@(127.0.0.1:3306)/final?charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai"
	Db, err = gorm.Open(mysql.Open(dns), &gorm.Config{
		// gorm日志模式：silent
		//Logger: logger.Default.LogMode(logger.Silent),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	//db, err = gorm.Open("mysql", "root:Xi1154191642@(127.0.0.1:3306)/new?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("连接数据库失败，请检查参数：", err)
		os.Exit(1)
	}
	_ = Db.AutoMigrate(&User{}, &Article{})
}
