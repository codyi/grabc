package libs

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/utils"
	"github.com/codyi/grabc/models"
	"strings"
)

//用于保存注册的beego.Controller
var BeegoC *beego.Controller

var RegisterUserIdFunc func(c *beego.Controller) int

//忽律检查权限的网址
var IgnoreRoutes map[string][]string = make(map[string][]string, 0)

//获取用户全部可以访问的路由
func AccessRoutes() map[string][]string {
	allAccessRoutes := make(map[string][]string, 0)

	//如果存在用户ID，则获取用户对应的权限
	if RegisterUserIdFunc != nil && RegisterUserIdFunc(BeegoC) > 0 {
		allAccessRoutes = models.Route{}.ListByUserId(RegisterUserIdFunc(BeegoC))
	}

	//将忽律检查的路由和可以登录的路由合并
	for controller, routes := range IgnoreRoutes {
		if allAccessRoutes[controller] == nil {
			allAccessRoutes[controller] = routes
		} else {
			for _, r := range routes {
				allAccessRoutes[controller] = append(allAccessRoutes[controller], r)
			}
		}

	}

	return allAccessRoutes
}

//权限检查
func CheckAccess(controllerName, routeName string, allAccessRoutes map[string][]string) bool {

	//检查路由
	controllerName = strings.ToLower(controllerName)
	routeName = strings.ToLower(routeName)
	if allAccessRoutes[controllerName] != nil {
		if utils.InSlice(routeName, allAccessRoutes[controllerName]) {
			return true
		}

		if utils.InSlice("*", allAccessRoutes[controllerName]) {
			return true
		}
	}

	//如果用户存在*的权限，将可以进入全部页面
	if allAccessRoutes["*"] != nil && utils.InSlice("*", allAccessRoutes["*"]) {
		return true
	}

	return false
}
