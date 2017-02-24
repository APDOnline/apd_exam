package route

import (
	"github.com/apd_exam/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func InitRoute(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	userAPI := e.Group("api/user")
	userAPI.GET("/getList", controller.GetUsers)
	userAPI.POST("/register", controller.RegisterAccount)
	userAPI.POST("/login", controller.Login)

	questionAPI := e.Group("api/book")
	questionAPI.GET("/", controller.GetAllBooks)
	questionAPI.GET("/:bookID/questionList", controller.GetQuestionsByBook)
}
