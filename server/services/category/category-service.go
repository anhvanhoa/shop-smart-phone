package categoryService

import (
	"server/config"
	modelCategory "server/models/category"

	"github.com/go-playground/validator/v10"
	"github.com/gosimple/slug"
)

func CreateCategory(category *modelCategory.CategoryModel) error {
	// Validate
	validate := validator.New()
	err := validate.Struct(category)
	if err != nil {
		return err
	}
	// convert name to slug
	category.Slug = slug.Make(category.Name)
	if err := config.DB.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func GetCategories() ([]modelCategory.CategoryModel, error) {
	var categories []modelCategory.CategoryModel
	config.DB.Where("status = ?", 1).Find(&categories)
	if err := config.DB.Find(modelCategory.CategoryModel{}).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
