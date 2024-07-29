package config

import (
	modelBrand "server/models/brand"
	modelCategory "server/models/category"

	"github.com/gofor-little/env"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Database connection
	var dsn string = env.Get("DATABASE_URL", "")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	// Migrate the schema
	db.AutoMigrate(
		&modelBrand.BrandModel{},
		&modelCategory.CategoryModel{},
	)
	DB = db
}
