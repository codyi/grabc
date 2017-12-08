package role

import (
	. "grabc/views/layout"
)

type Assignment struct {
	BaseTemplate
}

func (this *Assignment) Html() string {
	html := `
<div class="box box-info">
    <div class="box-body">
        <table class="table route_warp">
            <tbody>
                <tr>
                    <td style="width: 40%">
                        <select multiple="" size="20" id="select_unassignment_permissions">
                        </select>
                    </td>
                    <td>
                        <div>
                            <button id="btn_add_permission" class="btn btn-primary">>>添加</button>
                        </div>
                        <div style="margin-top: 15px">
                            <button id="btn_remove_permission" class="btn btn-danger"><<删除</button>
                        </div>
                    </td>
                    <td style="width: 40%">
                        <select multiple="" size="20" id="select_assignment_permissions">
                        </select>
                    </td>
                </tr>
            </tbody>
        </table>
        <input type="hidden" name="role_id" id="role_id" value="{{.model.Id}}">
    </div>
</div>
<script>
    var unassignment_permissions = new Array();
    var assignment_permissions = new Array();
    {{range $index,$name := .unassignmentPermssionNames}}
    unassignment_permissions.push("{{$name}}")
    {{end}}
    {{range $index,$name := .assignmentPermssionNames}}
    assignment_permissions.push("{{$name}}")
    {{end}}

    $(function(){
        $.showSelectOption("#select_assignment_permissions", assignment_permissions);
        $.showSelectOption("#select_unassignment_permissions", unassignment_permissions);

        //添加权限
        $.addPermission();
        //删除权限
        $.removePermission();
    });

    //添加权限
    $.addPermission = function () {
        $("#btn_add_permission").click(function () {
            var select_values = $("#select_unassignment_permissions").val();

            if (select_values.length > 0) {
                $("#btn_add_permission").attr("disabled","disabled");

                $(select_values).each(function (index, value) {
                    $.ajax({
                        type:"post",
                        url:"/role/ajaxassignment",
                        data:{permission_name:value,role_id:$("#role_id").val()},
                        dataType:"json",
                        async:false,
                        success:function (response) {
                            if (response.Code == 200) {
                                assignment_permissions.push(value);
                                unassignment_permissions = $.removeItem(value, unassignment_permissions);
                                $("#select_unassignment_permissions option[value='"+value+"']").remove();
                                $.showSelectOption("#select_assignment_permissions", assignment_permissions);
                            }
                        }
                    });
                });

                $("#btn_add_permission").removeAttr("disabled");
            }
        });
    };

    //删除权限
    $.removePermission = function () {
        $("#btn_remove_permission").click(function () {
            var select_values = $("#select_assignment_permissions").val();

            if (select_values.length > 0) {
                $("#btn_remove_permission").attr("disabled","disabled");

                $(select_values).each(function (index, value) {
                    $.ajax({
                        type:"post",
                        url:"/role/ajaxunassignment",
                        data:{permission_name:value,role_id:$("#role_id").val()},
                        dataType:"json",
                        async:false,
                        success:function (response) {
                            if (response.Code == 200) {
                                unassignment_permissions.push(value)
                                assignment_permissions = $.removeItem(value, assignment_permissions);
                                $("#select_assignment_permissions option[value='"+value+"']").remove();
                                $.showSelectOption("#select_unassignment_permissions", unassignment_permissions);
                            }
                        }
                    });
                });

                $("#btn_remove_permission").removeAttr("disabled");
            }
        });
    };
</script>
`

	return this.DealHtml(html)
}
