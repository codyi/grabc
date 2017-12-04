package controllers

import (
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
func (this *PermissionController) Get() {
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

	this.htmlData["model"] = permissionModel
	this.AddBreadcrumbs("权限管理", this.URLFor("PermissionController.Index"))
	this.AddBreadcrumbs("查看", this.URLFor("PermissionController.Get", "permission_id", permission_id))
	this.AddBreadcrumbs(permissionModel.Name, "")
	this.ShowHtml(&permission.Get{})
}

//permision delete page
func (this *PermissionController) Delete() {
}

//permision assignment page
func (this *PermissionController) Assignment() {
}
