package libs

import (
	"github.com/astaxie/beego/utils"
	"grabc/models"
	"strings"
)

//用于定义用户接口
type IUserIdentify interface {
	GetId() int //返回当前登录用户的ID
}

//用于保存外部注册的登录用户实例
var Identify *IUserIdentify

//忽律检查权限的网址
var IgnoreRoutes map[string][]string

//404页面网址
var NoPermissionUrl string = ""

//权限检查
func CheckAccess(controllerName, routeName string) bool {
	allAccessRoutes := make(map[string][]string, 0)

	//如果存在用户ID，则获取用户对应的权限
	if (*Identify).GetId() > 0 {
		allAccessRoutes = models.Route{}.ListByUserId((*Identify).GetId())
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
