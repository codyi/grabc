package controllers

import (
	"github.com/astaxie/beego/utils"
	"github.com/codyi/grabc/libs"
	"github.com/codyi/grabc/models"
	"strings"
)

type AssignmentController struct {
	BaseController
}

//用户授权列表
func (this *AssignmentController) Index() {
	page_index, err := this.GetInt("page_index")

	//分页设置
	pagination := libs.Pagination{}
	pagination.PageCount = 20
	pagination.Url = this.URLFor("AssignmentController.Index")

	if err == nil {
		pagination.PageIndex = page_index
	} else {
		pagination.PageIndex = 1
	}

	userList, pageTotal, err := (*models.UserModel).UserList(pagination.PageIndex, pagination.PageCount)

	if err != nil {
		this.AddErrorMessage(err.Error())
	}

	//查找用户对应的角色，并获取角色名称
	var userIds []int
	for id, _ := range userList {
		userIds = append(userIds, id)
	}

	type userData struct {
		Id        int
		Name      string
		RoleNames []string
	}

	var userItems []userData
	roleAssignmentModel := models.AssignmentRole{}
	roleModel := models.Role{}

	for id, name := range userList {
		u := userData{}
		u.Id = id
		u.Name = name

		roleIds, err := roleAssignmentModel.FindRoleIdsByUserId(u.Id)
		if err == nil && len(roleIds) > 0 {
			u.RoleNames, _ = roleModel.ListNamesByIds(roleIds)
		}

		userItems = append(userItems, u)
	}

	pagination.PageTotal = pageTotal
	this.Data["userItems"] = userItems
	this.Data["pages"] = pagination
	this.AddBreadcrumbs("授权管理", this.URLFor("AssignmentController.Index"))
	this.ShowHtml()
}

//用户授权
func (this *AssignmentController) User() {
	user_id, err := this.GetInt("user_id")

	if err != nil {
		this.redirectMessage(this.URLFor("AssignmentController.Index"), "用户ID不正确", MESSAGE_TYPE_ERROR)
	}

	user_name := (*models.UserModel).FindNameById(user_id)

	if user_name == "" {
		this.redirectMessage(this.URLFor("AssignmentController.Index"), "用户没有找到", MESSAGE_TYPE_ERROR)
	}

	//获取全部的权限
	allRoles := models.Role{}.FindAll()
	//获取全部已经授权的角色ID
	allAssignmentRoleIds, err := models.AssignmentRole{}.FindRoleIdsByUserId(user_id)

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

	this.Data["name"] = user_name
	this.Data["user_id"] = user_id
	this.Data["assignmentRoles"] = assignmentRoles
	this.Data["unassignmentRoles"] = unassignmentRoles
	this.AddBreadcrumbs("用户授权", this.URLFor("AssignmentController.Index"))
	this.ShowHtml()
}

func (this *AssignmentController) AjaxAdd() {
	data := JsonData{}

	if this.isPost() {
		param_role := strings.TrimSpace(this.GetString("role"))
		int_param_user_id, user_id_err := this.GetInt("user_id")

		roleModel := models.Role{}

		if param_role != "" {
			roleModel.FindByName(param_role)

			if roleModel.Id <= 0 {
				data.Code = 400
				data.Message = "角色不存在"
				this.ShowJSON(&data)
			}
		} else {
			data.Code = 400
			data.Message = "角色不能为空"
			this.ShowJSON(&data)
		}

		if user_id_err != nil {
			data.Code = 400
			data.Message = user_id_err.Error()
			this.ShowJSON(&data)
		}

		roleAssignmentModel := models.AssignmentRole{}
		roleAssignmentModel.UserId = int_param_user_id
		roleAssignmentModel.RoleId = roleModel.Id
		if isInsert, err := roleAssignmentModel.Insert(); isInsert {
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

func (this *AssignmentController) AjaxRemove() {
	data := JsonData{}

	if this.isPost() {
		param_role := strings.TrimSpace(this.GetString("role"))
		int_param_user_id, user_id_err := this.GetInt("user_id")

		roleModel := models.Role{}

		if param_role != "" {
			roleModel.FindByName(param_role)

			if roleModel.Id <= 0 {
				data.Code = 400
				data.Message = "角色不存在"
				this.ShowJSON(&data)
			}
		} else {
			data.Code = 400
			data.Message = "角色不能为空"
			this.ShowJSON(&data)
		}

		if user_id_err != nil {
			data.Code = 400
			data.Message = user_id_err.Error()
			this.ShowJSON(&data)
		}

		roleAssignmentModel := models.AssignmentRole{}
		if err := roleAssignmentModel.FindByRoleIdAndUserId(roleModel.Id, int_param_user_id); err == nil {
			if is_delete, err := roleAssignmentModel.Delete(); is_delete {
				data.Code = 200
				data.Message = "取消授权成功"
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
