package assignment

import (
	. "github.com/grabc/views/layout"
)

type Index struct {
	BaseTemplate
}

func (this *Index) Html() string {
	html := `
    <div class="box box-info">
        <div class="box-body">
            <table class="table table-bordered table-striped">
                <thead>
                    <tr>
                        <td>用户</td>
                        <td>用户姓名</td>
                        <td>已授权</td>
                        <td class="row_operate">操作</td>
                    </tr>
                </thead>
                <tbody>
                {{range $id,$user:=.userItems}}
                    <tr>
                        <td>{{$user.Id}}</td>
                        <td>{{$user.Name}}</td>
                        <td>
                        {{range $index,$roleName:=$user.RoleNames}}
                        【{{$roleName}}】&nbsp;
                        {{end}}
                        </td>
                        <td>
                            <a href="/assignment/user?user_id={{$user.Id}}" title="授权">
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
