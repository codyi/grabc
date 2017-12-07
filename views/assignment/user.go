package assignment

import (
	. "grabc/views/layout"
)

type User struct {
	BaseTemplate
}

func (this *User) Html() string {
	this.SelfJsAppend("http://127.0.0.1:8080/static/html/role_assignment.js")
	html := `
<div class="box box-info">
    <div class="box-header">
        授权用户：<span>{{.name}}</span>
    </div>
    <div class="box-body">
        <table class="table route_warp">
            <tbody>
                <tr>
                    <td style="width: 40%">
                        <select multiple="" size="20" id="unassignment_role">
                        </select>
                    </td>
                    <td>
                        <div>
                            <button id="add_role" class="btn btn-primary">>>添加</button>
                        </div>
                        <div style="margin-top: 15px">
                            <button id="remove_role" class="btn btn-danger"><<删除</button>
                        </div>
                    </td>
                    <td style="width: 40%">
                        <select multiple="" size="20" id="assignment_role">
                        </select>
                    </td>
                </tr>
            </tbody>
        </table>
        <input type="hidden" name="permission_id" id="permission_id" value="{{.model.Id}}">
    </div>
</div>
<script>
    var unassignmentRoles = new Array();
    var assignmentRoles = new Array();
    {{range $index,$role := .unassignmentRoles}}
    unassignmentRoles.push("{{$role}}")
    {{end}}
    {{range $index,$role := .assignmentRoles}}
    assignmentRoles.push("{{$role}}")
    {{end}}
</script>
`
	return this.DealHtml(html)
}
