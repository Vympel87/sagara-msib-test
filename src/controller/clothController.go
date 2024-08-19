package controller

import (
	"net/http"
	"strconv"

	"server/src/model"
	schemas "server/src/schema"
	"server/src/service"
	"server/src/util"

	"github.com/gin-gonic/gin"
)

type ClothController struct {
    Service service.ClothService
}

func (c *ClothController) ListCloth(ctx *gin.Context) {
    cloths, err := c.Service.ListCloth(ctx.Request.Context())
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    formattedCloths := make([]map[string]interface{}, len(cloths))
    for i, cloth := range cloths {
        formattedCloths[i] = map[string]interface{}{
            "id":    cloth.ID,
            "name":  cloth.Name,
            "color": cloth.Color,
            "size":  cloth.Size,
            "price": util.FormatCurrency(cloth.Price),
            "stock": util.FormatStock(cloth.Stock),
        }
    }

    ctx.JSON(http.StatusOK, formattedCloths)
}

func (c *ClothController) GetCloth(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    cloth, err := c.Service.GetCloth(ctx.Request.Context(), id)
    if err != nil {
        ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
        return
    }

    formattedCloth := map[string]interface{}{
        "id":    cloth.ID,
        "name":  cloth.Name,
        "color": cloth.Color,
        "size":  cloth.Size,
        "price": util.FormatCurrency(cloth.Price),
        "stock": util.FormatStock(cloth.Stock),
    }

    ctx.JSON(http.StatusOK, formattedCloth)
}

func (c *ClothController) CreateCloth(ctx *gin.Context) {
    var clothSchema schemas.ClothCreateUpdateSchema
    if err := ctx.ShouldBindJSON(&clothSchema); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newCloth := model.Cloth{
        Name:  clothSchema.Name,
        Color: clothSchema.Color,
        Size:  clothSchema.Size,
        Price: clothSchema.Price,
        Stock: clothSchema.Stock,
    }

    createdCloth, err := c.Service.CreateCloth(ctx.Request.Context(), &newCloth)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    formattedCloth := map[string]interface{}{
        "id":    createdCloth.ID,
        "name":  createdCloth.Name,
        "color": createdCloth.Color,
        "size":  createdCloth.Size,
        "price": util.FormatCurrency(createdCloth.Price),
        "stock": util.FormatStock(createdCloth.Stock),
    }

    ctx.JSON(http.StatusCreated, formattedCloth)
}

func (c *ClothController) UpdateCloth(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var clothSchema schemas.ClothCreateUpdateSchema
    if err := ctx.ShouldBindJSON(&clothSchema); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    updatedCloth := model.Cloth{
        Name:  clothSchema.Name,
        Color: clothSchema.Color,
        Size:  clothSchema.Size,
        Price: clothSchema.Price,
        Stock: clothSchema.Stock,
    }

    cloth, err := c.Service.UpdateCloth(ctx.Request.Context(), id, &updatedCloth)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    formattedCloth := map[string]interface{}{
        "id":    cloth.ID,
        "name":  cloth.Name,
        "color": cloth.Color,
        "size":  cloth.Size,
        "price": util.FormatCurrency(cloth.Price),
        "stock": util.FormatStock(cloth.Stock), 
    }

    ctx.JSON(http.StatusOK, formattedCloth)
}


func (c *ClothController) DeleteCloth(ctx *gin.Context) {
    id, err := strconv.Atoi(ctx.Param("id"))
    if err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    err = c.Service.DeleteCloth(ctx.Request.Context(), id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    ctx.Status(http.StatusNoContent)
}