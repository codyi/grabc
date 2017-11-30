package controllers

import (
	. "grabc/views"
)

type AssignmentController struct {
	BaseController
}

func (this *AssignmentController) Index() {
	routeIndex := &RouteIndex{}
	this.ShowHtml(routeIndex)
}
