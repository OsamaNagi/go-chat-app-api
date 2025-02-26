package routes

import (
	"go-react-chat-app/controllers"

	"github.com/gin-gonic/gin"
)

func MessageRoutes(r *gin.Engine) {
	r.POST("/messages", controllers.CreateMessage)
	r.GET("/messages", controllers.ListMessages)
}
