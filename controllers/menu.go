package controllers

import (
	"github.com/codyi/grabc/libs"
	"github.com/codyi/grabc/models"
	"github.com/codyi/grabc/views/menu"
	"strconv"
	"strings"
)

type MenuController struct {
	BaseController
}

//菜单管理首页
func (this *MenuController) Index() {
	page_index := strings.TrimSpace(this.GetString("page_index"))

	//分页设置
	pagination := libs.Pagination{}
	pagination.PageCount = 20
	pagination.Url = this.URLFor("MenuController.Index")

	if s, err := strconv.Atoi(page_index); err == nil {
		pagination.PageIndex = s
	} else {
		pagination.PageIndex = 1
	}

	menus, pageTotal, err := models.Menu{}.List(pagination.PageIndex, pagination.PageCount)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	pagination.PageTotal = pageTotal
	this.htmlData["menus"] = menus
	this.htmlData["pages"] = pagination
	this.AddBreadcrumbs("菜单管理", this.URLFor("MenuController.Index"))
	this.ShowHtml(&menu.Index{})
}

//添加菜单页面
func (this *MenuController) Post() {
	menuModel := models.Menu{}

	if this.isPost() {
		menu_name := strings.TrimSpace(this.GetString("menu_name"))
		menu_order := strings.TrimSpace(this.GetString("menu_order"))
		menu_route := strings.TrimSpace(this.GetString("menu_route"))
		menu_parent := strings.TrimSpace(this.GetString("menu_parent"))

		var int_menu_order, int_menu_parent int

		if i, err := strconv.Atoi(menu_order); err != nil {
			this.AddErrorMessage("排序必须是数字")
		} else {
			int_menu_order = i
		}

		if i, err := strconv.Atoi(menu_parent); err != nil {
			this.AddErrorMessage("父级分类必须是数字")
		} else {
			int_menu_parent = i
		}

		menuModel.Name = menu_name
		menuModel.Order = int_menu_order
		menuModel.Parent = int_menu_parent
		menuModel.Url = menu_route

		if isInsert, _ := menuModel.Insert(); isInsert {
			this.AddSuccessMessage("添加成功")
		} else {
			this.AddErrorMessage("添加失败")
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
	this.htmlData["model"] = menuModel
	this.htmlData["routes"] = selectRoutes
	this.htmlData["parents"] = selectParents
	this.AddBreadcrumbs("菜单管理", this.URLFor("MenuController.Index"))
	this.AddBreadcrumbs("新增", this.URLFor("MenuController.Add"))
	this.ShowHtml(&menu.Post{})
}

//修改菜单页面
func (this *MenuController) Put() {
	menuModel := models.Menu{}

	menu_id := strings.TrimSpace(this.GetString("menu_id"))

	if id, err := strconv.Atoi(menu_id); err == nil {
		if err := menuModel.FindById(id); err != nil {
			this.AddErrorMessage("数据获取失败")
		}
	} else {
		this.AddErrorMessage("数据不存在")
	}

	if this.isPost() && !menuModel.IsNewRecord() {
		menu_name := strings.TrimSpace(this.GetString("menu_name"))
		menu_order := strings.TrimSpace(this.GetString("menu_order"))
		menu_route := strings.TrimSpace(this.GetString("menu_route"))
		menu_parent := strings.TrimSpace(this.GetString("menu_parent"))

		var int_menu_order, int_menu_parent int

		if i, err := strconv.Atoi(menu_order); err != nil {
			this.AddErrorMessage("排序必须是数字")
		} else {
			int_menu_order = i
		}

		if i, err := strconv.Atoi(menu_parent); err != nil {
			this.AddErrorMessage("父级分类必须是数字")
		} else {
			int_menu_parent = i
		}

		menuModel.Name = menu_name
		menuModel.Order = int_menu_order
		menuModel.Parent = int_menu_parent
		menuModel.Url = menu_route

		if err := menuModel.Update(); err == nil {
			this.AddSuccessMessage("修改成功")
		} else {
			this.AddErrorMessage("修改失败")
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
	this.htmlData["model"] = menuModel
	this.htmlData["routes"] = selectRoutes
	this.htmlData["parents"] = selectParents
	this.AddBreadcrumbs("菜单管理", this.URLFor("MenuController.Index"))
	this.AddBreadcrumbs("修改", this.URLFor("MenuController.Put", "menu_id", menu_id))
	this.ShowHtml(&menu.Put{})
}

//删除菜单
func (this *MenuController) Delete() {
	data := JsonData{}
	if this.isPost() {
		menu_id, err := strconv.Atoi(strings.TrimSpace(this.GetString("menu_id")))

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
