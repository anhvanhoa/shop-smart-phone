package category

import (
	"server/common"
	"server/controllers"

	"github.com/gin-gonic/gin"
)

type routeCategory struct {
	common.RouteCommon
}

func (rC *routeCategory) routes() {
	controllerCategory := controllers.NewControllerCategory()
	routes := rC.GinEngine.Group(rC.Path)
	routes.POST("", controllerCategory.Create)
}

func InitRoute(ginEngine *gin.Engine) {
	r := &routeCategory{
		RouteCommon: common.RouteCommon{
			Path:      "category",
			GinEngine: ginEngine,
		},
	}
	r.routes()
}
