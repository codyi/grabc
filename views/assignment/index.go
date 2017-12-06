package assignment

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
            <a class="btn btn-primary" href="/user/add" role="button">新增权限</a>
        </div>
    </div>
    <div class="box box-info">
        <div class="box-body">
            <table class="table table-bordered table-striped">
                <thead>
                    <tr>
                        <td>用户</td>
                        <td>用户姓名</td>
                        <td class="row_operate">授权</td>
                    </tr>
                </thead>
                <tbody>
                {{range $id,$name:=.userList}}
                    <tr>
                        <td>{{$id}}</td>
                        <td>{{$name}}</td>
                        <td>
                            <a href="/assignment/user?user_id={{$id}}" title="授权">
                                <span class="glyphicon glyphicon-eye-open"></span>
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
