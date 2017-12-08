package controllers

import (
	"github.com/codyi/grabc/libs"
	"github.com/codyi/grabc/models"
	"github.com/codyi/grabc/views/role"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

//角色管理首页
func (this *RoleController) Index() {
	page_index := strings.TrimSpace(this.GetString("page_index"))

	//分页设置
	pagination := libs.Pagination{}
	pagination.PageCount = 20
	pagination.Url = this.URLFor("RoleController.Index")

	if s, err := strconv.Atoi(page_index); err == nil {
		pagination.PageIndex = s
	} else {
		pagination.PageIndex = 1
	}

	roles, pageTotal, err := models.Role{}.List(pagination.PageIndex, pagination.PageCount)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	pagination.PageTotal = pageTotal
	this.htmlData["roles"] = roles
	this.htmlData["pages"] = pagination
	this.AddBreadcrumbs("角色管理", this.URLFor("RoleController.Index"))
	this.ShowHtml(&role.Index{})
}

//添加角色页面
func (this *RoleController) Post() {
	roleModel := models.Role{}

	if this.isPost() {
		role_name := strings.TrimSpace(this.GetString("role_name"))
		role_desc := strings.TrimSpace(this.GetString("role_desc"))

		roleModel.Name = role_name
		roleModel.Description = role_desc

		if isInsert, _ := roleModel.Insert(); isInsert {
			this.AddSuccessMessage("添加成功")
		} else {
			this.AddErrorMessage("添加失败")
		}

	}

	this.htmlData["model"] = roleModel
	this.AddBreadcrumbs("角色管理", this.URLFor("RoleController.Index"))
	this.AddBreadcrumbs("新增", this.URLFor("RoleController.Add"))
	this.ShowHtml(&role.Post{})
}

//角色修改
func (this *RoleController) Put() {
	roleModel := models.Role{}
	role_id := strings.TrimSpace(this.GetString("role_id"))

	if id, err := strconv.Atoi(role_id); err == nil {
		if err := roleModel.FindById(id); err != nil {
			this.AddErrorMessage("数据获取失败")
		}
	} else {
		this.AddErrorMessage("数据不存在")
	}

	if this.isPost() && !roleModel.IsNewRecord() {
		role_name := strings.TrimSpace(this.GetString("role_name"))
		role_desc := strings.TrimSpace(this.GetString("role_desc"))

		roleModel.Name = role_name
		roleModel.Description = role_desc

		if err := roleModel.Update(); err == nil {
			this.AddSuccessMessage("修改成功")
		} else {
			this.AddErrorMessage("修改失败")
		}

	}

	this.htmlData["model"] = roleModel
	this.AddBreadcrumbs("角色管理", this.URLFor("RoleController.Index"))
	this.AddBreadcrumbs("修改", this.URLFor("RoleController.Put", "role_id", role_id))
	this.AddBreadcrumbs(roleModel.Name, "")
	this.ShowHtml(&role.Put{})
}

//角色授权展示页面
func (this *RoleController) Assignment() {
	roleModel := models.Role{}
	role_id := strings.TrimSpace(this.GetString("role_id"))

	if id, err := strconv.Atoi(role_id); err == nil {
		if err := roleModel.FindById(id); err != nil {
			this.AddErrorMessage("数据获取失败")
		}
	} else {
		this.AddErrorMessage("数据不存在")
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

	this.htmlData["model"] = roleModel
	this.htmlData["unassignmentPermssionNames"] = unassignmentPermssionNames
	this.htmlData["assignmentPermssionNames"] = assignmentPermssionNames
	this.AddBreadcrumbs("角色管理", this.URLFor("RoleController.Index"))
	this.AddBreadcrumbs("权限授权", this.URLFor("RoleController.Assignment", "role_id", role_id))
	this.AddBreadcrumbs(roleModel.Name, "")
	this.ShowHtml(&role.Assignment{})
}

//ajax-权限授权给角色
func (this *RoleController) AjaxAssignment() {
	data := JsonData{}

	if this.isPost() {
		paramPermissionName := strings.TrimSpace(this.GetString("permission_name"))
		paramRoleId := strings.TrimSpace(this.GetString("role_id"))

		permissionModel := models.Permission{}
		if err := permissionModel.FindByName(paramPermissionName); err != nil {
			data.Code = 400
			data.Message = "权限不能为空"
			this.ShowJSON(&data)
			return
		}

		var intRoleId int
		if id, err := strconv.Atoi(paramRoleId); err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
			return
		} else {
			intRoleId = id
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
		paramRoleId := strings.TrimSpace(this.GetString("role_id"))

		permissionModel := models.Permission{}
		if err := permissionModel.FindByName(paramPermissionName); err != nil {
			data.Code = 400
			data.Message = "权限不能为空"
			this.ShowJSON(&data)
			return
		}

		var intRoleId int
		if id, err := strconv.Atoi(paramRoleId); err != nil {
			data.Code = 400
			data.Message = err.Error()
			this.ShowJSON(&data)
			return
		} else {
			intRoleId = id
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
