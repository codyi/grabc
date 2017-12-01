package controllers

import (
	"grabc/views/route"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) Index() {
	this.ShowHtml(&route.Index{})
}
