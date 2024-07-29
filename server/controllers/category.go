package controllers

import (
	"net/http"
	modelCategory "server/models/category"
	categoryService "server/services/category"

	"github.com/gin-gonic/gin"
)

type controllerCategory struct{}

type ControllerCategory interface {
	Create(ctx *gin.Context)
}

func NewControllerCategory() *controllerCategory {
	return &controllerCategory{}
}

// getAlbums responds with the list of all albums as JSON.
func (controllerC *controllerCategory) Create(ctx *gin.Context) {
	// Create category
	var category modelCategory.CategoryModel
	if err := ctx.BindJSON(&category); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save category
	if err := categoryService.CreateCategory(&category); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Response
	ctx.IndentedJSON(http.StatusOK, category)
}
