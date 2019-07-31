package controllers

import (
	"common"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"models"
)

type WxRobotMsgController struct {
	beego.Controller
}

// @router /wx/robot/msg/add [post]
func (c *WxRobotMsgController) AddRobotMsgJob() {
	logs.Info("================")
	resultMap := make(map[string]interface{})
	resultMap["message"] = "添加成功！"
	resultMap["code"] = "200"
	cronTab := c.GetString("cron")
	key := c.GetString("key")
	msg := c.GetString("msg")
	msgType := c.GetString("msgType")
	status, err := c.GetInt8("status")
	if err != nil {
		status = 0;
	}
	if msgType == "" {
		msgType = "text"
	}
	//将新增的job消息保存数据库中
	robotMsg := models.WxRobotMsg{CronTab:cronTab, Key:key, Msg:msg, MsgType:msgType, Status:status}
	models.WxRobotMsgSave(&robotMsg)
	logs.Info(robotMsg)

	//添加job
	common.AddJob(robotMsg.CronTab, robotMsg.Id)

	c.Data["json"] = resultMap
	c.ServeJSON()
	c.StopRun()
	return
}

func addNewJob(robotMsg models.WxRobotMsg){
	common.AddJob(robotMsg.CronTab, robotMsg.Id)
}