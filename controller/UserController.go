package controller

import (
	"fmt"
)

//User ...
//用户的控制器
type User struct {
}

func (u *User) hello() {
	fmt.Println("你好啊啊")
}
