package controller

import (
	"bytes"
	"github.com/astaxie/beego"
	"strings"
	"text/template"
)

type HtmlTemplate interface {
	Html() string
}

type BaseController struct {
	beego.Controller
	controllerName string
	actionName     string
	htmlData       map[interface{}]interface{}
}

// Prepare
func (this *BaseController) Prepare() {
	this.htmlData = make(map[interface{}]interface{})
	controlerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controlerName[0 : len(controlerName)-10])
	this.actionName = strings.ToLower(actionName)
}

//显示页面
func (this *BaseController) ServerHtml(html HtmlTemplate) {
	tmpl, err := template.New("grabc").Parse(html.Html())

	if err != nil {
		this.Ctx.WriteString(err.Error())
	}

	var htmlContent bytes.Buffer
	tmpl.Execute(&htmlContent, this.htmlData)
	this.Ctx.WriteString(htmlContent.String())
	this.StopRun()
}
