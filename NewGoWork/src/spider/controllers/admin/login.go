package admin

import(
	"spider/util"
	"strconv"
	"spider/models"
)

type LoginController struct{
	AdminController
}

func (self *LoginController) Login(){
	self.TplName = "login/login.html"
}

func (self *LoginController) AjaxLogin(){
	username := self.GetString("username")
	password := self.GetString("password")
	password = util.Md5(password, false)//存的是32位的md5密码 前面不能有空格的切记
	user, err := models.GetUserByName(username)
	if err != nil || user == nil{
		self.ToJson(MSG_ERR,"用户不存在", nil)
	}
	if password != user.Password{
		self.ToJson(MSG_ERR,"密码不正确", nil)
	}
	if user.Status != 1{
		self.ToJson(MSG_ERR,"该账户已被禁用", nil)
	}
	authkey := util.Md5(self.GetClientIp() + "|" + password, false)
	self.Ctx.SetCookie("auth", strconv.Itoa(user.Id) + "|" + authkey, 7 * 86400)
	self.ToJson(MSG_OK,"登录成功", nil)
}

//登出
func (self *LoginController) AjaxLoginOut() {
	self.Ctx.SetCookie("auth", "")
	self.ToJson(MSG_OK,"退出成功", nil)
}