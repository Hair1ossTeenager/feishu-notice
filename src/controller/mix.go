package controller

import (
	"feishu-notice/src/manager"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Mix(c echo.Context) error {
	sendStatus := "send failed"
	if sendNoticeByEvent(c) {
		sendStatus = "send success"
	}
	return c.String(http.StatusOK, sendStatus)
}

//  根据 event 选择对应的渠道发送通知
func sendNoticeByEvent(c echo.Context) bool {
	event := c.Request().Header.Get("X-Gitlab-Event")
	switch event {
	case "Push Hook":
		push := new(manager.Push)
		return push.PushNotice(c)
	case "Merge Request Hook":
		merge := new(manager.Merge)
		return merge.MergeNotice(c)
	case "Pipeline Hook":
		pipeline := new(manager.Pipeline)
		return pipeline.PipelineNotice(c)
	default:
		return false
	}
}