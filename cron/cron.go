package cron

import (
	"flag"
	"fmt"
	"github.com/robfig/cron/v3"
	log "github.com/sirupsen/logrus"
	"indicator/config"
	"indicator/service"
	"net/url"
)

var configFile = flag.String("config", "app.yaml", "配置文件路径")

var conf = config.Init(*configFile)

func StartSchedule() {
	log.Info("Indicator server started successfully~")
	crontab := cron.New(cron.WithSeconds())
	// 添加定时任务
	_, _ = crontab.AddFunc(conf.Cron, Run)
	// 启动定时器
	crontab.Start()
	select {}
}

func Run() {
	for i := range conf.Applicant {
		a := conf.Applicant[i]
		res := service.Congratulate(a.Code, a.Name)
		var msg = ""
		if res {
			msg = "%s, 你中签了！！！"
		} else {
			msg = "%s, 这次又没有中啊"
		}
		content := fmt.Sprintf(msg, a.Nickname)
		//推送消息到微信
		result := service.SendWechatNotification(conf.ServerChan.Secret, url.QueryEscape(content))
		log.Info(result)
	}
}
