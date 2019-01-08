package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type App struct {
	Id        int
	Appname   string
	Token     string
	UserId    int
	Count     int
	Desc      string
	Status    int
	CreatedAt time.Time
}

type AppDetail struct {
	App
	Username   string
	StatusDesc string
}

func (self *App) TableName() string {
	return "app"
}

func AppAdd(app *App) (int64, error) {
	return orm.NewOrm().Insert(app)
}

func GetAppList(page, pageSize int, filters ...interface{}) ([]*App, int64) {
	offset := (page - 1) * pageSize
	list := make([]*App, 0)
	query := orm.NewOrm().QueryTable("app")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("id").Limit(pageSize, offset).All(&list)
	return list, total
}

func (app *App) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(app, fields...); err != nil {
		return err
	}
	return nil
}

func GetAppById(id int) (*App, error) {
	app := new(App)
	err := orm.NewOrm().QueryTable(app.TableName()).Filter("id", id).One(app)
	if err != nil {
		return nil, err
	}
	return app, nil
}
