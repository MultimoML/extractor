package routes

import (
	"extractor/internal/controllers"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	extractor := router.Group("/extractor")
	extractor.GET("/self", controllers.Self())

	v1 := extractor.Group("/v1")
	v1.GET("/info", controllers.Info())
	v1.POST("/extract", controllers.Extract())

	/*
		router.POST("/user", controllers.CreateUser())
		router.GET("/user/:userId", controllers.GetAUser())

		router.PUT("/user/:userId", controllers.EditAUser())
		router.DELETE("/user/:userId", controllers.DeleteAUser())
		router.GET("/users", controllers.GetAllUsers())
	*/
}
