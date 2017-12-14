package libs

import (
	"github.com/codyi/grabc/models"
	"strings"
)

//重新整理菜单，返还可以显示的菜单数据
type newMenu struct {
	Url  string
	Name string
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

	allAccessRoutes := AccessRoutes()
	type temp struct {
		Parent *models.Menu
		Child  []*models.Menu
	}

	tt := make(map[int]temp, 0)
	//归类子菜单，并检查完子菜单权限
	for _, m := range menus {
		if m.Parent == 0 {
			t := temp{}
			t.Parent = m
			tt[m.Id] = t
		} else {
			r := strings.Split(m.Route, "/")
			controllerName := r[0]
			routeName := r[1]
			if CheckAccess(controllerName, routeName, allAccessRoutes) {
				t := tt[m.Parent]
				if t.Parent == nil && len(t.Child) == 0 {
					t = temp{}
				}

				t.Child = append(t.Child, m)
				tt[m.Parent] = t
			}
		}
	}

	//检查完父级菜单权限，如果有子菜单，这个父级菜单将显示
	for i, t := range tt {
		if len(t.Child) > 0 {
			continue
		}

		r := strings.Split(t.Parent.Route, "/")
		controllerName := r[0]
		routeName := r[1]
		if !CheckAccess(controllerName, routeName, allAccessRoutes) {
			delete(tt, i)
		}
	}

	for _, t := range tt {
		m := MenuGroup{}
		p := newMenu{}
		p.Name = t.Parent.Name
		p.Url = "/" + t.Parent.Route

		m.Parent = p

		childMenus := make([]newMenu, 0)

		for _, m := range t.Child {
			nm := newMenu{}
			nm.Name = m.Name
			nm.Url = "/" + m.Route
			childMenus = append(childMenus, nm)
		}

		m.Child = childMenus
		returnMenus = append(returnMenus, &m)
	}

	return returnMenus
}
