package brandService

import (
	"server/config"
	modelBrand "server/models/brand"
)

func CreateBrand(brand *modelBrand.BrandModel) error {
	if err := config.DB.Create(brand).Error; err != nil {
		return err
	}
	return nil
}
