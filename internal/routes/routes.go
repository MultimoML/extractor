package routes

import (
	"extractor-timer/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/self", controllers.Self())
	router.GET("/info", controllers.Info())
	router.POST("/extract", controllers.Extract())

	/*
		router.POST("/user", controllers.CreateUser())
		router.GET("/user/:userId", controllers.GetAUser())

		router.PUT("/user/:userId", controllers.EditAUser())
		router.DELETE("/user/:userId", controllers.DeleteAUser())
		router.GET("/users", controllers.GetAllUsers())
	*/
}
