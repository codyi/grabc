package grabc

import (
	"github.com/astaxie/beego"
	"grabc/controllers"
	"grabc/libs"
	"grabc/models"
)

//init function
func init() {
	//注册rabc访问地址
	beego.AutoRouter(&controllers.RouteController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.PermissionController{})
	beego.AutoRouter(&controllers.AssignmentController{})

	libs.IgnoreRoutes = make(map[string][]string, 0)
}

//注册需要检查的routes
func RegisterController(controllers ...beego.ControllerInterface) {
	for _, controller := range controllers {
		libs.RegisterControllers = append((libs.RegisterControllers), controller)
	}
}

//注册登录用户
func RegisterIdentify(i libs.IUserIdentify) {
	libs.Identify = &i
}

//注册用户的模型表，用于获取用户的id和用户名称
func RegisterUserModel(m models.IUserModel) {
	models.UserModel = &m
}

//增加忽律检查的地址，例如site/login,这个就不需要检查权限
func AppendIgnoreRoute(c, r string) {
	if libs.IgnoreRoutes[c] == nil {
		libs.IgnoreRoutes[c] = make([]string, 0)
	}

	libs.IgnoreRoutes[c] = append(libs.IgnoreRoutes[c], r)
}

//权限检查
func CheckAccess(controllerName, routeName string) bool {
	return libs.CheckAccess(controllerName, routeName)
}

//没有权限跳转的页面
func NoPermissionUrl(url string) {
	libs.NoPermissionUrl = url
}
