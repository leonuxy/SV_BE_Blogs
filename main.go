package main

import (
	"sv_be_posts/config"
	"sv_be_posts/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.ConnectDB()

	router := gin.Default()
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
