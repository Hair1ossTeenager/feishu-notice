package router

import (
	"feishu-notice/src/controller"
	"github.com/labstack/echo/v4"
)

/*
/feishu/xxx api 组
 */
func Feishu(e *echo.Echo)  {
	api := e.Group("feishu")
	pipeLine(api)
	push(api)
	merge(api)
	mix(api)
}

/*
/feishu/pipeline/:web_hook
gitlab pipeline 飞书通知 api
 */
func pipeLine(api *echo.Group) {
	api.Any("/pipeline/:web_hook", controller.Pipeline)
}

/*
/feishu/push/:web_hook
gitlab push 飞书通知 api
*/
func push(api *echo.Group) {
	api.Any("/push/:web_hook", controller.Push)
}

/*
/feishu/merge/:web_hook
gitlab merge 飞书通知 api
*/
func merge(api *echo.Group) {
	api.Any("/merge/:web_hook", controller.Merge)
}

/*
/feishu/:web_hook
gitlab push/merge/pipeline 飞书通知 api
*/
func mix(api *echo.Group) {
	api.Any("/mix/:web_hook", controller.Mix)
}
