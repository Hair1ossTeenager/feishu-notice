package controller

import (
	"feishu-notice/src/manager"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Push(c echo.Context) error {
	sendStatus := "send failed"
	push := new(manager.Push)
	if push.PushNotice(c) {
		sendStatus = "send success"
	}
	return c.String(http.StatusOK, sendStatus)
}
