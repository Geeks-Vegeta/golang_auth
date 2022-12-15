package routes

import (
	"app/controllers"

	"github.com/gin-gonic/gin"
)

func HomeRoute(c *gin.Engine) {
	public := c.Group("/")
	public.GET("/", controllers.HomeController)
}
