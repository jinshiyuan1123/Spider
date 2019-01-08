package main

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"spider/config"
	_ "spider/models"
	_ "spider/routers"
	"spider/spider"
)

func main() {
	//ilog.Info("start") .
	beego.BConfig.WebConfig.Session.SessionOn = true //不开启下面调用出错 空指针
	//这个地方登陆拦截
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("auth").(string)
		if !ok && ctx.Request.RequestURI != "/v1/login" && ctx.Request.RequestURI != "/v1/register" {
			//time := ctx.Request.Header.Get("time")

			resopnse := config.Resopnse{config.MSG_ERR, "未登录", nil}
			bytes, _ := json.Marshal(resopnse)
			ctx.ResponseWriter.Write(bytes)
		}
	}

	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)
	//s, err := spider.NewSpider("booktxt")
	//if err != nil{
	//	ilog.Fatal("new Spider error: ", err.Error())
	//}
	//err = s.SpiderUrl("http://www.booktxt.net")
	//if err != nil{
	//	ilog.Fatal("new Document error: ", err.Error())
	//}
	go spider.Start()
	s := beego.VERSION
	fmt.Println(s)
	beego.Run(":8089")

}
