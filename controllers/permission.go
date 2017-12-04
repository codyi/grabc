package controllers

import (
	"github.com/astaxie/beego/utils"
	"grabc/libs"
	"grabc/models"
	"grabc/views/permission"
	"strconv"
	"strings"
)

type PermissionController struct {
	BaseController
}

//permision index page
func (this *PermissionController) Index() {
	page_index := strings.TrimSpace(this.GetString("page_index"))

	pagination := libs.Pagination{}
	pagination.PageCount = 20
	pagination.Url = this.URLFor("PermissionController.Index")

	if s, err := strconv.Atoi(page_index); err == nil {
		pagination.PageIndex = s
	} else {
		pagination.PageIndex = 1
	}

	permissions, pageTotal, err := models.Permission{}.FindAll(pagination.PageIndex, pagination.PageCount)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	pagination.PageTotal = pageTotal
	this.htmlData["permissions"] = permissions
	this.htmlData["pages"] = pagination
	this.AddBreadcrumbs("权限管理", this.URLFor("PermissionController.Index"))
	this.ShowHtml(&permission.Index{})
}

//permision add page
func (this *PermissionController) Add() {
	permissionModel := models.Permission{}

	if this.isPost() {
		permission_name := strings.TrimSpace(this.GetString("permission_name"))
		permission_desc := strings.TrimSpace(this.GetString("permission_desc"))

		permissionModel.Name = permission_name
		permissionModel.Description = permission_desc

		if isInsert, _ := permissionModel.Insert(); isInsert {
			this.AddSuccessMessage("添加成功")
		} else {
			this.AddErrorMessage("添加失败")
		}

	}

	this.htmlData["model"] = permissionModel
	this.AddBreadcrumbs("权限管理", this.URLFor("PermissionController.Index"))
	this.AddBreadcrumbs("新增", this.URLFor("PermissionController.Add"))
	this.ShowHtml(&permission.Add{})
}

//permision update page
func (this *PermissionController) Put() {
	permissionModel := models.Permission{}
	permission_id := strings.TrimSpace(this.GetString("permission_id"))

	if id, err := strconv.Atoi(permission_id); err == nil {
		if err := permissionModel.FindById(id); err != nil {
			this.AddErrorMessage("数据获取失败")
		}
	} else {
		this.AddErrorMessage("数据不存在")
	}

	if this.isPost() && !permissionModel.IsNewRecord() {
		permission_name := strings.TrimSpace(this.GetString("permission_name"))
		permission_desc := strings.TrimSpace(this.GetString("permission_desc"))

		permissionModel.Name = permission_name
		permissionModel.Description = permission_desc

		if err := permissionModel.Update(); err == nil {
			this.AddSuccessMessage("修改成功")
		} else {
			this.AddErrorMessage("修改失败")
		}

	}

	this.htmlData["model"] = permissionModel
	this.AddBreadcrumbs("权限管理", this.URLFor("PermissionController.Index"))
	this.AddBreadcrumbs("修改", this.URLFor("PermissionController.Put", "permission_id", permission_id))
	this.AddBreadcrumbs(permissionModel.Name, "")
	this.ShowHtml(&permission.Update{})
}

//permision view page
func (this *PermissionController) Assignment() {
	permissionModel := models.Permission{}
	permission_id := strings.TrimSpace(this.GetString("permission_id"))

	if id, err := strconv.Atoi(permission_id); err == nil {
		if err := permissionModel.FindById(id); err != nil {
			this.AddErrorMessage("数据获取失败")
		}
	} else {
		this.AddErrorMessage("数据不存在")
	}

	if permissionModel.IsNewRecord() {
		this.AddErrorMessage("数据不存在")

	}

	//获取已经授权的路由
	routeModel := models.Route{}
	permissionRouteModel := models.PermissionRoute{}
	allPermissionRouteModels, err := permissionRouteModel.FindAllByPermissionId(permissionModel.Id)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	var assignmentRoutes []string
	if allPermissionRouteModels != nil {
		var allPermissionRouteIds []int

		for _, pr := range allPermissionRouteModels {
			allPermissionRouteIds = append(allPermissionRouteIds, pr.RouteId)
		}

		prs, _ := routeModel.FindAllByIds(allPermissionRouteIds)

		if prs != nil {
			for _, r := range prs {
				assignmentRoutes = append(assignmentRoutes, r.Route)
			}
		}
	}

	//获取全部可以授权的路由
	routes, _ := routeModel.FindAll()
	var allRoutes []string

	if routes != nil {
		for _, route := range routes {
			if !utils.InSlice(route.Route, assignmentRoutes) {
				allRoutes = append(allRoutes, route.Route)
			}
		}
	}

	this.htmlData["model"] = permissionModel
	this.htmlData["allRoutes"] = allRoutes
	this.htmlData["assignmentRoutes"] = assignmentRoutes
	this.AddBreadcrumbs("权限管理", this.URLFor("PermissionController.Index"))
	this.AddBreadcrumbs("授权", this.URLFor("PermissionController.Assignment", "permission_id", permission_id))
	this.AddBreadcrumbs(permissionModel.Name, "")
	this.ShowHtml(&permission.Assignment{})
}

//route permission ajax add page
func (this *PermissionController) AssignmentRoute() {
	data := JsonData{}

	if this.isPost() {
		route := strings.TrimSpace(this.GetString("route"))
		permissionId := strings.TrimSpace(this.GetString("permissionId"))

		routeModel := models.Route{}
		permissionRouteModel := models.PermissionRoute{}

		if route != "" {
			routeModel.FindByRoute(route)

			if routeModel.Id > 0 {
				permissionRouteModel.RouteId = routeModel.Id
			} else {
				data.Code = 400
				data.Message = "路由不存在"
				this.ShowJSON(&data)
				return
			}
		} else {
			data.Code = 400
			data.Message = "路由不能为空"
			this.ShowJSON(&data)
			return
		}

		if i, err := strconv.Atoi(permissionId); err == nil {
			permissionRouteModel.PermissionId = i
		} else {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
			return
		}

		if isInsert, err := permissionRouteModel.Insert(); isInsert {
			data.Code = 200
			data.Message = "添加成功"
			data.Data = make(map[string]interface{})
			data.Data["route"] = routeModel.Route
		} else {
			data.Code = 400
			data.Message = err.Error()
		}

	} else {
		data.Code = 400
		data.Message = "非法请求"
	}

	this.ShowJSON(&data)
}

//route permission ajax add page
func (this *PermissionController) RemoveRoute() {
	data := JsonData{}

	if this.isPost() {
		param_route := strings.TrimSpace(this.GetString("route"))
		param_permision_id := strings.TrimSpace(this.GetString("permissionId"))
		permissionRouteModel := models.PermissionRoute{}

		var int_param_route_id, int_param_permission_id int
		routeModel := models.Route{}

		if param_route != "" {
			routeModel.FindByRoute(param_route)

			if routeModel.Id > 0 {
				int_param_route_id = routeModel.Id
			} else {
				data.Code = 400
				data.Message = "路由不存在"
				this.ShowJSON(&data)
				return
			}
		} else {
			data.Code = 400
			data.Message = "路由不能为空"
			this.ShowJSON(&data)
			return
		}

		if id, err := strconv.Atoi(param_permision_id); err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
			return
		} else {
			int_param_permission_id = id
		}

		if is_delete, err := permissionRouteModel.Delete(int_param_route_id, int_param_permission_id); is_delete {
			data.Code = 200
			data.Message = "删除成功"
			data.Data = make(map[string]interface{})
			data.Data["route"] = routeModel.Route
		} else {
			data.Code = 400
			data.Message = err.Error()
		}

	} else {
		data.Code = 400
		data.Message = "非法请求"
	}

	this.ShowJSON(&data)
}

//permision delete page
func (this *PermissionController) Delete() {
	data := JsonData{}
	if this.isPost() {
		permision_id, err := strconv.Atoi(strings.TrimSpace(this.GetString("permission_id")))

		if err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
			return
		}

		permissionModel := models.Permission{}
		if err := permissionModel.FindById(permision_id); err != nil {
			data.Code = 400
			data.Message = "数据获取失败"
			this.ShowJSON(&data)
			return
		}

		permissionRouteModel := models.PermissionRoute{}
		err = permissionRouteModel.DeleteByPermissionId(permissionModel.Id)

		if err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
			return
		}

		if is_delete, err := permissionModel.Delete(); is_delete {
			data.Code = 200
			data.Message = "删除成功"
			this.ShowJSON(&data)
			return
		} else {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
			return
		}
	} else {
		data.Code = 400
		data.Message = "非法请求"
	}

	this.ShowJSON(&data)
}
