package routes

import (
	"server/src/controller"

	"github.com/gin-gonic/gin"
)

func ClothRoutes(routerGroup *gin.RouterGroup, clothController *controller.ClothController) {
    routerGroup.GET("/", clothController.ListCloth)
    routerGroup.GET("/:id", clothController.GetCloth)
    routerGroup.POST("/", clothController.CreateCloth)
    routerGroup.PUT("/:id", clothController.UpdateCloth)
    routerGroup.DELETE("/:id", clothController.DeleteCloth)
}
