package controller

import (
	"encoding/json"
	"errors"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

var typeRegistry = make(map[string]reflect.Type) //组件注册列表

//InitController ... 主控制器
type InitController struct {
}

//在调用某个子类之前会提前调用 可在这里写一些通用的逻辑
//适用于全部的基础类 必须返回true 只有返回true的情况才会继续运行
func (cInit *InitController) InitBeforeFn(param interface{}) bool {
	return true
}

//在调用某个子类之前会提前调用 可在这里写一些通用的逻辑
//适用于全部的基础类 必须返回true 只有返回true的情况才会继续运行
func (cInit *InitController) InitAfterFn(param interface{}) bool {
	return true
}

//返回对应的map 和 错误 也就是获取到控制器的名字和对应调用的方法名字
//格式 为 User/addUser
func (cInit *InitController) SplitControllerAndFnName(controllerFnCode string) (map[string]string, error) {
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

//获取到每次传递过来的参数
func (cInit *InitController) GetParams(c *gin.Context) (interface{}, error) {
	var a map[string]interface{}
	return a, nil
}

//向注册表中注册函数  该方法会在子控制器中调用
func RegisterType(elem interface{}) {
	//在根据Elem返回封装好的结构类型 elem的作用相当于返回一个格式化的内容
	//先用TypeOf获取到变量的类型
	t := reflect.TypeOf(elem).Elem() //获取到一个类型接口类型
	typeRegistry[t.Name()] = t       //将这个类型接口类型保存在注册表中并用对应的字符串映射起来
}

//根据方法名字动态的实例化结构体
func NewStruct(name string) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	if !ok {
		return nil, false
	}
	//New 根据保存的类型接口 实例出一个具体的对象
	//elem 返回一个格式化可操作的对象 你可以理解链式调用的起点
	//Interface 返回一个具体的对象
	return reflect.New(elem).Elem().Interface(), true
}

//ReturnJosnSuccess... 返回正确的结果
//第一参数为信息,第二个参数为数据
func ReturnJsonSuccess(msg string, returnData interface{}) string {
	json, _ := json.Marshal(ResponseResult{
		ErrorCode:  0,
		ErrorMsg:   "",
		SuccessMsg: msg,
		Data:       returnData,
	})
	return string(json)
}

//ReturnJosnError... 返回正确的结果
//失败不能返回数据
func ReturnJsonError(msg string) string {
	var returnData []int = make([]int, 2)
	json, _ := json.Marshal(ResponseResult{
		ErrorCode:  1,
		ErrorMsg:   msg,
		SuccessMsg: "",
		Data:       returnData,
	})

	return string(json)
}

type ResponseResult struct {
	ErrorCode  int32       `json:"errorCode"`
	ErrorMsg   string      `json:"errorMsg"`
	SuccessMsg string      `json:"successMsg"`
	Data       interface{} `json:"data"`
}
