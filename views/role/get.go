package role

import (
	. "grabc/views/layout"
)

type Get struct {
	BaseTemplate
}

func (this *Get) Html() string {
	html := ``

	return this.DealHtml(html)
}
