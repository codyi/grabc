package menu

import (
	. "github.com/codyi/grabc/views/layout"
)

type Index struct {
	BaseTemplate
}

func (this *Index) Html() string {
	html := `
	<div class="box box-primary">
            <div class="box-body">
                <a class="btn btn-primary" href="/menu/post" menu="button">新增菜单</a>
            </div>
        </div>
        <div class="box box-info">
            <div class="box-body">
                <table class="table table-bordered table-striped">
                    <thead>
                        <tr>
                            <td>ID</td>
                            <td>菜单名称</td>
                            <td>菜单路由</td>
                            <td>菜单排序</td>
                            <td>父级菜单</td>
                            <td>创建时间</td>
                            <td class="row_operate">操作</td>
                        </tr>
                    </thead>
                    <tbody>
                    {{range $index,$menu:=.menus}}
                        <tr>
                            <td>{{$menu.Id}}</td>
                            <td>{{$menu.Name}}</td>
                            <td>{{$menu.Route}}</td>
                            <td>{{$menu.Order}}</td>
                            <td>{{$menu.GetParentName}}</td>
                            <td>{{unixTimeFormat $menu.CreateAt "2006-01-02"}}</td>
                            <td>
                                <a href="/menu/put?menu_id={{$menu.Id}}" title="更新">
                                    <span class="glyphicon glyphicon-pencil"></span>
                                </a>
                                <a href="/menu/delete?menu_id={{$menu.Id}}" title="删除" data-confirm="您确定要删除此项吗？">
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
