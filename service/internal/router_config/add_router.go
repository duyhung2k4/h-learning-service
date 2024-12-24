package routerconfig

import (
	constant "app/internal/constants"

	"github.com/gin-gonic/gin"
)

func AddRouter(r *gin.Engine, config RouterConfig) {
	routerGroup := r.Group("api/v1")
	handle := []gin.HandlerFunc{}
	handle = append(handle, config.Middleware...)
	handle = append(handle, config.Handle)

	switch config.Method {
	case constant.GET_HTTP:
		routerGroup.GET(config.Endpoint, handle...)
	case constant.POST_HTTP:
		routerGroup.POST(config.Endpoint, handle...)
	case constant.PUT_HTTP:
		routerGroup.PUT(config.Endpoint, handle...)
	case constant.PATCH_HTTP:
		routerGroup.PATCH(config.Endpoint, handle...)
	case constant.DELETE_HTTP:
		routerGroup.DELETE(config.Endpoint, handle...)
	}
}