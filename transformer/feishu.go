package transformer

import (
	"fmt"

	"github.com/aggronmagi/prom-webhook/model"
)

// TransformToFeiShuPost transform alertmanager notification to feishu post message
func TransformToFeiShuPost(notification model.Notification) (ret interface{}, robotURL string, err error) {

	groupKey := notification.GroupKey
	status := notification.Status

	annotations := notification.CommonAnnotations
	robotURL = annotations["robot"]

	fmt.Printf("%#v\n", notification)

	title := fmt.Sprintf("通知组%s(当前状态:%s)", groupKey, status)

	contents := make([]*model.PostContent, 0, len(notification.Alerts))

	contents = append(contents, &model.PostContent{
		Tag:  "link",
		Text: "链接\n",
		HRef: notification.ExternalURL,
	})

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		contents = append(contents, &model.PostContent{
			Tag:  "item",
			Text: annotations["summary"],
		})
		contents = append(contents, &model.PostContent{
			Tag:  "desc",
			Text: annotations["description"],
		})
		contents = append(contents, &model.PostContent{
			Tag:  "start",
			Text: alert.StartsAt.Format("15:04:05"),
		})
		contents = append(contents, &model.PostContent{
			Tag:  "e",
			Text: "\n",
		})
	}

	ret = model.NewFeiShuPost(title, contents)

	return
}
