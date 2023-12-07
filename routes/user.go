package routes

import (
	"github.com/gin-gonic/gin"
	"weblog/constants"
	"weblog/controllers"
	"weblog/middlewares"
)

func UserRoutes(route *gin.Engine, controller controllers.IUserController) {
	route.POST(constants.LoginPath, controller.LoginController)
	route.POST(constants.SignupPath, controller.SignupController)
	route.GET(constants.GetAllUsersPath, middlewares.Authentication, controller.GetAllUsersController)
	route.GET(constants.GetUserPath, middlewares.Authentication, controller.GetUserController)
	route.PUT(constants.UpdateUserPath, middlewares.Authentication, controller.UpdateUserController)
	route.DELETE(constants.DeleteUserPath, middlewares.Authentication, controller.DeleteUserController)
}
