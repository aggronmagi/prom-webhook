## Alertmanager Webhook

Webhook service support send Prometheus 2.0 alert message to Dingtalk/Feishu.

## How To Use

```
cd cmd/webhook
go build
# 钉钉 
webhook -d=true -a=:5001 -w=https://oapi.dingtalk.com/robot/send?access_token=xxxx
# 飞书
webhook -a=:5001 -w=https://open.feishu.cn/open-apis/bot/v2/hook/xxxxx
```

Usage of ./webhook:
  -a, --addr string      默认监听地址 (default ":5001")
  -d, --dingtalk         true: 使用钉钉 default:飞书
  -w, --webhook string   webhook地址


使用-w 参数指定机器人web地址，或者 在告警规则内添加roboturl字段

```
groups:
- name: hostStatsAlert
  rules:
  - alert: hostCpuUsageAlert
    expr: sum(avg without (cpu)(irate(node_cpu{mode!='idle'}[5m]))) by (instance) > 0.85
    for: 1m
    labels:
      severity: page
    annotations:
      summary: "Instance {{ $labels.instance }} CPU usgae high"
      description: "{{ $labels.instance }} CPU usage above 85% (current value: {{ $value }})"
      roboturl: "https://oapi.dingtalk.com/robot/send?access_token=xxxx"
```
