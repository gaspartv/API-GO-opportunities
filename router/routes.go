package router

import (
	"github.com/gaspartv/API-GO-opportunities/docs"
	"github.com/gaspartv/API-GO-opportunities/handler"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
    // Initialize handler
    handler.InitializeHandler()
    basePath := "/api/v1"
    docs.SwaggerInfo.BasePath = basePath
    v1 := router.Group("/api/v1")
    {
        v1.GET("/opening", handler.ShowOpeningHandler)
        v1.POST("/opening", handler.CreateOpeningHandler)
        v1.DELETE("/opening", handler.DeleteOpeningHandler)
        v1.PUT("/opening", handler.UpdateOpeningHandler)
        v1.GET("/openings", handler.ListOpeningHandlers)
    }
    // Initialize swagger
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
