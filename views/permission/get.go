package permission

import (
	. "grabc/views/layout"
)

type Get struct {
	BaseTemplate
}

func (this *Get) Html() string {
	html := `
<div class="box box-info">
    <div class="box-body">
        <table id="w0" class="table table-striped table-bordered detail-view">
            <tbody>
                <tr><th>名称</th><td>{{.model.Name}}</td></tr>
                <tr><th>描述</th><td>{{.model.Description}}</td></tr>
            </tbody>
        </table>
    </div>
</div>
	`

	return this.DealHtml(html)
}
