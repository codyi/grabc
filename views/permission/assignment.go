package permission

import (
	. "github.com/codyi/grabc/views/layout"
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
                        <select multiple="" size="20" id="select_unassignment_routes">
                        </select>
                    </td>
                    <td>
                        <div>
                            <button id="btn_add_route" class="btn btn-primary">>>添加</button>
                        </div>
                        <div style="margin-top: 15px">
                            <button id="btn_remove_route" class="btn btn-danger"><<删除</button>
                        </div>
                    </td>
                    <td style="width: 40%">
                        <select multiple="" size="20" id="select_assignment_routes">
                        </select>
                    </td>
                </tr>
            </tbody>
        </table>
        <input type="hidden" name="permission_id" id="permission_id" value="{{.model.Id}}">
    </div>
</div>
<script>
    var unassignment_routes = new Array();
    var assignment_routes = new Array();
    {{range $index,$route := .allRoutes}}
    unassignment_routes.push("{{$route}}")
    {{end}}
    {{range $index,$route := .assignmentRoutes}}
    assignment_routes.push("{{$route}}")
    {{end}}

    $(function(){
        $.showSelectOption("#select_assignment_routes", assignment_routes);
        $.showSelectOption("#select_unassignment_routes", unassignment_routes);

        //添加路由绑定事件
        $.addRoute();
        //删除路由绑定事件
        $.removeRoute();
    });

    //添加权限绑定事件
    $.addRoute = function () {
        $("#btn_add_route").click(function () {
            var select_routes = $("#select_unassignment_routes").val();

            if (select_routes.length > 0) {
                $("#btn_add_route").attr("disabled","disabled");

                $(select_routes).each(function (index, value) {
                    $.ajax({
                        type:"post",
                        url:"/permission/ajaxaddroute",
                        data:{route:value,permissionId:$("#permission_id").val()},
                        dataType:"json",
                        async:false,
                        success:function (response) {
                            if (response.Code == 200) {
                                assignment_routes.push(response.Data.route);
                                unassignment_routes = $.removeItem(response.Data.route, unassignment_routes);
                                $("#select_unassignment_routes option[value='"+response.Data.route+"']").remove();
                                $.showSelectOption("#select_assignment_routes", assignment_routes);
                            }
                        }
                    });
                });

                $("#btn_add_route").removeAttr("disabled");
            }
        });
    };

    //删除路由绑定事件
    $.removeRoute = function () {
        $("#btn_remove_route").click(function () {
            var select_routes = $("#select_assignment_routes").val();

            if (select_routes.length > 0) {
                $("#btn_remove_route").attr("disabled","disabled");

                $(select_routes).each(function (index, value) {
                    $.ajax({
                        type:"post",
                        url:"/permission/ajaxremoveroute",
                        data:{route:value,permissionId:$("#permission_id").val()},
                        dataType:"json",
                        async:false,
                        success:function (response) {
                            if (response.Code == 200) {
                                unassignment_routes.push(response.Data.route);
                                assignment_routes = $.removeItem(response.Data.route, assignment_routes);
                                $("#select_assignment_routes option[value='"+response.Data.route,+"']").remove();
                                $.showSelectOption("#select_unassignment_routes", unassignment_routes);
                            }
                        }
                    });
                });

                $("#btn_remove_route").removeAttr("disabled");
            }
        });
    };
</script>
	`

	return this.DealHtml(html)
}
