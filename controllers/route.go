package controllers

import (
	"github.com/astaxie/beego/utils"
	"github.com/codyi/grabc/libs"
	"github.com/codyi/grabc/models"
	"github.com/codyi/grabc/views/route"
	"strings"
)

type RouteController struct {
	BaseController
}

//路由首页
func (this *RouteController) Index() {
	routeModel := models.Route{}
	routes, _ := routeModel.FindAll()
	var addRoutes []string

	if routes != nil {
		for _, route := range routes {
			addRoutes = append(addRoutes, route.Url)
		}
	}

	allRoutes := libs.AllRoutes()
	var notAddRoutes []string

	for _, route := range allRoutes {
		if !utils.InSlice(route, addRoutes) {
			notAddRoutes = append(notAddRoutes, route)
		}
	}

	this.htmlData["notAddRoutes"] = notAddRoutes
	this.htmlData["addRoutes"] = addRoutes
	this.AddBreadcrumbs("路由管理", this.URLFor("RouteController.Index"))
	this.ShowHtml(&route.Index{})
}

//route ajax add page
func (this *RouteController) AjaxAdd() {
	data := JsonData{}

	if this.isPost() {
		route := strings.TrimSpace(this.GetString("route"))

		routeModel := models.Route{}
		routeModel.Url = route

		if isInsert, _ := routeModel.Insert(); isInsert {
			data.Code = 200
			data.Message = "添加成功"
			data.Data = make(map[string]interface{})
			data.Data["route"] = route
		} else {
			data.Code = 400
			data.Message = "添加失败"
		}

	} else {
		data.Code = 400
		data.Message = "非法请求"
	}

	this.ShowJSON(&data)
}

//route ajax remove page
func (this *RouteController) AjaxRemove() {
	data := JsonData{}

	if this.isPost() {
		route := strings.TrimSpace(this.GetString("route"))

		routeModel := models.Route{}

		if isDelete, _ := routeModel.DeleteByRoute(route); isDelete {
			data.Code = 200
			data.Message = "删除成功"
		} else {
			data.Code = 400
			data.Message = "删除成功"
		}

	} else {
		data.Code = 400
		data.Message = "非法请求"
	}

	this.ShowJSON(&data)
}
