package grbac

import (
	"fmt"
	"github.com/astaxie/beego"
	"grabc/controllers"
)

func init() {
	beego.AutoRouter(&controllers.RouteController{})
	beego.AutoRouter(&controllers.RoleController{})
	beego.AutoRouter(&controllers.PermissionController{})
	beego.AutoRouter(&controllers.AssignmentController{})
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
