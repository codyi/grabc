package permission

import (
	. "github.com/grabc/views/layout"
)

type Update struct {
	BaseTemplate
	Form
}

func (this *Update) Html() string {
	html := `
<div class="box box-info">
    <div class="box-body">
        <form action="/permission/put?permission_id={{.model.Id}}" method="post" class="form-horizontal">` + this.FormHtml() + `</form>
    </div>
</div>`

	return this.DealHtml(html)
}
