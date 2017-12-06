package controllers

import (
	"grabc/libs"
	"grabc/models"
	"grabc/views/role"
	"strconv"
	"strings"
)

type RoleController struct {
	BaseController
}

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

	roles, pageTotal, err := models.Role{}.FindAll(pagination.PageIndex, pagination.PageCount)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	pagination.PageTotal = pageTotal
	this.htmlData["roles"] = roles
	this.htmlData["pages"] = pagination
	this.AddBreadcrumbs("角色管理", this.URLFor("RoleController.Index"))
	this.ShowHtml(&role.Index{})
}

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

func (this *RoleController) Get() {
	this.ShowHtml(&role.Get{})
}

func (this *RoleController) Delete() {
	this.ShowHtml(&role.Delete{})
}
