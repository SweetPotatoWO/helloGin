package route

func init() {
	SetUserRoute(GetGinIns(), GetURLRoute()) //设置路由
}

//GetURLRoute 这里填写Url全部注册路径
func GetURLRoute() map[string]URLRouteOne {
	var ret map[string]URLRouteOne
	ret = make(map[string]URLRouteOne)

	ret["/"] = URLRouteOne{
		URL:  "/",
		Type: "get",
		Path: "User/Hello",
	}

	ret["/hello"] = URLRouteOne{
		URL:  "/hello",
		Type: "post",
		Path: "User/HelloPost",
	}

	return ret
}
