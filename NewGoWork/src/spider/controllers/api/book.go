package api

import (
	"spider/models"
	"spider/util"
	"strconv"
	"time"
)

type BookController struct {
	ApiController
}

//获取小说列表
func (self *BookController) GetAll() {
	books, _ := models.GetBookList()
	self.ToJson(MSG_OK, "成功", books)
}

// 分页获取指定小说的章节列表, 每页10条
// url参数：bookid => 小说id; page => 页码
func (self *BookController) GetChapters() {
	bookid, err := self.GetInt("bookid")
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	page, err := self.GetInt("page")
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	chapters, _ := models.GetChapterPage(page, 10, "book_id", bookid)
	self.ToJson(MSG_OK, "success", chapters)
}

// 获取指定章节详细信息
// url参数: id => 章节id
func (self *BookController) GetChapter() {
	id, err := self.GetInt("id")
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	chapter, err := models.GetChapterById(id)
	if err != nil {
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "success", chapter)
}

//后面是我加的
func (self *BookController) Register() {
	username := self.GetString("username")
	password := self.GetString("password")
	if len(username) == 0 {
		self.ToJson(MSG_ERR, "用户名不能为空", nil)
		return
	}
	if len(password) == 0 || len(password) < 6 {
		self.ToJson(MSG_ERR, "密码长度不规范", nil)
		return
	}
	md5 := util.Md5(password, false)
	user := models.User{Username: username, Password: md5, Level: 99, Status: 1, CreatedAt: time.Now()}
	_, e := models.UserAdd(&user)
	if e != nil {
		self.ToJson(MSG_ERR, "用户名已存在", nil)
	} else {
		self.Ctx.SetCookie("auth", "ceshishiyong", 7*86400) //这个返回直接就在postman上就能看到的
		self.ToJson(MSG_OK, "注册成功", nil)
	}
}

//后面是我加的
func (self *BookController) Login() {
	username := self.GetString("username")
	password := self.GetString("password")
	password = util.Md5(password, false) //存的是32位的md5密码 前面不能有空格的切记
	user, err := models.GetUserByName(username)
	if err != nil || user == nil {
		self.ToJson(MSG_ERR, "用户不存在", nil)
	}
	if password != user.Password {
		self.ToJson(MSG_ERR, "密码不正确", nil)
	}
	if user.Status != 1 {
		self.ToJson(MSG_ERR, "该账户已被禁用", nil)
	}
	authkey := util.Md5(self.GetClientIp()+"|"+password, false)
	self.Ctx.SetCookie("auth", strconv.Itoa(user.Id)+"|"+authkey, 7*86400)
	//self.Ctx.SetCookie("auth", "hahahahahaahahahahh")
	self.ToJson(MSG_OK, "登录成功", nil)
}
