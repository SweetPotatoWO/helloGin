package main

import (
	"fmt"
)

func main() {

	fmt.Println((*User)(nil)) //新建一个User对象 并初始化为nil
	var a User
	fmt.Println(a)
}

type User struct {
}
