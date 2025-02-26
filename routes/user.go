package routes

import (
	"go-react-chat-app/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	r.POST("/user", controllers.CreateUser)
	r.POST("/login", controllers.Login)
}
