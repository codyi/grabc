package permission

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
            <a class="btn btn-primary" href="/permission/add" role="button">新增权限</a>
        </div>
    </div>
    <div class="box box-info">
        <div class="box-body">
            <table class="table table-bordered table-striped">
                <thead>
                    <tr>
                        <td>ID</td>
                        <td>权限名称</td>
                        <td>权限描述</td>
                        <td>操作</td>
                    </tr>
                </thead>
                <tbody>
                {{range $index,$permission:=.permissions}}
                    <tr>
                        <td>{{$permission.Id}}</td>
                        <td>{{$permission.Name}}</td>
                        <td>{{$permission.Description}}</td>
                        <td>操作</td>
                    </tr>
                    {{end}}
                </tbody>
            </table>
            <nav aria-label="Page navigation" class="pagination_warp">
                <ul class="pagination">
                    <li>
                        <a href="#" aria-label="Previous">
                            <span aria-hidden="true">&laquo;</span>
                        </a>
                    </li>
                    <li><a href="#">1</a></li>
                    <li><a href="#">2</a></li>
                    <li><a href="#">3</a></li>
                    <li><a href="#">4</a></li>
                    <li><a href="#">5</a></li>
                    <li>
                        <a href="#" aria-label="Next">
                            <span aria-hidden="true">&raquo;</span>
                        </a>
                    </li>
                </ul>
            </nav>
        </div>
    </div>
`
	return this.DealHtml(html)
}
