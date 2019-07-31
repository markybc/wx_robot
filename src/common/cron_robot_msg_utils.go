package common

import (
	"github.com/astaxie/beego/httplib"
	"github.com/robfig/cron"
	"log"
	"models"
	"github.com/astaxie/beego/logs"
	"ni.com/commonutils"
)

var CRON_JOB = cron.New()


func StartAllJob(){
	InitDbJob()
	CRON_JOB.Start()
}

func InitDbJob() {
	var list []models.WxRobotMsg
	models.WxRobotMsgList(&list)
	for i := 0; i < len(list); i++ {
		AddJob(list[i].CronTab, list[i].Id)
	}

}

func AddJob(crontab string, id int64) {
	s, err := cron.Parse(crontab)
	if err != nil {
		log.Println("Parse error",commonutils.ToStr(id))
	}
	h := SendRobotMarkdownMsg{Id: id}
	logs.Info("add job id :" + commonutils.ToStr(id))
	CRON_JOB.Schedule(s, h)
}

func (c SendRobotMarkdownMsg) Run() {
	var wxMsg = models.WxRobotMsg{}
	models.GetWxRobotMsg(&wxMsg,c.Id)
	if wxMsg.Status == 1 {
		logs.Info("run job :" , wxMsg)
		req := httplib.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=" + wxMsg.Key)

		if wxMsg.MsgType == "markdown" {
			msgMap := map[string]interface{}{"msgtype": "markdown", "markdown": map[string]string{"content": wxMsg.Msg}}
			logs.Info(msgMap)
			req.JSONBody(msgMap)
		}else if wxMsg.MsgType == "text" {
			msgMap := map[string]interface{}{"msgtype": "text", "text": map[string]string{"content": wxMsg.Msg, "mentioned_list": "@all"}}
			logs.Info(msgMap)
			req.JSONBody(msgMap)
		}else {
			return
		}

		resultResp, _ := req.Response()
		if resultResp.StatusCode == 200 {
			logs.Info("run job success : " , wxMsg)
		}
	}else {
		logs.Info("no run job :" , wxMsg)
	}
}

type SendRobotMarkdownMsg struct {
	Id     int64
}
