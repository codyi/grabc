package grabc

import (
	"fmt"
	"github.com/astaxie/beego"
	"grabc/controllers"
	"grabc/libs"
)

var RegisterControllers []beego.ControllerInterface

func init() {
	beego.AutoRouter(&controllers.RouteController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.PermissionController{})
	beego.AutoRouter(&controllers.AssignmentController{})

	RegisterController(&controllers.RouteController{}, &controllers.RoleController{}, &controllers.PermissionController{}, &controllers.AssignmentController{})
	libs.RegisterControllers = &RegisterControllers
}

//注册需要检查的routes
func RegisterController(controllers ...beego.ControllerInterface) {
	for _, controller := range controllers {
		RegisterControllers = append(RegisterControllers, controller)
	}
}

//用户接口
type Identify interface {
	GetUid() int
}

type Rbac struct {
}

func (this *Rbac) CheckAccess() bool {
	fmt.Println(123456)
	return false
}
