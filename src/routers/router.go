package routers

import (
	"controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&controllers.WxRobotMsgController{})
	beego.Router("/", &controllers.MainController{})

}
