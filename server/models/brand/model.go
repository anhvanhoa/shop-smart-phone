package modelBrand

import "server/common"

type BrandModel struct {
	common.ModelCommon
	Name string `json:"name" gorm:"type:varchar(150)" validate:"required"`
	Logo string `json:"logo" gorm:"type:varchar(255)" validate:"required"`
}

func (brand *BrandModel) TableName() string {
	return "brands"
}
