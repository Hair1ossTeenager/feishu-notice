package router

import "github.com/labstack/echo/v4"

func Init(e *echo.Echo)  {
	api(e)
	static(e)
}

// api 路由
func api(e *echo.Echo)  {
	Root(e)
	Feishu(e)
}

// 静态文件路由
func static(e *echo.Echo)  {
	// web icon
	e.File("/favicon.ico", "favicon.ico")
}