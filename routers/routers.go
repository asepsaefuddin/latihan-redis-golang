package routers

import (
	"gin-quickstart/controllers"
	"gin-quickstart/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouters(r *gin.Engine) { //-> gin engine (roting dan middleware)
	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		// create
		auth.POST("/anime", controllers.CreateAnime)
		// get all
		auth.GET("/anime", controllers.GetAllAnime)
		// get by id
		auth.GET("/anime/:id", controllers.GetByIdAnime)
		// update
		auth.PUT("/anime/:id", controllers.UpdateAnime)
		// deleted
		auth.DELETE("/anime/:id", controllers.DeleteAnime)
	}
}
