package admin

import (
	"spider/models"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/utils/pagination"
	"spider/util"
	"strconv"
	"time"
)

type UserController struct {
	AdminController
}

func (self *UserController) Users() {
	page, err := self.GetInt("p")
	if err != nil {
		page = 1
	}
	users, total := models.GetUserList(page, self.pageSize, "status__gte", 0)
	self.Data["users"] = users
	self.Data["total"] = total
	if total < 1 {
		self.Data["hasdata"] = false
	} else {
		self.Data["hasdata"] = true
	}

	paginator := pagination.SetPaginator(self.Ctx, self.pageSize, total)
	self.Data["paginator"] = paginator
	self.display()
}

func (self *UserController) Edit() {
	id, _ := self.GetInt("id")
	user, _ := models.GetUserById(id)
	self.Data["user"] = user
	self.display()
}

func (self *UserController) ChangePwd() {
	self.display()
}

func (self *UserController) AjaxAdd() {
	username := self.GetString("username")
	if len(username) < 1 {
		self.ToJson(MSG_ERR, "The username can not be empty", nil)
	}
	user := new(models.User)
	user.Username = username
	user.Password = util.Md5("123456", false)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Level = 99
	user.Status = 0
	id, err := models.UserAdd(user)
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	if id < 1 {
		self.ToJson(MSG_ERR, "新增用户失败，换个用户名再试", nil)
	}
	self.ToJson(MSG_OK, "新增用户成功", nil)
}

func (self *UserController) AjaxDelete() {
	id, err := self.GetInt("id")
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	user, err := models.GetUserById(id)
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	user.Status = -1
	user.UpdatedAt = time.Now()
	err = user.Update()
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "用户【"+user.Username+"】已删除", nil)
}

func (self *UserController) AjaxEnable() {
	id, err := self.GetInt("id")
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	user, err := models.GetUserById(id)
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	user.Status = 1
	user.UpdatedAt = time.Now()
	err = user.Update()
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "用户【"+user.Username+"】已启用", nil)
}

func (self *UserController) AjaxDisable() {
	id, err := self.GetInt("id")
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	user, err := models.GetUserById(id)
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	user.Status = 0
	user.UpdatedAt = time.Now()
	err = user.Update()
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "用户【"+user.Username+"】已禁用", nil)
}

func (self *UserController) AjaxEdit() {
	id, err := self.GetInt("id")
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	level, err := self.GetInt("level")
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	user, err := models.GetUserById(id)
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	user.Level = level
	user.UpdatedAt = time.Now()
	err = user.Update()
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "用户【"+user.Username+"】已修改成功", nil)
}

func (self *UserController) AjaxChangePwd() {
	old_pwd := self.GetString("old_pwd")
	new_pwd := self.GetString("new_pwd")
	con_pwd := self.GetString("con_pwd")
	logs.Info(self.login_user)
	logs.Info(self.login_userId)
	user, _ := models.GetUserById(self.login_user.Id)

	old_pwd = util.Md5(old_pwd, false)
	if old_pwd != user.Password {
		self.ToJson(MSG_ERR, "旧密码不正确", nil)
	}
	if new_pwd != con_pwd {
		self.ToJson(MSG_ERR, "两次输入的新密码不一致", nil)
	}
	user.Password = util.Md5(new_pwd, false)
	err := user.Update()
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}

	self.login_user = user
	authkey := util.Md5(self.GetClientIp()+"|"+user.Password, false)
	self.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
	self.ToJson(MSG_OK, "用户【"+user.Username+"】的密码已修改成功", nil)
}


