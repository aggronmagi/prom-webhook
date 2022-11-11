package main

import (
	"net/http"

	model "github.com/aggronmagi/prom-webhook/model"
	"github.com/aggronmagi/prom-webhook/notifier"
	"github.com/gin-gonic/gin"
	"github.com/spf13/pflag"
)

var (
	defaultRobot string
	useDingTalk  bool
	addr         string = ":5001"
	script       string
)

func init() {
	pflag.StringVarP(&defaultRobot, "webhook", "w", "", "webhook地址")
	pflag.BoolVarP(&useDingTalk, "dingtalk", "d", false, "true: 使用钉钉 default:飞书")
	pflag.StringVarP(&addr, "addr", "a", addr, "默认监听地址")
	//pflag.StringVarP(&script, "script", "s", script, "lua脚本，用于过滤及组装消息")
}

func main() {
	pflag.Parse()

	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if useDingTalk {
			err = notifier.SendDingTalk(notification, defaultRobot)
		} else {
			err = notifier.SendFeiShu(notification, defaultRobot)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "send to dingtalk successful!"})

	})

	router.Run(addr)
}
