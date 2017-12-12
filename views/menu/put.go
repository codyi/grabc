package menu

import (
	. "github.com/codyi/grabc/views/layout"
)

type Put struct {
	BaseTemplate
	Form
}

func (this *Put) Html() string {
	html := `
<div class="box box-info">
    <div class="box-body">
        <form action="/menu/put?menu_id={{.model.Id}}" method="post" class="form-horizontal">` + this.FormHtml() + `</form>
    </div>
</div>
`

	return this.DealHtml(html)
}
