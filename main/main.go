package main

import (
	"helloGin/config"
	"helloGin/route"
)

/**
 主运行函数
**/
func main() {
	route.GIN.Run(config.MYCONFIG.Port)
}
