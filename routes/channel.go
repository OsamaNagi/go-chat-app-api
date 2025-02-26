package routes

import (
	"go-react-chat-app/controllers"

	"github.com/gin-gonic/gin"
)

func ChannelRoutes(r *gin.Engine) {
	r.POST("/channels", controllers.CreateChannel)
	r.GET("/channels", controllers.ListChannels)
}
