package routes

import (
	"server/src/controller"

	"github.com/gin-gonic/gin"
)

func RootRoutes(clothController *controller.ClothController) *gin.Engine {
    router := gin.Default()

    apiGroup := router.Group("/api")
    {
        ClothRoutes(apiGroup.Group("/cloths"), clothController)
    }

    return router
}
