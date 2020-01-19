package controller

import (
	"errors"
	"helloGin/config"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var typeRegistry = make(map[string]reflect.Type) //组件注册列表

//Run ... 主运行函数
func Run(c *gin.Context) {
	var cInitObj *InitController
	cInitObj = &InitController{} //
	//1 先获取到请求所传递过来的参数
	param, err := cInitObj.getParams(c)
	if err != nil {
		log.Fatalf("获取到参数错误", err)
	}

	//2 执行整个控制器运行周期的before
	isContinue := cInitObj.initBeforeFn(param)
	if !isContinue {
		log.Fatalf("initBefore执行失败", err)
	}
	//3 对控制器的名字和方法名进行解析
	controllerAndFnName, err := cInitObj.splitControllerAndFnName(controllerFnCode)
	if err != nil {
		log.Fatalf("找不到控制器", err)
	}
	//获取到对应的控制器名字和方法名字
	var ControllerName = controllerAndFnName["controller"]
	var RuqFnName = controllerAndFnName["fnname"]

	//4 获取到对应配置文件的组件注册列表
	ModuleArr := config.MYCONFIG.Module
	if len(ModuleArr) == 0 {
		log.Fatalln("不存在注册的控制器")
	}
	//4 检测是否存在控制器
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
	//5 根据获取到控制器名字获取到对应类的实例
	initStruct, ok := newStruct(ControllerName) //返回核心的控制器
	if !ok {
		log.Fatalln("实例化核心的控制器失败")
	}
	inputs := make([]reflect.Value, 1) //参数
	inputs[0] = reflect.ValueOf(param) //输入对应的参数
	//6 根据类和方法 调用并执行
	ret := reflect.ValueOf(initStruct).MethodByName(RuqFnName).Call(inputs) //调用对应的方法
	//7 根据类型返回对应的结果
	c.JSON(200, gin.H{
		"status":  "posted",
		"message": ret,
		"nick":    "cc",
	})
	//反射映射实例话对应的类
	cInitObj.initAfterFn(param)
}

type ControllerInter interface {
	initBeforeFn() bool //在函数运行之前的函数
	initAfterFn()       //在函数运行之后的函数
	ReturnSuccess()     //返回结果-成功
	ReturnFail()        //返回结果-失败
}

//InitController ... 主控制器
type InitController struct {
}

//返回对应的map 和 错误 也就是获取到控制器的名字和对应调用的方法名字
//格式 为 User/addUser
func (cInit *InitController) splitControllerAndFnName(controllerFnCode string) (map[string]string, error) {
	if controllerFnCode == "" {
		return nil, errors.New("没有传递对应的控制器和方法代号")
	}
	strArr := strings.Split(controllerFnCode, "/")
	if len(strArr) != 2 {
		return nil, errors.New("传递的代码节点有问误,请用将控制器和方法隔开")
	}
	var controllerAndFnName map[string]string
	controllerAndFnName = make(map[string]string)
	controllerAndFnName["controller"] = strArr[0]
	controllerAndFnName["fnname"] = strArr[1]
	return controllerAndFnName, nil
}

//在调用某个子类之前会提前调用 可在这里写一些通用的逻辑
//适用于全部的基础类 必须返回true 只有返回true的情况才会继续运行
func (cInit *InitController) initBeforeFn(param interface{}) bool {
	return true
}

//在调用某个子类之前会提前调用 可在这里写一些通用的逻辑
//适用于全部的基础类 必须返回true 只有返回true的情况才会继续运行
func (cInit *InitController) initAfterFn(param interface{}) bool {
	return true
}

//获取到每次传递过来的参数
func (cInit *InitController) getParams(c *gin.Context) (interface{}, error) {
	return nil, nil
}

//向注册表中注册函数  该方法会在子控制器中调用
func registerType(elem interface{}) {
	//在根据Elem返回封装好的结构类型 elem的作用相当于返回一个格式化的内容
	//先用TypeOf获取到变量的类型
	t := reflect.TypeOf(elem).Elem() //获取到一个类型接口类型
	typeRegistry[t.Name()] = t       //将这个类型接口类型保存在注册表中并用对应的字符串映射起来
}

//根据方法名字动态的实例化结构体
func newStruct(name string) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	if !ok {
		return nil, false
	}
	//New 根据保存的类型接口 实例出一个具体的对象
	//elem 返回一个格式化可操作的对象 你可以理解链式调用的起点
	//Interface 返回一个具体的对象
	return reflect.New(elem).Elem().Interface(), true
}
