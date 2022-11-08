package transformer

import (
	"fmt"

	"github.com/aggronmagi/prom-webhook/model"
)

// TransformToFeiShuPost transform alertmanager notification to feishu post message
func TransformToFeiShuPost(notification model.Notification) (ret interface{}, robotURL string, err error) {

	groupKey := notification.CommonLabels["alertname"]
	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["robot"]

	//fmt.Printf("%#v\n", notification)

	title := fmt.Sprintf("%s-%s", groupKey, status)

	contents := make([]*model.PostContent, 0, len(notification.Alerts))

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		contents = append(contents, &model.PostContent{
			Tag: "text",
			Text: fmt.Sprintf("  %s 开始时间:%s\n",
				//annotations["summary"],
				annotations["description"],
				alert.StartsAt.Local().Format("15:04:05"),
			),
		})
	}

	ret = model.NewFeiShuPost(title, contents)

	return
}
