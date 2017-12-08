package assignment

import (
	. "grabc/views/layout"
)

type User struct {
	BaseTemplate
}

func (this *User) Html() string {
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
                        <select multiple="" size="20" id="select_unassignment_role">
                        </select>
                    </td>
                    <td>
                        <div>
                            <button id="btn_add_role" class="btn btn-primary">>>添加</button>
                        </div>
                        <div style="margin-top: 15px">
                            <button id="btn_remove_role" class="btn btn-danger"><<删除</button>
                        </div>
                    </td>
                    <td style="width: 40%">
                        <select multiple="" size="20" id="select_assignment_role">
                        </select>
                    </td>
                </tr>
            </tbody>
        </table>
        <input type="hidden" name="user_id" id="user_id" value="{{.user_id}}">
    </div>
</div>
<script>
    var unassignment_roles = new Array();
    var assignment_roles = new Array();
    {{range $index,$role := .unassignmentRoles}}
    unassignment_roles.push("{{$role}}")
    {{end}}
    {{range $index,$role := .assignmentRoles}}
    assignment_roles.push("{{$role}}")
    {{end}}

    $(function(){
        $.showSelectOption("#select_assignment_role", assignment_roles);
        $.showSelectOption("#select_unassignment_role", unassignment_roles);

        //添加路由
        $.addRole();
        //删除路由
        $.removeRole();
    });

    //添加角色
    $.addRole = function () {
        $("#btn_add_role").click(function () {
            var select_values = $("#select_unassignment_role").val();

            if (select_values.length > 0) {
                $("#btn_add_role").attr("disabled","disabled");

                $(select_values).each(function (index, value) {
                    $.ajax({
                        type:"post",
                        url:"/assignment/add",
                        data:{role:value,user_id:$("#user_id").val()},
                        dataType:"json",
                        async:false,
                        success:function (response) {
                            if (response.Code == 200) {
                                assignment_roles.push(value);
                                unassignment_roles = $.removeItem(value, unassignment_roles);
                                $("#select_unassignment_role option[value='"+value+"']").remove();
                                $.showSelectOption("#select_assignment_role", assignment_roles);
                            }
                        }
                    });
                });

                $("#btn_add_role").removeAttr("disabled");
            }
        });
    };

    //删除角色
    $.removeRole = function () {
        $("#btn_remove_role").click(function () {
            var select_values = $("#select_assignment_role").val();

            if (select_values.length > 0) {
                $("#btn_remove_role").attr("disabled","disabled");

                $(select_values).each(function (index, value) {
                    $.ajax({
                        type:"post",
                        url:"/assignment/remove",
                        data:{role:value,user_id:$("#user_id").val()},
                        dataType:"json",
                        async:false,
                        success:function (response) {
                            if (response.Code == 200) {
                                unassignment_roles.push(value)
                                assignment_roles = $.removeItem(value, assignment_roles);
                                $("#select_assignment_role option[value='"+value+"']").remove();
                                $.showSelectOption("#select_unassignment_role", unassignment_roles);
                            }
                        }
                    });
                });

                $("#btn_remove_role").removeAttr("disabled");
            }
        });
    };
</script>
`
	return this.DealHtml(html)
}
