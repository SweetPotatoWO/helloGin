package config

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

//Myconf ...
type Myconf struct {
	Servicename string
	Port        string
	Db          []Device
}

//Device ...
type Device struct {
	Dbid string
	Type string
	Port string
	Host string
	User string
	Pwd  string
}

var main_config_path = "/root/git/helloGin/config/config.yml"
var constant_config_path = "/root/git/helloGin/config/constant.json"

//MYCONFIG  全局变量
var MYCONFIG Myconf

//MYCONSTANT 赋值一个任意类型的全局变量
var MYCONSTANT interface{}

func init() {
	//加载服务器的常规配置
	yamlFile, err := ioutil.ReadFile(main_config_path)
	if err != nil {
		log.Fatalf("加载主配置文件失败 %v", err)
	}
	err = yaml.Unmarshal(yamlFile, &MYCONFIG)
	if err != nil {
		log.Fatalf("文件解析失败 : %v", err)
	}
	//加载常量配置
	bytecodes, err := ioutil.ReadFile(constant_config_path)
	if err != nil {
		log.Fatalf("加载常量配置文件失败")
	}
	err = json.Unmarshal(bytecodes, &MYCONSTANT)
}
