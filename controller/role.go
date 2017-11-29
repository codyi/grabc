package controller

import (
	. "grabc/views"
)

type RoleController struct {
	BaseController
}

func (this *RoleController) Index() {
	routeIndex := &RouteIndex{}
	this.ServerHtml(routeIndex)
}
