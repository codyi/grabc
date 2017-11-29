package grbac

import (
	"fmt"
	"github.com/astaxie/beego"
	"grabc/controller"
)

func init() {
	beego.AutoRouter(&controller.RouteController{})
	beego.AutoRouter(&controller.RoleController{})
	beego.AutoRouter(&controller.PermissionController{})
	beego.AutoRouter(&controller.AssignmentController{})
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
