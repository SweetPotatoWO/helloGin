package main

import (
	"fmt"
	"reflect"
)

func main() {

	user := User{
		Name: "zhangsan",
	}
	funcName := "Signup"
	v := make([]reflect.Value, 3)
	// 数组第一个值必须为 结构体实例的reflect.Value, 如果函数没有参数，这个值也必须有，v[1],v[2]不需要
	v[0] = reflect.ValueOf(user)
	v[1] = reflect.ValueOf("likun")
	v[2] = reflect.ValueOf(27)
	// m 类型为  reflect.Value
	m, _ := reflect.TypeOf(user).MethodByName(funcName)
	m.Func.Call(v)

}

type User struct {
	Name string
	age  string
}

func (u User) Signup(name string, age int) {
	fmt.Println("name123123:", name, "age:", age)
}

func (u User) Signin() {
	fmt.Println("name:", u.Name)
}
