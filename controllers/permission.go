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
	pagination.PageCount = 3
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
	this.ShowHtml(&permission.Index{})
}

//permision add page
func (this *PermissionController) Add() {

	if this.isPost() {
		permission_name := strings.TrimSpace(this.GetString("permission_name"))
		permission_desc := strings.TrimSpace(this.GetString("permission_desc"))

		permission := models.Permission{}
		permission.Name = permission_name
		permission.Description = permission_desc

		if isInsert, _ := permission.Insert(); isInsert {
			this.AddSuccessMessage("添加成功")
		} else {
			this.AddErrorMessage("添加失败")
		}

	}

	this.ShowHtml(&permission.Add{})
}
