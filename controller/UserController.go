package controller

import (
	"fmt"
)

func init() {
	registerType((*User)(nil)) //新建一个对象并初始化为nil
}

//User ...
//用户的控制器
type User struct{}

//Hello ...
func (u User) Hello(param interface{}) string {
	fmt.Println(param)
	fmt.Println("你好啊啊")
	return "dafafadfadf"
}
