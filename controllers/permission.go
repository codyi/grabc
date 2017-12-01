package controllers

import (
	"grabc/models"
	"grabc/views/permission"
	"strings"
)

type PermissionController struct {
	BaseController
}

func (this *PermissionController) Index() {
	var pageIndex, pageCount int
	pageIndex = 1
	pageCount = 10

	permissions, err := models.Permission{}.FindAll(pageIndex, pageCount)

	if err != nil {
		this.AddErrorMessage("信息获取失败")
	}
	// this.AddErrorMessage("信息获取失败")
	this.htmlData["permissions"] = permissions
	this.ShowHtml(&permission.Index{})
}

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
