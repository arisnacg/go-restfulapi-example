package main

import (
	"arisnacg/go-restfulapi-example/controllers"
	"arisnacg/go-restfulapi-example/middlewares"
	"arisnacg/go-restfulapi-example/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/me", controllers.CurrentUser)

	r.Run(":3000")
}
