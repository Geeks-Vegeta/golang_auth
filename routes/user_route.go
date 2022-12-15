package routes

import (
	"app/controllers"
	"app/middlewares"

	"github.com/gin-gonic/gin"
)

func User(router *gin.Engine) {
	public := router.Group("/auth")
	public.POST("/register", controllers.CreateUser())
	public.POST("/login", controllers.LoginUser())

	protected := router.Group("/user")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/all", controllers.GetAllUser())
	protected.GET("/current", controllers.CurrentUser)

}
