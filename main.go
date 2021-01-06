package main

import (
	"feishu-notice/src/router"
	"github.com/labstack/echo/v4"
	//"github.com/labstack/echo/v4/middleware"
)

func main()  {
	// Echo instance
	e := echo.New()

	// Routes
	router.Init(e)

	// Logger
	//e.Use(middleware.Logger())

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}