package main

import (
	"go-react-chat-app/initializers"
	"go-react-chat-app/routes"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/glebarez/go-sqlite"
)

func init() {
	initializers.ConnectToDB()
}

func main() {
	route := gin.Default()

	routes.UserRoutes(route)
	routes.ChannelRoutes(route)
	routes.MessageRoutes(route)

	err := route.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
