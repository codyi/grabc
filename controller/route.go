package controller

import (
	. "grabc/views"
)

type RouteController struct {
	BaseController
}

func (this *RouteController) Index() {
	routeIndex := &RouteIndex{}
	this.htmlData["aa"] = "bb"
	this.ServerHtml(routeIndex)
}
