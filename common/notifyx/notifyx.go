package notifyx

import (
	"bytes"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strings"
	"text/template"
	"time"
)

var MailService *Mail

type Message map[string]interface{}

type Notifiable interface {
	Send(msg Message)
}

var queue = make(chan Message, 100)

func init() {
	go run()
}

// Push 把消息推入队列
func Push(msg Message) {
	queue <- msg
}

func run() {
	for msg := range queue {
		// 根据任务配置发送通知
		taskType, taskTypeOk := msg["task_type"]
		_, taskReceiverIdOk := msg["task_receiver_id"]
		_, nameOk := msg["name"]
		_, outputOk := msg["output"]
		_, statusOk := msg["status"]
		if !taskTypeOk || !taskReceiverIdOk || !nameOk || !outputOk || !statusOk {
			logx.Errorf("#notify#参数不完整#%+v", msg)
			continue
		}
		msg["content"] = fmt.Sprintf("============\n============\n============\n"+
			"任务名称: %s\n状态: %s\n输出:\n %s\n", msg["name"], msg["status"], msg["output"])
		switch taskType.(string) {
		case "mail":
			// 邮件
			//mail := Mail{}
			go MailService.Send(msg)
		case "2":
			// Slack
			//slack := Slack{}
			//go slack.Send(msg)
		case "3":
			// WebHook
			//webHook := WebHook{}
			//go webHook.Send(msg)
		}
		time.Sleep(1 * time.Second)
	}
}

func parseNotifyTemplate(notifyTemplate string, msg Message) string {
	tmpl, err := template.New("notify").Parse(notifyTemplate)
	if err != nil {
		return fmt.Sprintf("解析通知模板失败: %s", err)
	}
	var buf bytes.Buffer
	tmpl.Execute(&buf, map[string]interface{}{
		"TaskId":   msg["task_id"],
		"TaskName": msg["name"],
		"Status":   msg["status"],
		"Result":   msg["output"],
		"Remark":   msg["remark"],
	})

	return buf.String()
}

func InStringSlice(slice []string, element string) bool {
	element = strings.TrimSpace(element)
	for _, v := range slice {
		if strings.TrimSpace(v) == element {
			return true
		}
	}

	return false
}
