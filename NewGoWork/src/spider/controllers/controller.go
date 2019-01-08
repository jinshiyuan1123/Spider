package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type BaseController struct { //全部是继承  子类调用父类所有方法 beego.Controller
	beego.Controller
}

//{{$timestamp}}  可以给postman动态设置请求时间戳
//如果传入的时间戳小于服务器当前时间  返回false  提示权限不足
//
//如果传入的时间戳大于服务器当前时间  返回true  可以正常访问
func (self *BaseController) Prepare() { //不用这个方法了
	self.Ctx.Request.Cookies()
	//headers := self.Ctx.Request.Header //这个只是我模拟的 对所有请求做拦截
	//time_str := headers.Get("time")//请求头中的信息 获取是不区分大小写的 传过来就是大了 但是我们能用小写获取到

	// int float string 互相转换 必须是同一个类型才能互相加减
	//i_time, _ := strconv.ParseFloat(time_str,64)//整数型的时间戳了
	//registerTime := time.Now().Unix()
	//s := string(registerTime)
	//f, _ := strconv.ParseFloat(s, 64)

	//if time.Now().After(registerTime.AddDate(1, 0, 0)) {
	//	fmt.Println("已经超过一年")
	//}
	//s := self.GetString("token")
	//fmt.Println("s=",s)
	//if len(s)==0 {
	//	self.ToJson(config.MSG_ERR,"请先登录",nil)
	//}
}

// 固定返回的json数据格式
// msgno: 错误码
// msg: 错误信息
// data: 返回数据
func (self *BaseController) ToJson(msgno int, msg string, data interface{}) {
	out := make(map[string]interface{})
	out["status"] = msgno
	out["msg"] = msg
	out["data"] = data
	self.Data["json"] = out
	self.ServeJSON()
}

//获取用户IP地址
func (self *BaseController) GetClientIp() string {
	s := strings.Split(self.Ctx.Request.RemoteAddr, ":")
	return s[0]
}
