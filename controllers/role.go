package controllers

import (
	"github.com/codyi/grabc/libs"
	"github.com/codyi/grabc/models"
	"strings"
)

type RoleController struct {
	BaseController
}

//角色管理首页
func (this *RoleController) Index() {
	page_index, err := this.GetInt("page_index")

	//分页设置
	pagination := libs.Pagination{}
	pagination.PageCount = 20
	pagination.Url = this.URLFor("RoleController.Index")

	if err == nil {
		pagination.PageIndex = page_index
	} else {
		pagination.PageIndex = 1
	}

	roles, pageTotal, err := models.Role{}.List(pagination.PageIndex, pagination.PageCount)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	pagination.PageTotal = pageTotal
	this.Data["roles"] = roles
	this.Data["pages"] = pagination
	this.AddBreadcrumbs("角色管理", this.URLFor("RoleController.Index"))
	this.ShowHtml()
}

//添加角色页面
func (this *RoleController) Post() {
	roleModel := models.Role{}

	if this.isPost() {
		roleModel.Name = strings.TrimSpace(this.GetString("role_name"))
		roleModel.Description = strings.TrimSpace(this.GetString("role_desc"))

		if isInsert, _ := roleModel.Insert(); isInsert {
			this.redirectMessage(this.URLFor("RoleController.Index"), "添加成功", MESSAGE_TYPE_SUCCESS)
		} else {
			this.redirectMessage(this.URLFor("RoleController.Index"), "添加失败", MESSAGE_TYPE_ERROR)
		}

	}

	this.Data["model"] = roleModel
	this.AddBreadcrumbs("角色管理", this.URLFor("RoleController.Index"))
	this.AddBreadcrumbs("新增", this.URLFor("RoleController.Add"))
	this.ShowHtml()
}

//角色修改
func (this *RoleController) Put() {
	roleModel := models.Role{}

	if id, err := this.GetInt("role_id"); err == nil {
		if err := roleModel.FindById(id); err != nil || roleModel.IsNewRecord() {
			this.redirectMessage(this.URLFor("RoleController.Index"), "数据获取失败", MESSAGE_TYPE_ERROR)
		}
	} else {
		this.redirectMessage(this.URLFor("RoleController.Index"), "数据不存在", MESSAGE_TYPE_ERROR)
	}

	if this.isPost() {
		roleModel.Name = strings.TrimSpace(this.GetString("role_name"))
		roleModel.Description = strings.TrimSpace(this.GetString("role_desc"))

		if err := roleModel.Update(); err == nil {
			this.redirectMessage(this.URLFor("RoleController.Index"), "修改成功", MESSAGE_TYPE_SUCCESS)
		} else {
			this.redirectMessage(this.URLFor("RoleController.Put", "role_id", this.GetString("role_id")), "修改失败", MESSAGE_TYPE_ERROR)
		}

	}

	this.Data["model"] = roleModel
	this.AddBreadcrumbs("角色管理", this.URLFor("RoleController.Index"))
	this.AddBreadcrumbs("修改", this.URLFor("RoleController.Put", "role_id", this.GetString("role_id")))
	this.AddBreadcrumbs(roleModel.Name, "")
	this.ShowHtml()
}

//角色授权展示页面
func (this *RoleController) Assignment() {
	roleModel := models.Role{}

	if id, err := this.GetInt("role_id"); err == nil {
		if err := roleModel.FindById(id); err != nil || roleModel.IsNewRecord() {
			this.redirectMessage(this.URLFor("RoleController.Index"), "数据获取失败", MESSAGE_TYPE_ERROR)
		}
	} else {
		this.redirectMessage(this.URLFor("RoleController.Index"), "数据不存在", MESSAGE_TYPE_ERROR)
	}

	//获取全部权限
	allPermissions, err := models.Permission{}.FindAll()
	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	//获取已经授权的权限关系
	allPermissionAssignments, err := models.AssignmentPermission{}.FindAllByRoleId(roleModel.Id)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	var assignmentPermssionNames, unassignmentPermssionNames []string
	for _, p := range allPermissions {
		isFound := false
		for _, pa := range allPermissionAssignments {
			if p.Id == pa.PermissionId {
				isFound = true
				assignmentPermssionNames = append(assignmentPermssionNames, p.Name)
				break
			}
		}

		if !isFound {
			unassignmentPermssionNames = append(unassignmentPermssionNames, p.Name)
		}
	}

	this.Data["model"] = roleModel
	this.Data["unassignmentPermssionNames"] = unassignmentPermssionNames
	this.Data["assignmentPermssionNames"] = assignmentPermssionNames
	this.AddBreadcrumbs("角色管理", this.URLFor("RoleController.Index"))
	this.AddBreadcrumbs("权限授权", this.URLFor("RoleController.Assignment", "role_id", this.GetString("role_id")))
	this.AddBreadcrumbs(roleModel.Name, "")
	this.ShowHtml()
}

//ajax-权限授权给角色
func (this *RoleController) AjaxAssignment() {
	data := JsonData{}

	if this.isPost() {
		paramPermissionName := strings.TrimSpace(this.GetString("permission_name"))
		intRoleId, err := this.GetInt("role_id")

		if err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
		}

		permissionModel := models.Permission{}
		if err := permissionModel.FindByName(paramPermissionName); err != nil || permissionModel.IsNewRecord() {
			data.Code = 400
			data.Message = "权限不能为空"
			this.ShowJSON(&data)
		}

		perAssignmentModel := models.AssignmentPermission{}
		perAssignmentModel.RoleId = intRoleId
		perAssignmentModel.PermissionId = permissionModel.Id
		if isInsert, err := perAssignmentModel.Insert(); isInsert {
			data.Code = 200
			data.Message = "授权成功"
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

//ajax-角色取消权限
func (this *RoleController) AjaxUnassignment() {
	data := JsonData{}

	if this.isPost() {
		paramPermissionName := strings.TrimSpace(this.GetString("permission_name"))
		intRoleId, err := this.GetInt("role_id")

		if err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
		}

		permissionModel := models.Permission{}
		if err := permissionModel.FindByName(paramPermissionName); err != nil || permissionModel.IsNewRecord() {
			data.Code = 400
			data.Message = "权限不能为空"
			this.ShowJSON(&data)
		}

		perAssignmentModel := models.AssignmentPermission{}
		if is_delete, err := perAssignmentModel.Delete(intRoleId, permissionModel.Id); err == nil {
			if is_delete {
				data.Code = 200
				data.Message = "删除成功"
			} else {
				data.Code = 400
				data.Message = err.Error()
			}
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

//role delete page
func (this *RoleController) Delete() {
	data := JsonData{}
	if this.isPost() {
		role_id, err := this.GetInt("role_id")

		if err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
		}

		roleModel := models.Role{}
		if err := roleModel.FindById(role_id); err != nil {
			data.Code = 400
			data.Message = "数据获取失败"
			this.ShowJSON(&data)
		}

		if is_delete, err := roleModel.Delete(); is_delete {
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
