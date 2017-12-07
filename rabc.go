package grabc

import (
	"github.com/astaxie/beego"
	"grabc/controllers"
	"grabc/libs"
	"grabc/models"
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

//忽律检查的网址
type ignoreRoute struct {
	Controller string
	Route      string
}

var ignoreRoutes []ignoreRoute

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
	ignoreRoutes = make([]ignoreRoute, 0)
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
	i := ignoreRoute{}
	i.Controller = c
	i.Route = r
	ignoreRoutes = append(ignoreRoutes, i)
}

//权限检查
func CheckAccess(controllerName, routeName string) bool {
	//先检查是否在忽律的路由中
	if len(ignoreRoutes) > 0 {
		// for _, ignoreRoute := range ignoreRoutes {
		// 	return true
		// }
	}

	if identify.GetId() <= 0 {
		return false
	}

	return false
}
