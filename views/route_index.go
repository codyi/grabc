package views

type RouteIndex struct {
	baseTemplate
}

func (this *RouteIndex) Html() string {
	html := `
	<div class="container content_warp">
		<div class="row">
			<select multiple="" class="col-md-5" id="route_select">
			</select>
			<div class="col-md-1">
				<div>
					<button id="add_route" class="btn btn-primary">>>添加</button>
				</div>
				<div style="margin-top: 15px">
					<button id="remove_route" class="btn btn-danger"><<删除</button>
				</div>
			</div>
			<select multiple="" class="col-md-5" id="route_selected">
			</select>
		</div>
	</div>
	<script>
	    var addRoutes = new Array();
        var notAddRoutes = new Array();
		{{range $index,$route := .notAddRoutes}}
		notAddRoutes.push("{{$route}}")
        {{end}}
        {{range $index,$route := .addRoutes}}
        addRoutes.push("{{$route}}")
		{{end}}
    </script>
`
	return this.dealHtml(html)
}
