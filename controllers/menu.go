package controllers

import (
	"github.com/codyi/grabc/libs"
	"github.com/codyi/grabc/models"
	"strings"
)

type MenuController struct {
	BaseController
}

//菜单管理首页
func (this *MenuController) Index() {
	page_index, err := this.GetInt("page_index")

	//分页设置
	pagination := libs.Pagination{}
	pagination.PageCount = 20
	pagination.Url = this.URLFor("MenuController.Index")

	if err == nil {
		pagination.PageIndex = page_index
	} else {
		pagination.PageIndex = 1
	}

	menus, pageTotal, err := models.Menu{}.List(pagination.PageIndex, pagination.PageCount)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	pagination.PageTotal = pageTotal
	this.Data["menulist"] = menus
	this.Data["pages"] = pagination
	this.AddBreadcrumbs("菜单管理", this.URLFor("MenuController.Index"))
	this.ShowHtml()
}

//添加菜单页面
func (this *MenuController) Post() {
	menuModel := models.Menu{}

	if this.isPost() {
		menu_name := strings.TrimSpace(this.GetString("menu_name"))
		int_menu_order, menu_order_err := this.GetInt("menu_order")
		menu_route := strings.TrimSpace(this.GetString("menu_route"))
		int_menu_parent, menu_parent_err := this.GetInt("menu_parent")
		menu_icon := strings.TrimSpace(this.GetString("menu_icon"))

		if menu_order_err != nil {
			this.redirectMessage(this.URLFor("MenuController.Post"), "排序必须是数字", MESSAGE_TYPE_ERROR)
		}

		if menu_parent_err != nil {
			this.redirectMessage(this.URLFor("MenuController.Post"), "父级分类必须是数字", MESSAGE_TYPE_ERROR)
		}

		menuModel.Name = menu_name
		menuModel.Order = int_menu_order
		menuModel.Parent = int_menu_parent
		menuModel.Url = menu_route
		menuModel.Icon = menu_icon

		if isInsert, _ := menuModel.Insert(); isInsert {
			this.redirectMessage(this.URLFor("MenuController.Index"), "添加成功", MESSAGE_TYPE_SUCCESS)
		} else {
			this.redirectMessage(this.URLFor("MenuController.Post"), "添加失败", MESSAGE_TYPE_ERROR)
		}
	}

	routeModel := models.Route{}
	routes, _ := routeModel.FindAll()
	var selectRoutes []string

	if routes != nil {
		for _, route := range routes {
			if !strings.Contains(route.Url, "*") {
				selectRoutes = append(selectRoutes, route.Url)
			}
		}
	}

	parents, _ := menuModel.FindAllParent()
	selectParents := make(map[int]string, 0)

	if parents != nil {
		for _, p := range parents {
			selectParents[p.Id] = p.Name
		}
	}
	this.Data["model"] = menuModel
	this.Data["routes"] = selectRoutes
	this.Data["parents"] = selectParents
	this.AddBreadcrumbs("菜单管理", this.URLFor("MenuController.Index"))
	this.AddBreadcrumbs("新增", this.URLFor("MenuController.Add"))
	this.ShowHtml()
}

//修改菜单页面
func (this *MenuController) Put() {
	menu_id, err := this.GetInt("menu_id")

	if err != nil {
		this.redirectMessage(this.URLFor("MenuController.Index"), "数据不存在", MESSAGE_TYPE_ERROR)
	}

	menuModel := models.Menu{}
	if err = menuModel.FindById(menu_id); err != nil {
		this.redirectMessage(this.URLFor("MenuController.Index"), "数据获取失败", MESSAGE_TYPE_ERROR)
	}

	if this.isPost() {
		menu_name := strings.TrimSpace(this.GetString("menu_name"))
		int_menu_order, menu_order_err := this.GetInt("menu_order")
		menu_route := strings.TrimSpace(this.GetString("menu_route"))
		int_menu_parent, menu_parent_err := this.GetInt("menu_parent")
		menu_icon := strings.TrimSpace(this.GetString("menu_icon"))

		if menu_order_err != nil {
			this.redirectMessage(this.URLFor("MenuController.Put", "menu_id", this.GetString("menu_id")), "排序必须是数字", MESSAGE_TYPE_ERROR)
		}

		if menu_parent_err != nil {
			this.redirectMessage(this.URLFor("MenuController.Put", "menu_id", this.GetString("menu_id")), "父级分类必须是数字", MESSAGE_TYPE_ERROR)
		}

		menuModel.Name = menu_name
		menuModel.Order = int_menu_order
		menuModel.Parent = int_menu_parent
		menuModel.Url = menu_route
		menuModel.Icon = menu_icon

		if err := menuModel.Update(); err == nil {
			this.redirectMessage(this.URLFor("MenuController.Index"), "修改成功", MESSAGE_TYPE_SUCCESS)
		} else {
			this.redirectMessage(this.URLFor("MenuController.Put", "menu_id", this.GetString("menu_id")), "修改失败", MESSAGE_TYPE_ERROR)
		}

	}

	routeModel := models.Route{}
	routes, _ := routeModel.FindAll()
	var selectRoutes []string

	if routes != nil {
		for _, route := range routes {
			if !strings.Contains(route.Url, "*") {
				selectRoutes = append(selectRoutes, route.Url)
			}
		}
	}

	parents, _ := menuModel.FindAllParent()
	selectParents := make(map[int]string, 0)

	if parents != nil {
		for _, p := range parents {
			selectParents[p.Id] = p.Name
		}
	}
	this.Data["model"] = menuModel
	this.Data["routes"] = selectRoutes
	this.Data["parents"] = selectParents
	this.AddBreadcrumbs("菜单管理", this.URLFor("MenuController.Index"))
	this.AddBreadcrumbs("修改", this.URLFor("MenuController.Put", "menu_id", this.GetString("menu_id")))
	this.ShowHtml()
}

//删除菜单
func (this *MenuController) Delete() {
	data := JsonData{}
	if this.isPost() {
		menu_id, err := this.GetInt("menu_id")

		if err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
		}

		menuModel := models.Menu{}
		if err := menuModel.FindById(menu_id); err != nil {
			data.Code = 400
			data.Message = "数据获取失败"
			this.ShowJSON(&data)
		}

		if is_delete, err := menuModel.Delete(); is_delete {
			data.Code = 200
			data.Message = "删除成功"
			this.ShowJSON(&data)
		} else {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
		}
	} else {
		data.Code = 400
		data.Message = "非法请求"
	}

	this.ShowJSON(&data)
}
