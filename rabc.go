package grabc

import (
	"fmt"
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
var RegisterControllers []beego.ControllerInterface

//用于保存外部注册的登录用户实例
var Identify IUserIdentify

//用于保存外部注册的用户模型实例
var UserModel models.IUserModel

//init function
func init() {
	//注册rabc访问地址
	beego.AutoRouter(&controllers.RouteController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.PermissionController{})
	beego.AutoRouter(&controllers.AssignmentController{})

	RegisterController(&controllers.RouteController{}, &controllers.RoleController{}, &controllers.PermissionController{}, &controllers.AssignmentController{})
	libs.RegisterControllers = &RegisterControllers
	models.UserModel = &UserModel
}

//注册需要检查的routes
func RegisterController(controllers ...beego.ControllerInterface) {
	for _, controller := range controllers {
		RegisterControllers = append(RegisterControllers, controller)
	}
}

//权限检查
func CheckAccess() bool {
	fmt.Println(123456)
	return false
}
