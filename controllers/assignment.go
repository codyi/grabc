package controllers

import (
	"github.com/astaxie/beego/utils"
	"grabc/libs"
	"grabc/models"
	"grabc/views/assignment"
	"strconv"
	"strings"
)

type AssignmentController struct {
	BaseController
}

//用户授权列表
func (this *AssignmentController) Index() {
	page_index := strings.TrimSpace(this.GetString("page_index"))

	//分页设置
	pagination := libs.Pagination{}
	pagination.PageCount = 20
	pagination.Url = this.URLFor("AssignmentController.Index")

	if s, err := strconv.Atoi(page_index); err == nil {
		pagination.PageIndex = s
	} else {
		pagination.PageIndex = 1
	}

	userList, pageTotal, err := (*models.UserModel).UserList(pagination.PageIndex, pagination.PageCount)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	pagination.PageTotal = pageTotal
	this.htmlData["userList"] = userList
	this.htmlData["pages"] = pagination
	this.AddBreadcrumbs("授权管理", this.URLFor("AssignmentController.Index"))
	this.ShowHtml(&assignment.Index{})
}

//用户授权
func (this *AssignmentController) User() {
	param_user_id := strings.TrimSpace(this.GetString("user_id"))
	var user_id int
	if s, err := strconv.Atoi(param_user_id); err == nil {
		user_id = s
	} else {
		this.AddErrorMessage("用户ID不存在")
	}

	user_name := (*models.UserModel).FindNameById(user_id)

	if user_name == "" {
		this.AddErrorMessage("用户没有找到")
	}

	//获取全部的权限
	allRoles := models.Role{}.FindAll()
	//获取全部已经授权的角色ID
	allAssignmentRoleIds, err := models.RoleAssignment{}.FindRoleIdsByUserId(user_id)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	var assignmentRoles []string   //已经授权的角色
	var unassignmentRoles []string //未授权的角色

	//分离用户授权和未授权的角色
	for _, role := range allRoles {
		for _, roleId := range allAssignmentRoleIds {
			if role.Id == roleId {
				assignmentRoles = append(assignmentRoles, role.Name)
			}
		}

		if !utils.InSlice(role.Name, assignmentRoles) {
			unassignmentRoles = append(unassignmentRoles, role.Name)
		}
	}

	this.htmlData["name"] = user_name
	this.htmlData["assignmentRoles"] = assignmentRoles
	this.htmlData["unassignmentRoles"] = unassignmentRoles
	this.AddBreadcrumbs("用户授权", this.URLFor("AssignmentController.Index"))
	this.ShowHtml(&assignment.User{})
}
