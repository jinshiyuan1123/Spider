package admin

import(
	"spider/util"
	"time"
	"github.com/astaxie/beego/utils/pagination"
	"spider/models"
)

type AppController struct{
	AdminController
}

func (self *AppController)Apps(){
	page, err := self.GetInt("p")
	if err != nil{
		page = 1
	}
	apps, total := models.GetAppList(page, self.pageSize, "status__gte", 0)
	appDetails := make([]*models.AppDetail, 0)
	for _, app := range apps{
		user, _ := models.GetUserById(app.UserId)
		ad := new(models.AppDetail)
		ad.Username = user.Username
		ad.Id = app.Id
		ad.Appname = app.Appname
		ad.Token = app.Token
		ad.UserId = app.UserId
		ad.Status = app.Status
		ad.Count = app.Count
		ad.CreatedAt = app.CreatedAt
		ad.Desc = app.Desc
		appDetails = append(appDetails, ad)
	}
	self.Data["apps"] = appDetails
	self.Data["total"] = total
	if total < 1{
		self.Data["hasdata"] = false
	}else{
		self.Data["hasdata"] = true
	}

	paginator := pagination.SetPaginator(self.Ctx, self.pageSize, total)
	self.Data["paginator"] = paginator
	self.display()
	self.display()
}

func (self *AppController) MyApps(){
	page, err := self.GetInt("p")
	if err != nil{
		page = 1
	}
	apps, total := models.GetAppList(page, self.pageSize, "status__gte", 0, "user_id", self.login_userId)
	self.Data["apps"] = apps
	self.Data["total"] = total
	if total < 1{
		self.Data["hasdata"] = false
	}else{
		self.Data["hasdata"] = true
	}

	paginator := pagination.SetPaginator(self.Ctx, self.pageSize, total)
	self.Data["paginator"] = paginator
	self.display()
}

func (self *AppController)Add(){
	self.display()
}

func (self *AppController)AjaxAdd(){
	appname := self.GetString("appname")
	desc := self.GetString("desc")

	app := new(models.App)
	app.Appname = appname
	app.Desc = desc
	app.Count = 1000
	app.CreatedAt = time.Now()
	app.Status = 0
	app.Token = util.Md5(appname + app.CreatedAt.String(), false)
	app.UserId = self.login_userId

	_, err := models.AppAdd(app)
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "新增应用成功", nil)
}

func (self *AppController)AjaxPass(){
	id, err := self.GetInt("id")
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	app, err := models.GetAppById(id)
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	app.Status = 1
	err = app.Update()
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "应用【"+app.Appname+"】已审核通过", nil)
}

func (self *AppController)AjaxUnPass(){
	id, err := self.GetInt("id")
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	app, err := models.GetAppById(id)
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	app.Status = 2
	err = app.Update()
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "应用【"+app.Appname+"】审核不通过", nil)
}

func (self *AppController)AjaxDelete(){
	id, err := self.GetInt("id")
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	app, err := models.GetAppById(id)
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	app.Status = -1
	err = app.Update()
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "应用【"+app.Appname+"】已删除成功", nil)
}

func (self *AppController)AjaxApply(){
	id, err := self.GetInt("id")
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	app, err := models.GetAppById(id)
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	app.Status = 0
	err = app.Update()
	if err != nil{
		self.ToJson(MSG_ERR, err.Error(), nil)
	}
	self.ToJson(MSG_OK, "应用【"+app.Appname+"】已申请成功，等待审核", nil)
}