package controllers

import (
	"bytes"
	"github.com/astaxie/beego"
	"github.com/codyi/grabc/libs"
	"html/template"
	"strings"
)

const MESSAGE_TYPE_SUCCESS = "success"
const MESSAGE_TYPE_ERROR = "error"

func init() {
	libs.ExceptMethodAppend("ShowHtml", "ShowJSON")
	beego.AddFuncMap("unixTimeFormat", libs.UnixTimeFormat)
	beego.AddFuncMap("pagination", libs.PaginationRender)
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
	funcMap        template.FuncMap
	homeUrl        string
}

// redirect to url
func (this *BaseController) redirect(url string) {
	this.Controller.Redirect(url, 302)
	this.StopRun()
}

func (this *BaseController) redirectMessage(url, message, messageType string) {
	this.AddBreadcrumbs("消息提示", "")
	this.Data["redirect_url"] = url
	this.Data["message"] = message
	this.Data["message_type"] = messageType
	this.ShowHtml("layout/tip.html")
}

// Prepare
func (this *BaseController) Prepare() {
	controlerName, actionName := this.GetControllerAndAction()
	this.controllerName = strings.ToLower(controlerName[0 : len(controlerName)-10])
	this.actionName = strings.ToLower(actionName)
	libs.BeegoC = &this.Controller

	if !libs.CheckAccess(this.controllerName, this.actionName, libs.AccessRoutes()) {
		this.redirectMessage("/", "权限不足", MESSAGE_TYPE_ERROR)
	}
}

// RenderBytes returns the bytes of rendered template string. Do not send out response.
func (c *BaseController) RenderBytes() ([]byte, error) {
	buf, err := c.renderTemplate()
	//if the controller has set layout, then first get the tplName's content set the content to the layout
	if err == nil && libs.Template.LayoutName != "" {
		c.Data["LayoutContent"] = template.HTML(buf.String())

		if c.LayoutSections != nil {
			for sectionName, sectionTpl := range c.LayoutSections {
				if sectionTpl == "" {
					c.Data[sectionName] = ""
					continue
				}
				buf.Reset()
				err = beego.ExecuteViewPathTemplate(&buf, sectionTpl, libs.Template.LayoutPath, c.Data)
				if err != nil {
					return nil, err
				}
				c.Data[sectionName] = template.HTML(buf.String())
			}
		}

		buf.Reset()
		beego.ExecuteViewPathTemplate(&buf, libs.Template.LayoutName, libs.Template.LayoutPath, c.Data)
	}
	return buf.Bytes(), err
}

// Render sends the response with rendered template bytes as text/html type.
func (c *BaseController) Render() error {
	rb, err := c.RenderBytes()
	if err != nil {
		return err
	}

	if c.Ctx.ResponseWriter.Header().Get("Content-Type") == "" {
		c.Ctx.Output.Header("Content-Type", "text/html; charset=utf-8")
	}

	return c.Ctx.Output.Body(rb)
}

func (c *BaseController) renderTemplate() (bytes.Buffer, error) {
	var buf bytes.Buffer
	if c.TplName == "" {
		c.TplName = strings.ToLower(c.controllerName) + "/" + strings.ToLower(c.actionName) + "." + c.TplExt
	}
	if c.TplPrefix != "" {
		c.TplName = c.TplPrefix + c.TplName
	}
	return buf, beego.ExecuteViewPathTemplate(&buf, c.TplName, libs.Template.ViewPath, c.Data)
}

//重新定义beego的render
func (this *BaseController) ShowHtml(tpl ...string) {
	if len(tpl) > 0 {
		this.TplName = tpl[0]
	} else {
		this.TplName = this.controllerName + "/" + this.actionName + ".html"
	}

	this.Data["homeUrl"] = this.homeUrl
	this.Data["viewpaht"] = libs.Template.ViewPath
	this.Data["alert_messages"] = this.Alert
	this.Data["global_css"] = libs.Template.GlobalCss()
	this.Data["global_js"] = libs.Template.GlobalJs()
	this.Data["alert"] = this.ShowAlert()
	this.Data["breadcrumbs"] = this.ShowBreadcrumbs()
	this.Data["menus"] = libs.ShowMenu(this.controllerName, this.actionName)

	layoutData := this.GetSession("grabc_layout_data")

	if layoutData != nil {
		d := layoutData.(map[string]interface{})
		for name, value := range d {
			if _, isExist := this.Data[name]; isExist {
				panic("设置layout数据失败，因为" + name + "已经存在")
			}

			this.Data[name] = value
		}
	}

	this.Render()
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
