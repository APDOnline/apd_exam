package route

import (
	"github.com/apd_exam/controller"
	"github.com/labstack/echo"
)

func InitRoute(e *echo.Echo) {
	userAPI := e.Group("/user")
	userAPI.GET("/getList", controller.GetUsers)
	userAPI.POST("/register", controller.RegisterAccount)
	userAPI.POST("/login", controller.Login)
}
