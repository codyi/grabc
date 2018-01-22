package grabc

import (
	"github.com/astaxie/beego"
	"github.com/codyi/grabc/controllers"
	"github.com/codyi/grabc/libs"
	"github.com/codyi/grabc/models"
	"os"
	"path/filepath"
)

//init function
func init() {
	var c []beego.ControllerInterface
	c = append(c, &controllers.RouteController{}, &controllers.RoleController{}, &controllers.PermissionController{}, &controllers.AssignmentController{}, &controllers.MenuController{})

	for _, v := range c {
		//将路由注册到beego
		beego.AutoRouter(v)
		//将路由注册到grabc
		RegisterController(v)
	}

	//设置grabc默认视图、layout路径
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err.Error())
	}

	SetViewPath(filepath.Dir(dir) + "/github.com/codyi/grabc/views/")
	SetLayout("main.html", filepath.Dir(dir)+"/github.com/codyi/grabc/views/layout/")
}

//注册需要检查的routes
func RegisterController(controllers ...beego.ControllerInterface) {
	for _, controller := range controllers {
		libs.RegisterControllers = append((libs.RegisterControllers), controller)
	}
}

//注册用于获取登录用户ID的函数
func RegisterUserIdFunc(f func(c *beego.Controller) int) {
	libs.RegisterUserIdFunc = f
}

//注册当前请求beego的Controller
func SetBeegoController(c *beego.Controller) {
	libs.BeegoC = c
}

//注册用户的模型表，用于获取用户的id和用户名称
func RegisterUserModel(m models.IUserModel) {
	models.UserModel = &m
}

//增加忽律检查的地址，例如site/login,这个就不需要检查权限
func AppendIgnoreRoute(controllerName, actionName string) {
	if libs.IgnoreRoutes[controllerName] == nil {
		libs.IgnoreRoutes[controllerName] = make([]string, 0)
	}

	libs.IgnoreRoutes[controllerName] = append(libs.IgnoreRoutes[controllerName], actionName)
}

//权限检查
func CheckAccess(controllerName, routeName string) bool {
	return libs.CheckAccess(controllerName, routeName, libs.AccessRoutes())
}

//设置grabc的模板
func SetLayout(layoutName string, layoutPath string) {
	libs.Template.LayoutName = layoutName
	libs.Template.LayoutPath = layoutPath
	beego.AddViewPath(libs.Template.LayoutPath)
}

//设置grabc的模板的数据
func AddLayoutData(name string, value interface{}) {
	libs.Template.Data[name] = value
}

//设置grabc模板的路径
func SetViewPath(path string) {
	libs.Template.ViewPath = path
	beego.AddViewPath(libs.Template.ViewPath)
}
