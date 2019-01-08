package admin

import(
	"spider/util"
	"strings"
	"strconv"
	"spider/controllers"
	"spider/models"

	
)

// json 返回错误码
const (
	MSG_OK  = 0   // 成功
	MSG_ERR = -1  // 失败
)

type AdminController struct{
	controllers.BaseController
	controllerName string
	actionName string
	pageSize int
	login_userId int
	login_user *models.User
}

func (self *AdminController) Prepare(){
	self.pageSize = 10
	controllerName, actionName := self.GetControllerAndAction()
	self.controllerName = strings.ToLower(controllerName[0 : len(controllerName)-10])
	self.actionName = strings.ToLower(actionName)
	self.Data["cur_controller"] = self.controllerName
	self.Data["cur_action"] = self.actionName

	if self.auth(){
        self.Data["login_username"] = self.login_user.Username
	}
}

func (self *AdminController) auth() bool{
	auth := self.Ctx.GetCookie("auth")
	arr := strings.Split(auth, "|")
	self.login_userId = 0
	if len(arr) == 2{
		idStr, authkey := arr[0], arr[1]
		self.login_userId, _ = strconv.Atoi(idStr)
		if self.login_userId > 0{
			user,err := models.GetUserById(self.login_userId)
			if err == nil && authkey == util.Md5(self.GetClientIp() + "|" + user.Password, false){
				self.login_user = user
				return true
			}else{
				self.login_userId = 0
			}
		}
	}
	if self.login_userId < 1 && (self.controllerName != "login" && self.actionName != "loginin"){
		self.redirect("/admin/login")
	}
	return false
}

func (self *AdminController) display(tpl ...string){
	var tplname string
	if len(tpl) > 0 {
		tplname = strings.Join([]string{tpl[0], "html"}, ".")
	} else {
		tplname = self.controllerName + "/" + self.actionName + ".html"
	}
	self.Layout = "public/layout.html"
	self.TplName = tplname
}

// 重定向
func (self *AdminController) redirect(url string) {
	self.Redirect(url, 302)
	self.StopRun()
}