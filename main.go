package main

import (
	"github.com/gin-gonic/gin"
	"weblog/configs"
	"weblog/constants"
	"weblog/controllers"
	"weblog/database"
	"weblog/repositories"
	"weblog/routes"
	"weblog/services"
)

func main() {
	app := gin.Default()
	config, _ := configs.LoadDatabaseConfig()
	db := database.InitDatabase(config)

	userRepository := repositories.CreateUserRepository(db)
	userService := services.CreateUserService(userRepository)
	userController := controllers.CreateUserController(userService)

	routes.UserRoutes(app, userController)

	_ = app.Run(constants.Server)
}
