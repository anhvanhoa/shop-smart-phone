package modelCategory

import "server/common"

type CategoryModel struct {
	common.ModelCommon
	Name   string `json:"name" gorm:"type:varchar(150)" validate:"required"`
	Slug   string `json:"slug" gorm:"type:varchar(255)"`
	Status bool   `json:"status" gorm:"type:tinyint(1) default 1"`
}

func (category *CategoryModel) TableName() string {
	return "categories"
}
