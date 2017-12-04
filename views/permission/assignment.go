package permission

import (
	. "grabc/views/layout"
)

type Assignment struct {
	BaseTemplate
}

func (this *Assignment) Html() string {
	this.SelfJsAppend("http://127.0.0.1:8080/static/html/permission_assignment.js")
	html := `
<div class="box box-info">
    <div class="box-body">
        <table class="table route_warp">
            <tbody>
                <tr>
                    <td style="width: 40%">
                        <select multiple="" size="20" id="all_route">
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
                        <select multiple="" size="20" id="assignment_route">
                        </select>
                    </td>
                </tr>
            </tbody>
        </table>
        <input type="hidden" name="permission_id" id="permission_id" value="{{.model.Id}}">
    </div>
</div>
<script>
    var allRoutes = new Array();
    var assignmentRoutes = new Array();
    {{range $index,$route := .allRoutes}}
    allRoutes.push("{{$route}}")
    {{end}}
    {{range $index,$route := .assignmentRoutes}}
    assignmentRoutes.push("{{$route}}")
    {{end}}
</script>
	`

	return this.DealHtml(html)
}
