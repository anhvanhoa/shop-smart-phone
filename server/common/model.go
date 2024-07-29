package common

import "time"

type ModelCommon struct {
	Id        int       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
