package route

import (
	"github.com/gin-gonic/gin"
)

func init() {
	SetUserRoute(GetGinIns()) //设置路由
}

//SetUserRoute 在这里设置整个User模块的路由
//新建一个则填写一个
func SetUserRoute(r *gin.Engine) *gin.Engine {
	r.GET("/", hello)
	return r
}

func hello(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": "ddd",
		"nick":    "cc",
	})
}
