package main

import (
	"channelServe/models"
	_ "channelServe/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 注册数据库
	models.InitDB()
}
func main() {
	beego.Run()
}
