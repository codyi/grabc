package role

import (
	. "grabc/views/layout"
)

type Delete struct {
	BaseTemplate
}

func (this *Delete) Html() string {
	html := ``

	return this.DealHtml(html)
}
