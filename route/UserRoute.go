package route

import (
	"helloGin/controller"

	"github.com/gin-gonic/gin"
)

func init() {
	SetUserRoute(GetGinIns()) //设置路由
}

//SetUserRoute 在这里设置整个User模块的路由
//新建一个则填写一个
func SetUserRoute(r *gin.Engine) *gin.Engine {
	r.GET("/", controller.Run)
	return r
}
