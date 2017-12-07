package role

import (
	. "grabc/views/layout"
)

type Assignment struct {
	BaseTemplate
}

func (this *Assignment) Html() string {
	this.SelfJsAppend("http://127.0.0.1:8080/static/html/assignment_permission.js")
	html := `
<div class="box box-info">
    <div class="box-body">
        <table class="table route_warp">
            <tbody>
                <tr>
                    <td style="width: 40%">
                        <select multiple="" size="20" id="unassignment_permissions">
                        </select>
                    </td>
                    <td>
                        <div>
                            <button id="add_permission" class="btn btn-primary">>>添加</button>
                        </div>
                        <div style="margin-top: 15px">
                            <button id="remove_permission" class="btn btn-danger"><<删除</button>
                        </div>
                    </td>
                    <td style="width: 40%">
                        <select multiple="" size="20" id="assignment_permissions">
                        </select>
                    </td>
                </tr>
            </tbody>
        </table>
        <input type="hidden" name="role_id" id="role_id" value="{{.model.Id}}">
    </div>
</div>
<script>
    var unassignmentPermissions = new Array();
    var assignmentPermissions = new Array();
    {{range $index,$name := .unassignmentPermssionNames}}
    unassignmentPermissions.push("{{$name}}")
    {{end}}
    {{range $index,$name := .assignmentPermssionNames}}
    assignmentPermissions.push("{{$name}}")
    {{end}}
</script>
`

	return this.DealHtml(html)
}
