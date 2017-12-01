package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"grabc/libs"
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
	controllerName string
	actionName     string
	htmlData       map[interface{}]interface{}
}

// redirect to url
func (this *BaseController) redirect(url string) {
	this.Controller.Redirect(url, 301)
	this.StopRun()
}

// Prepare
func (this *BaseController) Prepare() {
	this.htmlData = make(map[interface{}]interface{})
	controlerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controlerName[0 : len(controlerName)-10])
	this.actionName = strings.ToLower(actionName)
}

//显示页面
func (this *BaseController) ShowHtml(html HtmlTemplate) {
	tmpl, err := template.New("grabc").Parse(html.Html())

	if err != nil {
		this.Ctx.WriteString(err.Error())
	}

	var htmlContent bytes.Buffer
	this.htmlData["alert_messages"] = this.Alert
	tmpl.Execute(&htmlContent, this.htmlData)
	this.Ctx.WriteString(htmlContent.String())
	this.StopRun()
}

//server Json
func (this *BaseController) ShowJSON(data *JsonData) {
	this.Data["json"] = data
	this.Controller.ServeJSON()
}

// 是否POST提交
func (this *BaseController) isPost() bool {
	return this.Ctx.Request.Method == "POST"
}
