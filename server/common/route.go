package common

import "github.com/gin-gonic/gin"

type RouteCommon struct {
	Path      string
	GinEngine *gin.Engine
}

type InterRouteCommon interface {
	routes()
}
