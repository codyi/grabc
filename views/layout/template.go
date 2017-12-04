package layout

var globalCss []string
var globalJs []string

//append css
func GloablCssAppend(css ...string) {
	for _, c := range css {
		globalCss = append(globalCss, c)
	}
}

//append js
func GloablJsAppend(js ...string) {
	for _, j := range js {
		globalJs = append(globalJs, j)
	}
}

func registerGloablCss() string {
	var cssHtml string
	if globalCss != nil {
		for _, c := range globalCss {
			cssHtml += "<link rel=\"stylesheet\" href=\"" + c + "\">"
		}
	}

	return cssHtml
}

func registerGloablJs() string {
	var jsHtml string
	if globalJs != nil {
		for _, j := range globalJs {
			jsHtml += "<script src=\"" + j + "\"></script>"
		}
	}

	return jsHtml
}

func init() {
	GloablCssAppend("https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css", "http://127.0.0.1:8080/static/html/global.css")
	GloablJsAppend("https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js", "https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js", "http://127.0.0.1:8080/static/html/global.js")
}

type BaseTemplate struct {
	selfCss []string
	selfJs  []string
}

//append css
func (this *BaseTemplate) SelfCssAppend(css ...string) {
	for _, c := range css {
		this.selfCss = append(this.selfCss, c)
	}
}

//append js
func (this *BaseTemplate) SelfJsAppend(js ...string) {
	for _, j := range js {
		this.selfJs = append(this.selfJs, j)
	}
}

func (this *BaseTemplate) registerSelfCss() string {
	var cssHtml string
	if this.selfCss != nil {
		for _, c := range this.selfCss {
			cssHtml += "<link rel=\"stylesheet\" href=\"" + c + "\">"
		}
	}

	return cssHtml
}

func (this *BaseTemplate) registerSelfJs() string {
	var jsHtml string
	if this.selfJs != nil {
		for _, j := range this.selfJs {
			jsHtml += "<script src=\"" + j + "\"></script>"
		}
	}

	return jsHtml
}

func (this *BaseTemplate) DealHtml(html string) string {
	return this.GetHeaderHtml() + `<div class="container content_warp"><section class="content-header">
  	<ul class="breadcrumb">
	  <li>
        <a href='{{.homeUrl}}'>
          <span class="glyphicon glyphicon-dashboard" aria-hidden="true"></span>
          首页
        </a>
      </li>
      {{range $breadcrumb := .breadcrumbs}}
      <li>
      	{{if eq $breadcrumb.Url ""}}
          {{$breadcrumb.Label}}
        {{else}}
        <a href='{{$breadcrumb.Url}}'>
          {{$breadcrumb.Label}}
        </a>
        {{end}}
      </li>
      {{end}}
    </ul>
  </section>` + this.GetAlertHtml() + html + "</div>" + this.GetFooterHtml()
}

func (this *BaseTemplate) GetHeaderHtml() string {
	return `
<!DOCTYPE html>
<html lang='zh-cn'>
<head>
	<meta charset='UTF-8'/>
	<meta name='viewport' content='width=device-width, initial-scale=1'>
	<meta name='csrf-param' content='_csrf-backend'>
	<title>路由列表</title>
	` + registerGloablCss() + this.registerSelfCss() + `
</head>
<body>
	<div class="navbar navbar-inverse navbar-fixed-top">
    <div class="container">
      <div class="navbar-collapse collapse" role="navigation">
        <ul class="nav navbar-nav" id="top_menu">
          <li class="hidden-sm hidden-md">
          	<a href="/route/index">路由管理</a></li>
          <li><a href="/permission/index">权限管理</a></li>
          <li><a href="/role/index">角色管理</a></li>
          <li><a href="/assignment/index">角色分配</a></li>
        </ul>
      </div>
    </div>
  </div>
`
}

func (this *BaseTemplate) GetAlertHtml() string {
	return `{{$error_len := len .alert_messages.Error_messages}}
{{$success_len := len .alert_messages.Success_messages}}
{{$info_len := len .alert_messages.Info_messages}}
{{$warning_len := len .alert_messages.Warning_messages}}

{{if gt $error_len 0}}
<div class="alert alert-danger fade in">
	<button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
	<i class='icon fa fa-ban'></i>
	{{range $message := .alert_messages.Error_messages}}
		{{$message}}
	{{end}}
</div>
{{end}}

{{if gt $success_len 0}}
<div class="alert alert-success fade in">
	<button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
	<i class='icon fa fa-check'></i>
	{{range $message := .alert_messages.Success_messages}}
		{{$message}}
	{{end}}
</div>
{{end}}

{{if gt $info_len 0}}
<div class="alert alert-info fade in">
	<button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
	<i class='icon fa fa-info'></i>
	{{range $message := .alert_messages.Info_messages}}
		{{$message}}
	{{end}}
</div>
{{end}}

{{if gt $warning_len 0}}
<div class="alert alert-warning fade in">
	<button type="button" class="close" data-dismiss="alert" aria-hidden="true">×</button>
	<i class='icon fa fa-warning'></i>
	{{range $message := .alert_messages.Warning_messages}}
		{{$message}}
	{{end}}
</div>
{{end}}
`
}
func (this *BaseTemplate) GetFooterHtml() string {
	return registerGloablJs() + this.registerSelfJs() + "</body></html>"
}
