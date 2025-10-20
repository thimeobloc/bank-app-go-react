package routes

import (
	"banque-app/backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, c *controllers.UserController) {
	r.POST("/auth/register", c.Register)
}
