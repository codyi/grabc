package route

import (
	. "grabc/views/layout"
)

type Index struct {
	BaseTemplate
}

func (this *Index) Html() string {
	this.SelfJsAppend("http://127.0.0.1:8080/static/html/route.js")
	html := `
    <div class="box box-info">
        <div class="box-body">
            <table class="table route_warp">
                <tbody>
                    <tr>
                        <td style="width: 40%">
                            <select multiple="" size="20" id="route_select">
                            </select>
                        </td>
                        <td>
                            <div>
                                <button id="add_route" class="btn btn-primary">>>添加</button>
                            </div>
                            <div style="margin-top: 15px">
                                <button id="remove_route" class="btn btn-danger"><<删除</button>
                            </div>
                        </td>
                        <td style="width: 40%">
                            <select multiple="" size="20" id="route_selected">
                            </select>
                        </td>
                    </tr>
                </tbody>
            </table>
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
	return this.DealHtml(html)
}
