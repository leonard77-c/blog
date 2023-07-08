登录
    url:api/login
    interface:
    {
        "username"      string
        "password"      string
    }
    return:
    {
        "status":       int     ERROR_USER_NOT_EXIST(1003用户不存在)    ERROR_PASSWORD_WRONG(1002密码错误)    SUCCSE(200成功)
        "data":         data.Username,
        "id":           data.ID,
        "message":      errmsg.GetErrMsg(code),
    }


发布文章
    url:api/article/add
    interface:
    {
        category:{
                    "id"        uint               //类型id
                    "name"      string             //直接赋值为空
                }
        "title"                 string
        "desc"                  string
        "content"               string
        "img"                   string                    //图片url
        "comment_count"         int
        "read_count"            int
    }
    return:
    {
        "status":  code     ERROR(500)      SUCCSE(200)
        "data":    data
        "message": errmsg.GetErrMsg(code)
    }


注册
    url:api/register
    interface:
    {
        "username"      string
        "password"      string
    }
    return:
    {
        "status":       int     ERROR_USERNAME_USED(1004用户已存在)        SUCCSE(200成功)
        "message":      errmsg.GetErrMsg(code),
    }

搜索文章标题
    url:api/article/search?title="title"
    return:
    {
        "status":       int     ERROR_USERNAME_USED(2001文章不存在)        SUCCSE(200成功)
        "message":      errmsg.GetErrMsg(code),
        "artList":      []art
    }