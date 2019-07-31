package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["controllers:WxRobotMsgController"] = append(beego.GlobalControllerRouter["controllers:WxRobotMsgController"],
		beego.ControllerComments{
			Method: "AddRobotMsgJob",
			Router: `/wx/robot/msg/add`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

}
