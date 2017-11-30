package views

type baseTemplate struct {
}

func (this *baseTemplate) dealHtml(html string) string {
	return this.getHeaderHtml() + html + this.getFooterHtml()
}

func (this *baseTemplate) getHeaderHtml() string {
	return `
<!DOCTYPE html>
<html lang='zh-cn'>
<head>
	<meta charset='UTF-8'/>
	<meta name='viewport' content='width=device-width, initial-scale=1'>
	<meta name='csrf-param' content='_csrf-backend'>
	<title>路由列表</title>
	<link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
	<link rel="stylesheet" href="http://127.0.0.1:8080/static/html/global.css" crossorigin="anonymous">
	` + getGlobalCss() + `
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

func (this *baseTemplate) getFooterHtml() string {
	return `
	<script src="https://cdn.bootcss.com/jquery/3.2.1/jquery.min.js"></script>
	<script src="https://cdn.bootcss.com/bootstrap/3.3.7/js/bootstrap.min.js" integrity="sha384-Tc5IQib027qvyjSMfHjOMaLkfuWVxZxUPnCJA7l2mCWNIpG9mGCD8wGNIcPD7Txa" crossorigin="anonymous"></script>
	<script src="http://127.0.0.1:8080/static/html/global.js"></script>
` + getGlobalJs() + "</body></html>"
}
