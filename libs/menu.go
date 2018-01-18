package libs

import (
	"fmt"
	"github.com/codyi/grabc/models"
	"strings"
)

//重新整理菜单，返还可以显示的菜单数据
type newMenu struct {
	Url  string
	Name string
	Icon string
}

type MenuGroup struct {
	Parent newMenu
	Child  []newMenu
}

//获取用户可以访问的菜单
func AccessMenus() []*MenuGroup {
	returnMenus := make([]*MenuGroup, 0)
	menu := models.Menu{}
	menus, err := menu.ListAll()

	if err != nil {
		return returnMenus
	}

	//找出全部的子菜单、父级菜单
	parentMenus := make([]*models.Menu, 0)
	childMenus := make([]*models.Menu, 0)
	for _, menu := range menus {
		if menu.Parent == 0 {
			parentMenus = append(parentMenus, menu)
		} else {
			childMenus = append(childMenus, menu)
		}
	}
	//归类菜单,并检查权限
	allAccessRoutes := AccessRoutes()

	for _, parentMenu := range parentMenus {
		mg := MenuGroup{}
		mg.Child = make([]newMenu, 0)

		//查找子菜单、并检查权限
		for _, childMenu := range childMenus {
			if childMenu.Parent != parentMenu.Id {
				continue
			}

			r := strings.Split(childMenu.Url, "/")
			controllerName := r[0]
			routeName := r[1]

			if CheckAccess(controllerName, routeName, allAccessRoutes) {
				cm := newMenu{}
				cm.Name = childMenu.Name
				cm.Url = "/" + childMenu.Url
				cm.Icon = childMenu.Icon
				mg.Child = append(mg.Child, cm)
			}
		}

		//如果不存在子菜单，将检查父级菜单的权限
		if len(mg.Child) == 0 {
			r := strings.Split(parentMenu.Url, "/")
			controllerName := r[0]
			routeName := r[1]
			if CheckAccess(controllerName, routeName, allAccessRoutes) {
				mg.Parent = newMenu{}
				mg.Parent.Name = parentMenu.Name
				mg.Parent.Url = "/" + parentMenu.Url
				mg.Parent.Icon = parentMenu.Icon
			}
		} else {
			mg.Parent = newMenu{}
			mg.Parent.Name = parentMenu.Name
			mg.Parent.Url = "/" + parentMenu.Url
			mg.Parent.Icon = parentMenu.Icon
		}

		returnMenus = append(returnMenus, &mg)
	}

	return returnMenus
}

func ShowMenu(controllName, actionName string) string {
	var activeUrl = ""
	menus := AccessMenus()
	//先进行精准的匹配选中的url链接
	for _, menu := range menus {
		if len(menu.Child) > 0 {
			for _, childMenu := range menu.Child {
				if strings.ToLower("/"+controllName+"/"+actionName) == strings.ToLower(childMenu.Url) {
					activeUrl = childMenu.Url
					goto SHOW
				}
			}
		} else {
			if strings.ToLower("/"+controllName+"/"+actionName) == strings.ToLower(menu.Parent.Url) {
				activeUrl = menu.Parent.Url
				goto SHOW
			}
		}
	}

	//如果精准没有找到，则尝试模糊查询
	//规则：只匹配controller的名称
	for _, menu := range menus {
		if len(menu.Child) > 0 {
			for _, childMenu := range menu.Child {
				if strings.Index(childMenu.Url, controllName) == 1 {
					activeUrl = childMenu.Url
					goto SHOW
				}
			}
		} else {
			if strings.Index(menu.Parent.Url, controllName) == 1 {
				activeUrl = menu.Parent.Url
				goto SHOW
			}
		}
	}
SHOW:

	html := `<ul class='sidebar-menu tree' data-widget='tree'>`
	for _, menu := range menus {
		if len(menu.Child) > 0 {
			childHtml := ""
			isActiveChild := false
			for _, childMenu := range menu.Child {
				activeClass := ""

				if activeUrl == childMenu.Url {
					activeClass = "active"
					isActiveChild = true
				}

				childHtml += fmt.Sprintf(`<li class="%s"><a href='%s'><i class="fa %s"></i>%s</a></li>`, activeClass, childMenu.Url, childMenu.Icon, childMenu.Name)
			}

			s := `<li class='treeview %s'><a href='#'><i class="fa %s"></i><span>%s</span><span class='pull-right-container'><i class='fa fa-angle-left pull-right'></i></span></a><ul class='treeview-menu'>%s</ul></li>`
			if isActiveChild {
				html += fmt.Sprintf(s, "active menu-open", menu.Parent.Icon, menu.Parent.Name, childHtml)
			} else {
				html += fmt.Sprintf(s, "", menu.Parent.Icon, menu.Parent.Name, childHtml)
			}
		} else {
			activeClass := ""
			s := `<li class='%s'><a href='%s'><i class="fa %s"></i><span>%s</span></a></li>`

			if activeUrl == menu.Parent.Url {
				activeClass = "active"
			}

			html += fmt.Sprintf(s, activeClass, menu.Parent.Url, menu.Parent.Icon, menu.Parent.Name)
		}

	}

	html += `</ul>`
	return html
}
