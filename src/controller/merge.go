package controller

import (
	"feishu-notice/src/manager"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Merge(c echo.Context) error {
	sendStatus := "send failed"
	merge := new(manager.Merge)
	if merge.MergeNotice(c) {
		sendStatus = "send success"
	}
	return c.String(http.StatusOK, sendStatus)
}
