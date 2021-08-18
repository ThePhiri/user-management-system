package routes

import (
	controller "github.com/thephiri/user-management-system/go-backend/controllers"
	"github.com/thephiri/user-management-system/go-backend/middleware"

	"github.com/gin-gonic/gin"
)

//UserRoutes function
func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authentication())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}
