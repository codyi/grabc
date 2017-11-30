package views

type RouteIndex struct {
	baseTemplate
}

func (this *RouteIndex) Html() string {
	html := `
	<div class="container" style="padding-top: 60px">
		<div class="row">
			<select multiple="" size="20" class="col-md-5" id="route_select">
			    <option value="/banner/*">/banner/*</option>
                <option value="/banner/1">/banner/*</option>
                <option value="/banner/2">/banner/*</option>
                <option value="/banner/3">/banner/*</option>
                <option value="/banner/4">/banner/*</option>
                <option value="/banner/5">/banner/*</option>
			</select>
			<div class="col-md-1">
				<div>
					<button id="add_route" class="btn btn-primary">>>添加</button>
				</div>
				<div style="margin-top: 15px">
					<button id="remove_route" class="btn btn-danger"><<删除</button>
				</div>
			</div>
			<select multiple="" size="20" class="col-md-5" id="route_selected">
			    {{range $index,$route := .insertRoutes}}
                <option value="{{$route}}">{{$route}}</option>
                {{end}}
			</select>
		</div>
	</div>
`
	return this.dealHtml(html)
}
