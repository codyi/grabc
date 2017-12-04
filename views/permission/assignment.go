package permission

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
                        <select multiple="" size="20" id="route_select">
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
                        <select multiple="" size="20" id="route_selected">
                        </select>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</div>
	`

	return this.DealHtml(html)
}
