package main

import (
	"banque-app/backend/controllers"
	"banque-app/backend/db"
	"banque-app/backend/repositories"
	"banque-app/backend/routes"
	"banque-app/backend/services"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	userRepo := &repositories.UserRepository{DB: db.DB}
	userService := &services.UserService{Repo: userRepo}
	userController := &controllers.UserController{Service: userService}

	r := gin.Default()
	routes.RegisterUserRoutes(r, userController)

	r.Run(":8080")
}
