package controller

import (
	. "grabc/views"
)

type PermissionController struct {
	BaseController
}

func (this *PermissionController) Index() {
	routeIndex := &RouteIndex{}
	this.ServerHtml(routeIndex)
}
