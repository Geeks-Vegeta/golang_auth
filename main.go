package main

import (
	"app/models"
	"app/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	models.Connection()

	routes.HomeRoute(route)
	routes.User(route)
	route.Run(":5000")

}
