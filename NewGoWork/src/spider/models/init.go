package models

import(
	"fmt"
	"spider/config"
	"net/url"
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
)

func init(){
	dbhost := config.AppConfig.GetString("db.host")
	dbport := config.AppConfig.GetString("db.port")
	dbuser := config.AppConfig.GetString("db.user")
	dbpwd := config.AppConfig.GetString("db.password")
	dbname := config.AppConfig.GetString("db.name")
	timezone := config.AppConfig.GetString("db.timezone")
	if dbport == ""{
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpwd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	if timezone != "" {
		dsn = dsn + "&loc=" + url.QueryEscape(timezone)
	}
	base := orm.RegisterDataBase("default", "mysql", dsn)
	fmt.Println("RegisterDataBase",base)
	orm.RegisterModel(new(Book), new(Chapter), new(Url), new(User), new(App))
}

func TableName(name string) string {
	return config.AppConfig.GetString("db.prefix") + name
}