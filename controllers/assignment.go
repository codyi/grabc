package controllers

import (
	"grabc/views/route"
)

type AssignmentController struct {
	BaseController
}

func (this *AssignmentController) Index() {
	this.ShowHtml(&route.Index{})
}
