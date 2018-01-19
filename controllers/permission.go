package controllers

import (
	"github.com/astaxie/beego/utils"
	"github.com/codyi/grabc/libs"
	"github.com/codyi/grabc/models"
	"strings"
)

type PermissionController struct {
	BaseController
}

//permision index page
func (this *PermissionController) Index() {
	page_index, err := this.GetInt("page_index")

	pagination := libs.Pagination{}
	pagination.PageCount = 20
	pagination.Url = this.URLFor("PermissionController.Index")

	if err == nil {
		pagination.PageIndex = page_index
	} else {
		pagination.PageIndex = 1
	}

	permissions, pageTotal, err := models.Permission{}.List(pagination.PageIndex, pagination.PageCount)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	pagination.PageTotal = pageTotal
	this.Data["permissions"] = permissions
	this.Data["pages"] = pagination
	this.AddBreadcrumbs("权限管理", this.URLFor("PermissionController.Index"))
	this.ShowHtml()
}

//permision add page
func (this *PermissionController) Add() {
	permissionModel := models.Permission{}

	if this.isPost() {
		permissionModel.Name = strings.TrimSpace(this.GetString("permission_name"))
		permissionModel.Description = strings.TrimSpace(this.GetString("permission_desc"))

		if isInsert, _ := permissionModel.Insert(); isInsert {
			this.redirectMessage(this.URLFor("PermissionController.Index"), "添加成功", MESSAGE_TYPE_SUCCESS)
		} else {
			this.redirectMessage(this.URLFor("PermissionController.Add"), "添加失败", MESSAGE_TYPE_ERROR)
		}

	}

	this.Data["model"] = permissionModel
	this.AddBreadcrumbs("权限管理", this.URLFor("PermissionController.Index"))
	this.AddBreadcrumbs("新增", this.URLFor("PermissionController.Add"))
	this.ShowHtml()
}

//permision update page
func (this *PermissionController) Put() {
	permissionModel := models.Permission{}
	if permission_id, err := this.GetInt("permission_id"); err != nil {
		this.redirectMessage(this.URLFor("PermissionController.Index"), "数据不存在", MESSAGE_TYPE_ERROR)
	} else if err := permissionModel.FindById(permission_id); err != nil || permissionModel.IsNewRecord() {
		this.redirectMessage(this.URLFor("PermissionController.Index"), "数据获取失败", MESSAGE_TYPE_ERROR)
	}

	if this.isPost() {
		permissionModel.Name = strings.TrimSpace(this.GetString("permission_name"))
		permissionModel.Description = strings.TrimSpace(this.GetString("permission_desc"))

		if err := permissionModel.Update(); err == nil {
			this.redirectMessage(this.URLFor("PermissionController.Index"), "修改成功", MESSAGE_TYPE_SUCCESS)
		} else {
			this.redirectMessage(this.URLFor("PermissionController.Put", "permission_id", this.GetString("permission_id")), "修改失败", MESSAGE_TYPE_ERROR)
		}

	}

	this.Data["model"] = permissionModel
	this.AddBreadcrumbs("权限管理", this.URLFor("PermissionController.Index"))
	this.AddBreadcrumbs("修改", this.URLFor("PermissionController.Put", "permission_id", this.GetString("permission_id")))
	this.AddBreadcrumbs(permissionModel.Name, "")
	this.ShowHtml()
}

//permision view page
func (this *PermissionController) Assignment() {
	permissionModel := models.Permission{}

	if id, err := this.GetInt("permission_id"); err == nil {
		if err := permissionModel.FindById(id); err != nil || permissionModel.IsNewRecord() {
			this.redirectMessage(this.URLFor("PermissionController.Index"), "数据获取失败", MESSAGE_TYPE_ERROR)
		}
	} else {
		this.redirectMessage(this.URLFor("PermissionController.Index"), "数据不存在", MESSAGE_TYPE_ERROR)
	}

	//获取已经授权的路由
	routeModel := models.Route{}
	routeAssignmentModel := models.AssignmentRoute{}
	allRouteAssignmentModels, err := routeAssignmentModel.FindAllByPermissionId(permissionModel.Id)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	var assignmentRoutes []string
	if allRouteAssignmentModels != nil {
		var allPermissionAssignmentIds []int

		for _, pr := range allRouteAssignmentModels {
			allPermissionAssignmentIds = append(allPermissionAssignmentIds, pr.RouteId)
		}

		prs, _ := routeModel.FindAllByIds(allPermissionAssignmentIds)

		if prs != nil {
			for _, r := range prs {
				assignmentRoutes = append(assignmentRoutes, r.Url)
			}
		}
	}

	//获取全部可以授权的路由
	routes, _ := routeModel.FindAll()
	var allRoutes []string

	if routes != nil {
		for _, route := range routes {
			if !utils.InSlice(route.Url, assignmentRoutes) {
				allRoutes = append(allRoutes, route.Url)
			}
		}
	}

	this.Data["model"] = permissionModel
	this.Data["allRoutes"] = allRoutes
	this.Data["assignmentRoutes"] = assignmentRoutes
	this.AddBreadcrumbs("权限管理", this.URLFor("PermissionController.Index"))
	this.AddBreadcrumbs("授权", this.URLFor("PermissionController.Assignment", "permission_id", this.GetString("permission_id")))
	this.AddBreadcrumbs(permissionModel.Name, "")
	this.ShowHtml()
}

//route permission ajax add page
func (this *PermissionController) AjaxAddRoute() {
	data := JsonData{}

	if this.isPost() {
		route := strings.TrimSpace(this.GetString("route"))

		routeModel := models.Route{}
		routeAssignmentModel := models.AssignmentRoute{}

		if route != "" {
			routeModel.FindByUrl(route)

			if routeModel.Id > 0 {
				routeAssignmentModel.RouteId = routeModel.Id
			} else {
				data.Code = 400
				data.Message = "路由不存在"
				this.ShowJSON(&data)
			}
		} else {
			data.Code = 400
			data.Message = "路由不能为空"
			this.ShowJSON(&data)
		}

		if i, err := this.GetInt("permissionId"); err == nil {
			routeAssignmentModel.PermissionId = i
		} else {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
		}

		if isInsert, err := routeAssignmentModel.Insert(); isInsert {
			data.Code = 200
			data.Message = "添加成功"
			data.Data = make(map[string]interface{})
			data.Data["route"] = routeModel.Url
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
func (this *PermissionController) AjaxRemoveRoute() {
	data := JsonData{}

	if this.isPost() {
		param_route := strings.TrimSpace(this.GetString("route"))
		routeAssignmentModel := models.AssignmentRoute{}

		var int_param_route_id int
		routeModel := models.Route{}

		if param_route != "" {
			routeModel.FindByUrl(param_route)

			if routeModel.Id > 0 {
				int_param_route_id = routeModel.Id
			} else {
				data.Code = 400
				data.Message = "路由不存在"
				this.ShowJSON(&data)
			}
		} else {
			data.Code = 400
			data.Message = "路由不能为空"
			this.ShowJSON(&data)
		}

		int_param_permission_id, err := this.GetInt("permissionId")

		if err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
		}

		if is_delete, err := routeAssignmentModel.Delete(int_param_route_id, int_param_permission_id); is_delete {
			data.Code = 200
			data.Message = "删除成功"
			data.Data = make(map[string]interface{})
			data.Data["route"] = routeModel.Url
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
		permision_id, err := this.GetInt("permission_id")

		if err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
		}

		permissionModel := models.Permission{}
		if err := permissionModel.FindById(permision_id); err != nil {
			data.Code = 400
			data.Message = "数据获取失败"
			this.ShowJSON(&data)
		}

		if is_delete, err := permissionModel.Delete(); is_delete {
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
