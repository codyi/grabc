package role

import (
	. "grabc/views/layout"
)

type Index struct {
	BaseTemplate
}

func (this *Index) Html() string {
	html := `
	<div class="box box-primary">
            <div class="box-body">
                <a class="btn btn-primary" href="/role/post" role="button">新增角色</a>
            </div>
        </div>
        <div class="box box-info">
            <div class="box-body">
                <table class="table table-bordered table-striped">
                    <thead>
                        <tr>
                            <td>ID</td>
                            <td>角色名称</td>
                            <td>角色描述</td>
                            <td>创建时间</td>
                            <td class="row_operate">操作</td>
                        </tr>
                    </thead>
                    <tbody>
                    {{range $index,$role:=.roles}}
                        <tr>
                            <td>{{$role.Id}}</td>
                            <td>{{$role.Name}}</td>
                            <td>{{$role.Description}}</td>
                            <td>{{unixTimeFormat $role.CreateAt "2006-01-02"}}</td>
                            <td>
                                <a href="/role/get?role_id={{$role.Id}}" title="授权">
                                    <span class="glyphicon glyphicon-eye-open"></span>
                                </a>
                                <a href="/role/put?role_id={{$role.Id}}" title="更新">
                                    <span class="glyphicon glyphicon-pencil"></span>
                                </a>
                                <a href="/role/delete?role_id={{$role.Id}}" title="删除" data-confirm="您确定要删除此项吗？">
                                    <span class="glyphicon glyphicon-trash"></span>
                                </a>
                            </td>
                        </tr>
                        {{end}}
                    </tbody>
                </table>
                {{pagination .pages}}
            </div>
        </div>
	`

	return this.DealHtml(html)
}
