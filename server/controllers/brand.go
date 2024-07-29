package controllers

import (
	modelBrand "server/models/brand"
	brandService "server/services/brand"

	"github.com/gin-gonic/gin"
)

type controllerBrand struct{}

type ControllerBrand interface {
	GetList(ctx *gin.Context)
}

// NewControllerBrand is a constructor
func NewControllerBrand() *controllerBrand {
	return &controllerBrand{}
}

func (controllerB *controllerBrand) GetList(ctx *gin.Context) {
	ctx.IndentedJSON(200, [3]string{"Nguyen", "Van", "Anh"})
}

func (controllerB *controllerBrand) Create(ctx *gin.Context) {
	// Create brand
	var brand *modelBrand.BrandModel
	if err := ctx.BindJSON(&brand); err != nil {
		ctx.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	// Save brand
	if err := brandService.CreateBrand(brand); err != nil {
		ctx.IndentedJSON(500, gin.H{"error": err.Error()})
		return
	}
	// Response
	ctx.IndentedJSON(200, gin.H{"data": brand})
}
