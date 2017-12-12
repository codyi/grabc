package menu

import (
	. "github.com/codyi/grabc/views/layout"
)

type Post struct {
	BaseTemplate
	Form
}

func (this *Post) Html() string {
	html := `
<div class="box box-info">
    <div class="box-body">
        <form action="/menu/post" method="post" class="form-horizontal">
        ` + this.FormHtml() + `
		</form>
    </div>
</div>
	`

	return this.DealHtml(html)
}
