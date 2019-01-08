package api

import(
	"spider/controllers"
)

// json 返回错误码
const (
	MSG_OK  = 0   // 成功
	MSG_ERR = -1  // 失败
)

// 基类
type ApiController struct{
	controllers.BaseController
}

