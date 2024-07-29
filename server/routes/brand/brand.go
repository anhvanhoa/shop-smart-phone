package brand

import (
	"server/common"
	"server/controllers"

	"github.com/gin-gonic/gin"
)

type routeBarnd struct {
	common.RouteCommon
}

func (rB *routeBarnd) routes() {
	controllerBrand := controllers.NewControllerBrand()
	routes := rB.GinEngine.Group(rB.Path)
	routes.GET("", controllerBrand.GetList)
	routes.POST("", controllerBrand.Create)
}

func InitRoute(ginEngine *gin.Engine) {
	r := &routeBarnd{
		RouteCommon: common.RouteCommon{
			Path:      "brand",
			GinEngine: ginEngine,
		},
	}
	r.routes()
}
