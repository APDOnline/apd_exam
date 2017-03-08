package route

import (
	"github.com/apd_exam/controller"
	"github.com/labstack/echo"
)

func InitRoute(e *echo.Echo) {

	userAPI := e.Group("api/user")
	userAPI.GET("/getList", controller.GetUsers)
	userAPI.POST("/register", controller.RegisterAccount)
	userAPI.POST("/login", controller.Login)

	questionAPI := e.Group("api/book")
	questionAPI.GET("/", controller.GetAllBooks)
	questionAPI.GET("/:bookID/questionList", controller.GetQuestionsByBook)

	examAPI := e.Group("api/exam")
	examAPI.GET("/get/:id", controller.GetExam)
	examAPI.POST("/create", controller.CreateExam)
	examAPI.POST("/updateQuestion", controller.UpdateQuestionToExam)
	examAPI.GET("/list", controller.GetExamList)
}
