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
	route.GET(constants.GetUserByEmailPath, middlewares.Authentication, controller.GetUserByEmailController)
	route.PUT(constants.UpdateUserByEmailPath, middlewares.Authentication, controller.UpdateController)
	route.DELETE(constants.DeleteUserByEmailPath, middlewares.Authentication, controller.DeleteUserByIdController)
}
