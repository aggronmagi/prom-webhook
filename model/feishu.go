package model

// curl -X POST -H "Content-Type: application/json" \
//	-d '{"msg_type":"text","content":{"text":"request example"}}' \

// 文档
// https://open.larksuite.com/document/uAjLw4CM/ukTMukTMukTM/bot-v3/use-custom-bots-in-a-group?lang=zh-CN

////////////////////////////////////////////////////////////
// 文本
// {
//     "msg_type": "text",
//     "content": {
//         "text": "新更新提醒"
//     }
// }

func NewFeiShuText(txt string) map[string]interface{} {
	return map[string]interface{}{
		"msg_type": "text",
		"content":  txt,
	}
}

////////////////////////////////////////////////////////////
// 富文本
// {
//     "msg_type": "post",
//     "content": {
//         "post": {
//             "zh_cn": {
//                 "title": "项目更新通知",
//                 "content": [
//                     [
//                         {
//                             "tag": "text",
//                             "text": "项目有更新: "
//                         },
//                         {
//                             "tag": "a",
//                             "text": "请查看",
//                             "href": "http://www.example.com/"
//                         }
//                     ]
//                 ]
//             }
//         }
//     }
// }

type PostContent struct {
	Tag  string `json:"tag"`
	Text string `json:"text"`
	HRef string `json:"hfre,omitempty"`
}

func NewFeiShuPost(title string, contents []*PostContent) map[string]interface{} {
	return map[string]interface{}{
		"msg_type": "post",
		"content": map[string]interface{}{
			"post": map[string]interface{}{
				"zh_cn": map[string]interface{}{
					"title":   title,
					"content": []interface{}{contents},
				},
			},
		},
	}
}

////////////////////////////////////////////////////////////
// 群名片
// https://open.larksuite.com/document/ukTMukTMukTM/ucjMxEjL3ITMx4yNyETM
// {
//     "msg_type": "share_chat",
//     "content":{
//         "share_chat_id": "oc_f5b1a7eb27ae2c7b6adc2a74faf339ff"
//     }
// }

////////////////////////////////////////////////////////////
// 图片
// {
//     "msg_type":"image",
//     "content":{
//         "image_key": "img_ecffc3b9-8f14-400f-a014-05eca1a4310g"
//     }
// }
