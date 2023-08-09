package routes

import (
	"github.com/golangast/goservershell/src/handler/get/home"
	"github.com/golangast/goservershell/src/handler/get/loginemail"
	"github.com/golangast/goservershell/src/handler/post/createuser"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	e.GET("/home", home.Home)
	e.GET("/loginemail/:email/:sitetoken", loginemail.LoginEmail)
	e.POST("/usercreate", createuser.Createuser)

	//#routes

}
