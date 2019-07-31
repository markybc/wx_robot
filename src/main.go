package main

import (
	"common"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"models"
	"os"
	_ "routers"
)

func main() {
	os.MkdirAll("/data/go/logs", os.ModePerm)
	logFilePath := "/data/go/logs/catalina.out"
	if !common.CheckFileIsExist(logFilePath) {
		os.Create("/data/go/logs/catalina.out")
	}
	//日志处理，一次引入，全局使用，无需重新引入
	logs.SetLogger(logs.AdapterFile, `{"filename":"/data/go/logs/catalina.out","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	//输出log时能显示输出文件名和行号（非必须）
	logs.EnableFuncCallDepth(true)
	os.Chmod("/data/go/logs/catalina.out", os.ModePerm)
	InitDb()
	logs.Info("wx_robot_msg is run ")

	common.StartAllJob()
	beego.Run()

}

func InitDb() {
	//dbhost := beego.AppConfig.String("db.host")
	//dbport := beego.AppConfig.String("db.port")
	//dbuser := beego.AppConfig.String("db.user")
	//dbpasswd := beego.AppConfig.String("db.passwd")
	//dbname := beego.AppConfig.String("db.name")

	dbhost := "192.168.66.94"
	dbport := "5432"
	dbuser := "trz_admin"
	dbpasswd := "reqKbs857sd"
	dbname := "trz"
	dataSource := "user=" + dbuser + " password=" + dbpasswd + " dbname=" + dbname + " host=" + dbhost + " port=" + dbport + " sslmode=disable"

	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", dataSource)
	orm.RegisterModel(new(models.WxRobotMsg))
}


