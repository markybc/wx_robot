package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type WxRobotMsg struct {
	Id      int64
	Key     string
	Msg     string
	CronTab string
	Status  int8
	MsgType string
}

func WxRobotMsgList(list *[]WxRobotMsg) {
	qs := orm.NewOrm().QueryTable("wx_robot_msg")
	_, err := qs.All(list)
	if err != nil {
		logs.Info(err)
	}
}

func WxRobotMsgSave(w *WxRobotMsg) {
	orm.NewOrm().Insert(w)
}

func GetWxRobotMsg(msg *WxRobotMsg, id int64) {
	err := orm.NewOrm().QueryTable("wx_robot_msg").Filter("id", id).One(msg)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		logs.Info("Returned Multi Rows Not One")
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		logs.Info("Not row found")
	}
}
