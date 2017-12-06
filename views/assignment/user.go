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
                        <select multiple="" size="20" id="all_route">
                        </select>
                    </td>
                    <td>
                        <div>
                            <button id="add_route" class="btn btn-primary">>>添加</button>
                        </div>
                        <div style="margin-top: 15px">
                            <button id="remove_route" class="btn btn-danger"><<删除</button>
                        </div>
                    </td>
                    <td style="width: 40%">
                        <select multiple="" size="20" id="assignment_route">
                        </select>
                    </td>
                </tr>
            </tbody>
        </table>
        <input type="hidden" name="permission_id" id="permission_id" value="{{.model.Id}}">
    </div>
</div>
`
	return this.DealHtml(html)
}
