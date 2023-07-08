package model

import (
	"Project/utils"
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);not null " json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type:varchar(500);not null" json:"password" validate:"required,min=6,max=120" label:"密码"`
}

// CheckUser 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	Db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED //1001
	}
	return errmsg.SUCCSE
}

// CheckLogin 后台登录验证
func CheckLogin(username string, password string) (User, int) {
	var user User
	//var PasswordErr error

	Db.Where("BINARY username = ?", username).First(&user)

	//PasswordErr = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if user.ID == 0 {
		return user, errmsg.ERROR_USER_NOT_EXIST
	}
	if user.Password != password {
		return user, errmsg.ERROR_PASSWORD_WRONG
	}
	//if PasswordErr != nil {
	//	return user, errmsg.ERROR_PASSWORD_WRONG
	//}
	//if user.Role != 1 {
	//	return user, errmsg.ERROR_USER_NO_RIGHT
	//}
	return user, errmsg.SUCCSE
}

func DoRegister(password string, username string) int {
	var user User
	Db.Where("BINARY username = ?", username).First(&user)
	if user.ID != 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	var user1 User
	var pro Profile
	Db.Last(&user1)
	var Id uint
	Id = user1.ID + 1
	user2 := User{Model: gorm.Model{ID: Id, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Username: username, Password: password}
	Db.Create(&user2)
	pro.ID = int(user.ID)
	pro.Name = username
	Db.Create(&pro)
	return errmsg.SUCCSE

}

// 发送邮件
func EmailSend(addr string) string {
	m := gomail.NewMessage(
		gomail.SetEncoding(gomail.Base64),
	)
	userName := "1154191642@qq.com"
	code := Create_Code()
	//message := "诞生于1996，梦想成为说唱领袖，so Pretty so Jiggy,没有太多讲究。爱唱跳舞的男孩，创作是他的王牌，随时随地可以表演，生活就是他的舞台大家的目光就像是我的兴奋剂大家好，我是来自BBT的 王 子 异"
	message := `
	<p> Hello %s,</p>

		<p style="text-indent:2em">The special code for you is</p>
		
		<p style="text-indent:2em">

	` + code

	m.SetHeader("From", "Hello!My friend"+"<"+userName+">")
	mailTo := []string{
		//"1642213006@qq.com",
		addr,
	}
	m.SetHeader("To", mailTo...)
	m.SetBody("text/html", fmt.Sprintf(message, "IKUN"))
	m.SetHeader("Subject", "click here and find your code")
	host := "smtp.qq.com"
	port := 25
	password := "tsdaitroaxqlgeii" // qq邮箱填授权码
	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	m.Attach("Kiss.jpg")
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
	return code

}

// 创建验证码
func Create_Code() string {
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	vcode := fmt.Sprintf("%06v", rnd.Int31n(1000000))
	return vcode
}

// 修改密码
func CodeChange(id uint, passw string) int {
	var user User
	err = Db.Model(&user).Where("id = ?", id).UpdateColumn("password", passw).Error
	if err != nil {
		return errmsg.ERROR
	}
	return 200
}
