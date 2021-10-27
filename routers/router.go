package routers

import (
	"channelServe/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:CreateChanel")
	beego.Router("/player", &controllers.MainController{}, "get:Show;post:PlayerPost")
}
