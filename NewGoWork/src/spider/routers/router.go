package routers

import (
	"github.com/astaxie/beego"
	"spider/controllers/admin"
	"spider/controllers/api"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSRouter("/register", &api.BookController{}, "post:Register"), //注册的接口  Register这个方法名要大写 不然找不到的

		beego.NSNamespace("/book",
			// /v1/book  获取小说列表
			beego.NSRouter("/", &api.BookController{}, "get:GetAll"),
			// /v1/book/getchapters?bookid=xxx&page=1  获取小说章节列表
			beego.NSRouter("/getchapters", &api.BookController{}, "get:GetChapters"),
			// /v1/book/chapter?id=xxx  获取章节内容
			beego.NSRouter("/chapter", &api.BookController{}, "get:GetChapter"),
		),
	)
	ns.Router("/login", &api.BookController{}, "post:Login")
	beego.AddNamespace(ns)

	nsAdmin := beego.NewNamespace("/admin",
		beego.NSRouter("/", &admin.IndexController{}, "get:Index"),
		beego.NSRouter("/login", &admin.LoginController{}, "get:Login"),
		beego.NSRouter("/users", &admin.UserController{}, "get:Users"),
		beego.NSRouter("/users/edit", &admin.UserController{}, "get:Edit"),
		beego.NSRouter("/users/changepwd", &admin.UserController{}, "get:ChangePwd"),
		beego.NSRouter("/apps", &admin.AppController{}, "get:Apps"),
		beego.NSRouter("/myapps", &admin.AppController{}, "get:MyApps"),
		beego.NSRouter("/app/add", &admin.AppController{}, "get:Add"),

		beego.NSNamespace("/service",
			beego.NSRouter("/login", &admin.LoginController{}, "post:AjaxLogin"),
			beego.NSRouter("/loginout", &admin.LoginController{}, "post:AjaxLoginOut"),

			beego.NSRouter("/user_add", &admin.UserController{}, "post:AjaxAdd"),
			beego.NSRouter("/user_delete", &admin.UserController{}, "post:AjaxDelete"),
			beego.NSRouter("/user_enable", &admin.UserController{}, "post:AjaxEnable"),
			beego.NSRouter("/user_disable", &admin.UserController{}, "post:AjaxDisable"),
			beego.NSRouter("/user_edit", &admin.UserController{}, "post:AjaxEdit"),
			beego.NSRouter("/password_change", &admin.UserController{}, "post:AjaxChangePwd"),

			beego.NSRouter("/app_add", &admin.AppController{}, "post:AjaxAdd"),
			beego.NSRouter("/app_delete", &admin.AppController{}, "post:AjaxDelete"),
			beego.NSRouter("/app_pass", &admin.AppController{}, "post:AjaxPass"),
			beego.NSRouter("/app_unpass", &admin.AppController{}, "post:AjaxUnPass"),
			beego.NSRouter("/app_apply", &admin.AppController{}, "post:AjaxApply"),
		),
	)
	beego.AddNamespace(nsAdmin)
}
