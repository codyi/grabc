package layout

import (
	"github.com/grabc/libs"
	"strings"
)

var DefaultLayout string

func init() {
	DefaultLayout = `
<!DOCTYPE html>
<html lang='zh-cn'>
<head>
	<meta charset='UTF-8'/>
	<meta name='viewport' content='width=device-width, initial-scale=1'>
	<meta name='csrf-param' content='_csrf-backend'>
	<title>路由列表</title>
	<link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css"/>
	<script src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script>
	<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js"></script>
	` + GetGlobalCss() + GetGlobalJs() + `
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
          <li><a href="/assignment/index">用户授权</a></li>
        </ul>
      </div>
    </div>
  </div>
  <div class="container content_warp"><section class="content-header">
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
  </section>
  {{.grabc_content}}
  </div>
<script>
$(function(){
	//设置活动导航
	$.setActiveMenu();

	$("[data-confirm]").click(function () {
		if (confirm($(this).attr("data-confirm"))) {
			var oThis = this;
			$.post($(this).attr("href"), function(response){
				console.log(response)
				if (response.Code == 200) {
					$(oThis).parent().parent().remove();
				}
			});
		}

		return false;
	});
});

//设置导航选中状态
$.setActiveMenu = function(){
	var pathname = window.location.pathname;

	if (pathname.substr(0, 1) == "/") {
		pathname = pathname.substr(1, pathname.length)
	}

	var controllerName = pathname.split("/")[0].toLocaleLowerCase(); //字符分割 


	$("#top_menu").children("li").each(function(){
		var url = $(this).children("a").attr("href").toLocaleLowerCase();

		if (url.indexOf(controllerName) == 1) {
			$(this).children("a").addClass("active")
		}
	});
}
</script>
</body></html>
`
}

type BaseTemplate struct {
}

func (this *BaseTemplate) DealHtml(html string) string {
	grabc_content := GetGlobalCss() + GetGlobalJs() + `` + this.GetAlertHtml() + html

	if libs.Template.Layout != "" {
		return strings.Replace(libs.Template.Layout, "{{.grabc_content}}", grabc_content, 1)
	} else {
		return strings.Replace(DefaultLayout, "{{.grabc_content}}", grabc_content, 1)
	}
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
