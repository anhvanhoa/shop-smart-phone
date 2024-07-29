package main

import (
	"server/config"
	"server/routes/brand"
	"server/routes/category"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.LoadEnv()
	config.ConnectDB()
	// init routes
	{
		category.InitRoute(router)
		brand.InitRoute(router)
	}
	router.Run("localhost:8080")
}
