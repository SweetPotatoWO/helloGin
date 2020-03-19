package route

import (
	"encoding/json"
	"helloGin/config"
	"helloGin/controller"
	"log"
	"reflect"
	"regexp"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

//只执行一次的函数种子
var once sync.Once

//RouteRegisterTable 路由注册的内存表
var RouteRegisterTable = make(map[string]string)

//GIN ... 框架的核心应用
var GIN *gin.Engine

//GetGinIns 单例模式 生成整个程序的核心gin
func GetGinIns() (g *gin.Engine) {
	once.Do(func() {
		GIN = gin.Default()
	}) //有且只运行一次
	return GIN
}

//URLRouteOne 每个Url注册表达的单元节点
type URLRouteOne struct {
	URL  string
	Type string
	Path string
}

//SetUserRoute 在这里设置整个User模块的路由
//新建一个则填写一个
func SetUserRoute(r *gin.Engine, URLTable map[string]URLRouteOne) *gin.Engine {
	if len(URLTable) <= 0 {
		log.Fatalln("URL映射表为空")
	}
	for _, value := range URLTable {
		if strings.ToLower(value.Type) == "get" {
			r.GET(value.URL, Run)
			RouteRegisterTable[value.URL] = value.Path //将注册过的保存起来
		} else if strings.ToLower(value.Type) == "post" {
			r.POST(value.URL, Run)
			RouteRegisterTable[value.URL] = value.Path //将注册过的保存起来
		} else if strings.ToLower(value.Type) == "put" {
			r.PUT(value.URL, Run)
			RouteRegisterTable[value.URL] = value.Path //将注册过的保存起来
		} else if strings.ToLower(value.Type) == "delete" {
			r.DELETE(value.URL, Run)
			RouteRegisterTable[value.URL] = value.Path //将注册过的保存起来
		} else if strings.ToLower(value.Type) == "patch" {
			r.PATCH(value.URL, Run)
			RouteRegisterTable[value.URL] = value.Path //将注册过的保存起来
		} else if strings.ToLower(value.Type) == "head" {
			r.HEAD(value.URL, Run)
			RouteRegisterTable[value.URL] = value.Path //将注册过的保存起来
		} else if strings.ToLower(value.Type) == "options" {
			r.OPTIONS(value.URL, Run)
			RouteRegisterTable[value.URL] = value.Path //将注册过的保存起来
		}
	}
	return r
}

//Run ... 主运行函数
func Run(c *gin.Context) {
	var cInitObj controller.InitController
	//1 先获取到请求所传递过来的参数
	param, err := cInitObj.GetParams(c)
	if err != nil {
		log.Fatalf("获取到参数错误%v", err)
	}

	//2 执行整个控制器运行周期的before
	isContinue := cInitObj.InitBeforeFn(param)
	if !isContinue {
		log.Fatalf("initBefore执行失败%v", err)
	}
	r := c.Request
	ReqPath := r.URL.Path //当前请求的路径
	var controllerFnCode string = ""
	for rkey, rvalue := range RouteRegisterTable {
		if matchReqUrl(ReqPath, rkey) { //如果请求的路径相同
			controllerFnCode = rvalue
			break
		}
	}
	if controllerFnCode == "" {
		log.Fatalln("请求的URL不存在")
	}

	//3 对控制器的名字和方法名进行解析
	controllerAndFnName, err := cInitObj.SplitControllerAndFnName(controllerFnCode)
	if err != nil {
		log.Fatalf("找不到控制器%v", err)
	}
	//获取到对应的控制器名字和方法名字
	ControllerName := controllerAndFnName["controller"]
	RuqFnName := controllerAndFnName["fnname"]

	//4 获取到对应配置文件的组件注册列表
	ModuleArr := config.MYCONFIG.Module
	if len(ModuleArr) <= 0 {
		log.Fatalln("不存在注册的控制器")
	}

	//5 检测是否存在控制器
	var isFlag = false
	for _, controllerOne := range ModuleArr {
		if controllerOne == ControllerName {
			isFlag = true
			break
		}
	}
	if !isFlag {
		log.Fatalln("当前传递的控制器,不在注册列表中")
	}
	//7 根据获取到控制器名字获取到对应类的实例
	initStruct, ok := controller.NewStruct(ControllerName) //返回核心的控制器
	if !ok {
		log.Fatalln("实例化核心的控制器失败")
	}

	inputs := make([]reflect.Value, 1) //参数
	inputs[0] = reflect.ValueOf(param) //输入对应的参数
	//8 根据类和方法 调用并执行
	ret := reflect.ValueOf(initStruct).MethodByName(RuqFnName).Call(inputs) //调用对应的方法

	//反射Value 转string string 转[]bety []bety转结构体 结构体转json
	resultStr := []byte(ret[0].String())
	var returnRet interface{}
	json.Unmarshal(resultStr, &returnRet)
	//9 根据类型返回对应的结果
	c.JSON(200, gin.H{
		"status":   "200",
		"response": returnRet,
	})
	//反射映射实例话对应的类
	cInitObj.InitAfterFn(param)
}

//链接的匹配度
//匹配的规则,将链接用/分割成各个的小组
//当存在: 通配符的链接则不进行匹配
//全部匹配通过返回true 匹配不过返回flase

func matchReqUrl(reqUrl string, registerUrl string) (flag bool) {
	if reqUrl == "" || registerUrl == "" {
		return false
	}
	reqUrlArr := strings.Split(reqUrl, "/")
	registerUrlArr := strings.Split(registerUrl, "/")
	if strings.Join(reqUrlArr, "") == strings.Join(registerUrlArr, "") {
		return true
	}
	flag = true //默认
	for key, regisv := range registerUrlArr {
		if !matchChar(regisv) { //是通配符吗? 是
			if regisv != reqUrlArr[key] {
				flag = false
			}
		}
	}
	return flag
}

//正则匹配出:uid之类的通配符
func matchChar(char string) (flag bool) {
	flag = false
	match, _ := regexp.MatchString(":(.*)", char)
	if match {
		flag = true
	}
	return flag
}
