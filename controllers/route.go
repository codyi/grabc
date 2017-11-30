package controllers

import (
	"grabc/models"
	. "grabc/views"
	"strings"
)

type RouteController struct {
	BaseController
}

//route index page
func (this *RouteController) Index() {
	routeIndex := &RouteIndex{}
	this.ShowHtml(routeIndex)
}

//route ajax add page
func (this *RouteController) Add() {
	data := JsonData{}

	if this.isPost() {
		route := strings.TrimSpace(this.GetString("route"))

		routeModel := models.Route{}
		routeModel.Route = route

		if isInsert, err := routeModel.Insert(); isInsert {
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
		data.Message = "添加失败"
	}

	this.ShowJSON(&data)
}
