package controllers

import (
	. "grabc/views"
)

type PermissionController struct {
	BaseController
}

func (this *PermissionController) Index() {
	routeIndex := &RouteIndex{}
	this.ShowHtml(routeIndex)
}
