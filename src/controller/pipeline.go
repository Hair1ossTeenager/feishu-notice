package controller

import (
	"feishu-notice/src/manager"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Pipeline(c echo.Context) error {
	sendStatus := "send failed"
	pipeline := new(manager.Pipeline)
	if pipeline.PipelineNotice(c) {
		sendStatus = "send success"
	}
	return c.String(http.StatusOK, sendStatus)
}
