package controller

func init() {
	var a *User
	RegisterType(a) //新建一个对象并初始化为nil
}

//User ...
//用户的控制器
type User struct{}

//Hello ...
func (u User) Hello(param interface{}) string {
	return ReturnJsonSuccess("返回成功", "dafda")
}

//helloUser
func (u User) HelloPost(param interface{}) string {
	return ReturnJsonSuccess("返回成功", "nmzdsl")
}
