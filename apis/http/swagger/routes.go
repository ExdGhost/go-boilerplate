package swagger

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "go-boilerplate/docs"
)

// NewSwaggerRoutes initializes all the Swagger UI routes
func NewSwaggerRoutes(router *gin.Engine) {
	// Inits Swagger UI endpoint
	url := ginSwagger.URL("http://localhost:80/swagger/doc.json") // The url pointing to API definition

	// Defines supported methods to parse into definitions
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.POST("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.PUT("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	router.DELETE("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
