package routes

import (
	"Project/api"
	"Project/middleware"
	"Project/model"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.Use(gin.Recovery())
	model.InitDb()
	router := r.Group("api")
	{
		router.POST("/login", api.Login)

	}

	admin := r.Group("api")
	{
		admin.POST("/article/add", api.AddArticle)                            //发布文章
		admin.POST("/article/praise", api.UpdatePraise)                       //更新点赞数
		admin.POST("/comment/add", api.AddComment)                            //发表评论
		admin.POST("/comment/delete", api.DeleteComment)                      //删除评论
		admin.POST("/register", api.Register)                                 //注册
		admin.POST("/article/search", api.SearchArt)                          //根据标题模糊查找笔记
		admin.POST("/return_cate", api.Return_Cate)                           //返回标签
		admin.POST("/admin/article/search", api.SearchUserArt)                //根据用户查找笔记
		admin.POST("/admin/comment/search", api.SearchUserComment)            //根据用户查找其评论
		admin.POST("/admin/article/info", api.GetArtInfo)                     //获取文章详细信息
		admin.POST("/cate/article/search", api.GetCateArt)                    //根据类型查询文章
		admin.POST("/email", api.Email_Control)                               //邮件
		admin.POST("/article/delete", api.DeleteArticle)                      //删除文章
		admin.POST("/admin/article/comment/search", api.SearchArticleComment) //查询文章评论
		admin.POST("/upload", api.UpLoad)                                     //上传图片
		admin.POST("/getimg", api.GetImg)                                     //获取头像
		admin.POST("/getpro", api.GetPro)                                     //获取个人信息
		admin.POST("/updateinfo", api.UpdatePro)                              //更新个人信息
		admin.POST("/password_change", api.CodeCge)                           //改密码
	}

	r.Run(":9090")
}
