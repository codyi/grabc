package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/codyi/grabc/libs"
	"strings"
	"text/template"
)

func init() {
	libs.ExceptMethodAppend("ShowHtml", "ShowJSON")
}

type HtmlTemplate interface {
	Html() string
}

type JsonData struct {
	Code    int
	Message string
	Data    map[string]interface{}
}

//Base controller
type BaseController struct {
	beego.Controller
	libs.Alert
	libs.Breadcrumbs
	controllerName string
	actionName     string
	htmlData       map[interface{}]interface{}
	funcMap        template.FuncMap
	homeUrl        string
}

// redirect to url
func (this *BaseController) redirect(url string) {
	this.Controller.Redirect(url, 302)
	this.StopRun()
}

// Prepare
func (this *BaseController) Prepare() {
	this.htmlData = make(map[interface{}]interface{})
	controlerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controlerName[0 : len(controlerName)-10])
	this.actionName = strings.ToLower(actionName)
	this.funcMap = template.FuncMap{
		"pagination":     libs.PaginationRender,
		"unixTimeFormat": libs.UnixTimeFormat,
	}

	if !libs.CheckAccess(this.controllerName, this.actionName, libs.AccessRoutes()) {
		this.redirect(libs.Http_403)
	}
}

//显示页面
func (this *BaseController) ShowHtml(html HtmlTemplate) {
	tmpl, err := template.New("grabc").Funcs(this.funcMap).Parse(html.Html())

	if err != nil {
		this.Ctx.WriteString(err.Error())
	}

	var htmlContent bytes.Buffer
	this.htmlData["homeUrl"] = this.homeUrl
	this.htmlData["alert_messages"] = this.Alert
	this.htmlData["breadcrumbs"] = this.Breadcrumbs.Items
	tmpl.Execute(&htmlContent, this.htmlData)
	this.Ctx.WriteString(htmlContent.String())
	this.StopRun()
}

//server Json
func (this *BaseController) ShowJSON(data *JsonData) {
	this.Data["json"] = data
	this.Controller.ServeJSON()
	this.StopRun()
}

// 是否POST提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}
