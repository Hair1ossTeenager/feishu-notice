package router

import (
	"feishu-notice/src/controller"
	"github.com/labstack/echo/v4"
	"net/http"
)

/*
/xxx api 组
 */
func Root(e *echo.Echo)  {
	api := e.Group("")
	root(api)
	heart(api)
}

/*
/
跟路径
 */
func root(api *echo.Group)  {
	api.Any("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello world!")
	})
}

/*
/heart
健康检查
 */
func heart(api *echo.Group)  {
	api.Any("/heart", controller.Heart)
}
