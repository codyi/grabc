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

func (this *AssignmentController) Post() {
	this.ShowHtml(&route.Index{})
}

func (this *AssignmentController) Put() {
	this.ShowHtml(&route.Index{})
}

func (this *AssignmentController) Get() {
	this.ShowHtml(&route.Index{})
}

func (this *AssignmentController) Delete() {
	this.ShowHtml(&route.Index{})
}
