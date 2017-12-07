package grabc

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"grabc/controllers"
	"grabc/libs"
	"grabc/models"
	"strings"
)

//用于定义用户接口
type IUserIdentify interface {
	GetId() int //返回当前登录用户的ID
}

//用于检测需要权限检查的的路由
var registerControllers []beego.ControllerInterface

//用于保存外部注册的登录用户实例
var identify IUserIdentify

//用于保存外部注册的用户模型实例
var userModel models.IUserModel

//忽律检查权限的网址
var ignoreRoutes map[string][]string

//init function
func init() {
	//注册rabc访问地址
	beego.AutoRouter(&controllers.RouteController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.PermissionController{})
	beego.AutoRouter(&controllers.AssignmentController{})

	RegisterController(&controllers.RouteController{}, &controllers.RoleController{}, &controllers.PermissionController{}, &controllers.AssignmentController{})
	libs.RegisterControllers = &registerControllers
	models.UserModel = &userModel
	ignoreRoutes = make(map[string][]string, 0)
}

//注册需要检查的routes
func RegisterController(controllers ...beego.ControllerInterface) {
	for _, controller := range controllers {
		registerControllers = append(registerControllers, controller)
	}
}

//注册登录用户
func RegisterIdentify(i IUserIdentify) {
	identify = i
}

//注册用户的模型表，用于获取用户的id和用户名称
func RegisterUserModel(m models.IUserModel) {
	userModel = m
}

//增加忽律检查的地址，例如site/login,这个就不需要检查权限
func AppendIgnoreRoute(c, r string) {
	if ignoreRoutes[c] == nil {
		ignoreRoutes[c] = make([]string, 0)
	}

	ignoreRoutes[c] = append(ignoreRoutes[c], r)
}

//权限检查
func CheckAccess(controllerName, routeName string) bool {
	allAccessRoutes := make(map[string][]string, 0)

	if identify.GetId() > 0 {
		allAccessRoutes = models.Route{}.ListByUserId(identify.GetId())
	}

	for controller, routes := range ignoreRoutes {
		if allAccessRoutes[controller] == nil {
			allAccessRoutes[controller] = routes
		} else {
			for _, r := range routes {
				allAccessRoutes[controller] = append(allAccessRoutes[controller], r)
			}
		}

	}

	controllerName = strings.ToLower(controllerName)
	routeName = strings.ToLower(routeName)
	fmt.Println(controllerName)
	fmt.Println(routeName)
	fmt.Println(allAccessRoutes)
	if allAccessRoutes[controllerName] != nil {
		if utils.InSlice(routeName, allAccessRoutes[controllerName]) {
			return true
		}

		if utils.InSlice("*", allAccessRoutes[controllerName]) {
			return true
		}
	}

	if allAccessRoutes["*"] != nil {
		return true
	}

	return false
}
