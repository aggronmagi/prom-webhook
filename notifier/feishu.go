package notifier

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/aggronmagi/prom-webhook/model"
	"github.com/aggronmagi/prom-webhook/transformer"
)

func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret
	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}

// SendFeiShu send markdown message to dingtalk
func SendFeiShu(notification model.Notification, defaultRobot string) (err error) {

	msg, robotURL, err := transformer.TransformToFeiShuPost(notification)

	if err != nil {
		return
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return
	}

	var dingTalkRobotURL string

	if robotURL != "" {
		dingTalkRobotURL = robotURL
	} else {
		dingTalkRobotURL = defaultRobot
	}

	if len(dingTalkRobotURL) == 0 {
		return nil
	}

	req, err := http.NewRequest(
		"POST",
		dingTalkRobotURL,
		bytes.NewBuffer(data))

	if err != nil {
		fmt.Println("dingtalk robot url not found ignore:")
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)

	return
}
