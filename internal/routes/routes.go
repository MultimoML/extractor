package routes

import (
	"net/http"

	"github.com/multimoml/extractor/internal/controllers"

	_ "github.com/multimoml/extractor/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	extractor := router.Group("/extractor")
	extractor.GET("/self", controllers.Self())

	extractor.GET("/openapi", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/extractor/openapi/index.html")
	})
	extractor.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/extractor/openapi/index.html")
	})
	extractor.GET("/openapi/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := extractor.Group("/v1")
	v1.GET("/info", controllers.Info())
	v1.POST("/extract", controllers.Extract())
}
