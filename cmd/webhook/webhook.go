package main

import (
	"flag"
	"net/http"

	model "github.com/aggronmagi/prom-webhook/model"
	"github.com/aggronmagi/prom-webhook/notifier"
	"github.com/gin-gonic/gin"
)

var (
	h            bool
	defaultRobot string
	useDingTalk  bool
	addr         string = "127.0.0.1:5001"
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&defaultRobot, "defaultRobot", "", "global dingtalk robot webhook, you can overwrite by alert rule with annotations dingtalkRobot")
	flag.BoolVar(&useDingTalk, "dingtalk", false, "true: use dingtalk default:feishu")
	flag.StringVar(&addr, "addr", addr, "默认监听地址")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

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
