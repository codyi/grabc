package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/codyi/grabc/libs"
	"path/filepath"
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

	if !libs.CheckAccess(this.controllerName, this.actionName, libs.AccessRoutes()) {
		this.redirect(libs.Http_403)
	}

	this.funcMap = template.FuncMap{
		"pagination":     libs.PaginationRender,
		"unixTimeFormat": libs.UnixTimeFormat,
	}
}

//显示页面
func (this *BaseController) ShowHtml(tpl ...string) {
	var tplName string

	if len(tpl) > 0 {
		tplName = libs.Template.ViewPath + tpl[0]
	} else {
		tplName = libs.Template.ViewPath + this.controllerName + "/" + this.actionName + ".html"
	}

	var viewContent bytes.Buffer
	this.htmlData["homeUrl"] = this.homeUrl
	this.htmlData["viewpaht"] = libs.Template.ViewPath
	this.htmlData["alert_messages"] = this.Alert
	this.htmlData["global_css"] = libs.Template.GlobalCss()
	this.htmlData["global_js"] = libs.Template.GlobalJs()
	this.htmlData["alert"] = this.ShowAlert()

	if len(libs.Template.Data) > 0 {
		for k, v := range libs.Template.Data {
			this.htmlData[k] = v
		}
	}

	//渲染视图模板
	tmpl, err := template.New(filepath.Base(tplName)).Funcs(this.funcMap).ParseFiles(tplName)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		this.StopRun()
	}

	tmpl.Execute(&viewContent, this.htmlData)

	//渲染layout模板
	tmpl, err = template.New(filepath.Base(libs.Template.Layout)).ParseFiles(libs.Template.Layout)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		this.StopRun()
	}

	libs.Template.Data["breadcrumbs"] = this.ShowBreadcrumbs()
	libs.Template.Data["LayoutContent"] = viewContent.String()
	libs.Template.Data["grabc_menus"] = libs.AccessMenus()
	var htmlContent bytes.Buffer
	tmpl.Execute(&htmlContent, libs.Template.Data)

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
