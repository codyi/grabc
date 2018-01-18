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
	//注册rabc访问地址
	beego.AutoRouter(&controllers.RouteController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.PermissionController{})
	beego.AutoRouter(&controllers.AssignmentController{})
	beego.AutoRouter(&controllers.MenuController{})
	libs.IgnoreRoutes = make(map[string][]string, 0)
	RegisterController(&controllers.RouteController{}, &controllers.RoleController{}, &controllers.PermissionController{}, &controllers.AssignmentController{}, &controllers.MenuController{})

	//设置grabc页面路径
	//如果使用默认的，不要设置或者置空
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
	return libs.CheckAccess(controllerName, routeName, libs.AccessRoutes())
}

//没有权限跳转的页面
func Http_403(url string) {
	libs.Http_403 = url
}

//设置grabc的模板
func SetLayout(name string, path string) {
	libs.Template.LayoutName = name
	libs.Template.LayoutPath = path
	beego.AddViewPath(libs.Template.LayoutPath)
}

//设置grabc的模板的数据
func AddLayoutData(name string, value interface{}) {
	libs.Template.Data[name] = value
}

//返回用户可以看到的导航
func AccessMenus() []*libs.MenuGroup {
	return libs.AccessMenus()
}

//设置模板的路径
func SetViewPath(path string) {
	libs.Template.ViewPath = path
	beego.AddViewPath(libs.Template.ViewPath)
}
