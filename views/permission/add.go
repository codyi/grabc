package permission

import (
	. "github.com/grabc/views/layout"
)

type Add struct {
	BaseTemplate
	Form
}

func (this *Add) Html() string {
	html := `
<div class="box box-info">
    <div class="box-body">
        <form action="/permission/add" method="post" class="form-horizontal">
        ` + this.FormHtml() + `
		</form>
    </div>
</div>`

	return this.DealHtml(html)
}
