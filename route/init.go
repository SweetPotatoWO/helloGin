package route

import (
	"sync"

	"github.com/gin-gonic/gin"
)

//只执行一次的函数种子
var once sync.Once

//GIN ... 框架的核心应用
var GIN *gin.Engine

//GetGinIns 单例模式 生成整个程序的核心gin
func GetGinIns() (g *gin.Engine) {
	once.Do(func() {
		GIN = gin.Default()
	}) //有且只运行一次
	return GIN
}
