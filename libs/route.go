package libs

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"reflect"
	"strings"
)

var RegisterControllers *[]beego.ControllerInterface

// 不会添加到路由方法
var exceptMethods []string

func ExceptMethodAppend(excetMethod ...string) {
	if len(excetMethod) > 0 {
		for _, v := range excetMethod {
			exceptMethods = append(exceptMethods, v)
		}
	}
}

//将注册的contorller反射，获取对应的方法
//注意：犹豫目前不能明确地确定一个方法隶属于controller的方法，还是继承过来的方法
//所以这里采用了将beego框架的一些方法给过滤掉了，并过滤了一些其它的方法
//同时也删除了get,post,delete,put这四个方法，因为父类中有这四个方法
//所以：采用了在controller中定义一个特殊的方法RABCMethods()[]string，如果定义了，那么将会把这些方法添加到路由当中
func AllRoutes() map[string][]string {
	//定义一个临时的结构体，用户获取beego.Controller的全部方法
	type TempController struct {
		beego.Controller
	}

	tempRefectVal := reflect.ValueOf(&TempController{})
	tempRt := tempRefectVal.Type()

	var tempMethods []string
	for i := 0; i < tempRt.NumMethod(); i++ {
		tempMethods = append(tempMethods, tempRt.Method(i).Name)
	}

	routes := make(map[string][]string)

	//遍历注册的控制器，获取器方法
	for _, controller := range *RegisterControllers {
		reflectVal := reflect.ValueOf(controller)
		rt := reflectVal.Type()
		ct := reflect.Indirect(reflectVal).Type()
		controllerName := strings.TrimSuffix(ct.Name(), "Controller")

		var controllerRoutes []string
		//* 是一个特殊的权限，代表这个控制器下全部的权限
		controllerRoutes = append(controllerRoutes, "*")

		for i := 0; i < rt.NumMethod(); i++ {
			if !utils.InSlice(rt.Method(i).Name, exceptMethods) && !utils.InSlice(rt.Method(i).Name, tempMethods) {
				controllerRoutes = append(controllerRoutes, rt.Method(i).Name)
			}
		}

		routes[controllerName] = controllerRoutes
	}

	return routes
}
