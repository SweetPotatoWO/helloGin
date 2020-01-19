package controller

import (
	"log"

	"github.com/gin-gonic/gin"
)

type ControllerInter interface {
	initBeforeFn() bool //在函数运行之前的函数
	Run()               //运行函数
	initAfterFn()       //在函数运行之后的函数
	ReturnSuccess()     //返回结果-成功
	ReturnFail()        //返回结果-失败
}

type InitController struct {
}

func (cInit *InitController) Run(c *gin.Engine, controllerFnCode string) {
	//1 先获取到请求所传递过来的参数
	param, err := cInit.getParams(c)
	if err != nil {
		log.Fatalf("获取到参数错误", err)
	}

	//2 执行整个控制器运行周期的before
	isContinue := cInit.initBeforeFn(param)
	if !isContinue {
		log.Fatalf("initBefore执行失败", err)
	}
	//3 对控制器的名字和方法名进行解析
	controllerAndFnName, err := cInit.splitControllerAndFnName(controllerFnCode)
	if err != nil {
		log.Fatalf("找不到控制器", err)
	}
	//实例化对应的控制器
	var ControllerName = controllerAndFnName["controller"]
	var RuqFnName = controllerAndFnName["fnname"]

	//
	cInit.initAfterFn(param)
}

//返回对应的map 和 错误 也就是获取到控制器的名字和对应调用的方法名字
//格式 为 User/addUser
func (cInit *InitController) splitControllerAndFnName(controllerFnCode string) (map[string]string, error) {

	return nil, nil
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
func (cInit *InitController) getParams(c *gin.Engine) (interface{}, error) {
	return nil, nil
}
